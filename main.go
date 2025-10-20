package main

import (
	"falling_rest/api"
	"fmt"
	"strings"

	//"encoding/json"
	"html/template"
	"net/http"
	"os"
	//"strings"
	//"strconv"
)

/*type Item struct {
	ID int `json:"id"`
	NAME string `json:name`
}

var testItems = []Item{
	{ID: 1, Name: "Apple"},
	{ID: 2, Name: "Bananae"},
}*/

var extensions = map[string]string{
	"css":  "text/css",
	"js":   "application/javascript",
	"ico":  "image/x-icon",
	"png":  "image/png",
	"gif":  "image/gif",
	"jpg":  "image/jpeg",
	"jpeg": "image/jpeg",
	"svg":  "image/svg+xml",
	"mp3":  "audio/mpeg",
	"wav":  "audio/x-wav",
	"ogg":  "application/ogg", //Why?
}

var endpoints = map[string]func() []byte{
	"seasonalFacts": api.Season,
}

func handler(w http.ResponseWriter, r *http.Request) {
	fullUrl := r.URL.String()
	suffix := strings.Split(fullUrl, ".")

	//Handle files
	if strings.HasSuffix(fullUrl, ".js") || strings.HasSuffix(fullUrl, ".css") || strings.HasSuffix(fullUrl, ".ico") {
		data, err := os.ReadFile(fmt.Sprintf("public%s", fullUrl))
		if err != nil {
			fmt.Println(err)
		} else {
			w.Header().Set("Content-Type", fmt.Sprintf("%s", extensions[suffix[1]]))
			w.Write(data)
		}
		return
	}

	//Handle API endpoints
	urlSplit := strings.Split(fullUrl, "/")
	if urlSplit[1] == "api" {
		apiEndpoint := urlSplit[2]
		if _, ok := endpoints[apiEndpoint]; !ok {
			return
		}
		data := endpoints[apiEndpoint]()

		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
		return
	}

	//Handle and render HTML
	fullUrl = strings.Replace(fullUrl, ".html", "", -1)

	if fullUrl == "/" {
		fullUrl = "/index"
	}

	content, err := os.ReadFile(fmt.Sprintf("public%s.html", fullUrl))
	if err != nil {
		//fmt.Println(err)
		content, err = os.ReadFile("public/404.html")
	}

	tmpl := template.Must(template.New("index").Parse(string(content)))

	switch r.Method {

	case http.MethodGet:
		tmpl.Execute(w, nil)

	case http.MethodPost:
		//Dont actually need POST requests but I'll leave them just in case
		if err := r.ParseForm(); err != nil {
			http.Error(w, "unable to parse", http.StatusBadRequest)
			return
		}

		message := r.FormValue("message")

		tmpl.Execute(w, message)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/", handler)

	fmt.Println("Server started")
	http.ListenAndServe(":8080", nil)
}

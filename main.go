package main

import (
	"bufio"
	"bytes"
	"falling_rest/api"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
	//"strconv"
)

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
	"seasonalFacts":   api.Season,
	"scientificFacts": api.Science,
	"leavesImages":    api.LeafImage,
}

var vars = map[string]interface{}{}

// How the fuck did I get this working...
func parseHTML(content []byte) string { //Fuck if I know what im doing here but hey worth a shot
	reader := bytes.NewReader(content)
	scanner := bufio.NewScanner(reader)
	var result string
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "{{") {
			variableName := strings.TrimSpace(strings.Split(strings.Split(line, "server.")[1], "}}")[0])
			if ptr, ok := vars[variableName].(*string); ok {
				parsed := strings.ReplaceAll(line, variableName, *ptr)
				parsedAgain := strings.ReplaceAll(parsed, "{{", "")
				parsedAgain2 := strings.ReplaceAll(parsedAgain, "server.", "")
				parsedAgain3 := strings.ReplaceAll(parsedAgain2, "}}", "")
				result += parsedAgain3 + "\n"
			} else {
				fmt.Println("variable not a string")
				result += line + "\n"
			}
		} else {
			result += line + "\n"
		}
	}
	return result
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
	var parsed string = parseHTML(content)
	tmpl := template.Must(template.New("index").Parse(parsed))

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

func loadEnv() {
	wd, _ := os.Getwd()
	godotenv.Load(filepath.Join(wd, ".env"))
}

func main() {
	loadEnv()

	var websiteName string = os.Getenv("websiteName")
	fmt.Println(websiteName)
	vars["websiteName"] = &websiteName

	http.HandleFunc("/", handler)

	fmt.Println("Server started")
	http.ListenAndServe(":8081", nil) //Steam webhelper uses port 8080 >:(
}

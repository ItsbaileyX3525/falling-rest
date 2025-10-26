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
	"sync"
	"time"

	"github.com/joho/godotenv"
)

var extensions = map[string]string{
	"css":  "text/css",
	"js":   "text/javascript",
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

var endpoints = map[string]func([]string) []byte{
	"seasonalFacts":   api.Season,
	"scientificFacts": api.Science,
	"leavesImages":    api.LeafImage,
	"motionImages":    api.MotionImage,
	"decode":          api.DecodeHash,
	"fallPeople":      api.People,
}

var authRoutes = map[string]http.HandlerFunc{
	"register": api.Register,
	"login":    api.Login,
	"logout":   api.Logout,
	"me":       api.Me,
}

var vars = map[string]interface{}{}

// Rate limiting: per-IP fixed window limiter
type clientLimit struct {
	count       int
	windowStart time.Time
}

var (
	rateLimiters = make(map[string]*clientLimit)
	rateMu       = sync.Mutex{}
)

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
				parsed := strings.ReplaceAll(line, variableName, "Undefined")
				parsedAgain := strings.ReplaceAll(parsed, "{{", "")
				parsedAgain2 := strings.ReplaceAll(parsedAgain, "server.", "")
				parsedAgain3 := strings.ReplaceAll(parsedAgain2, "}}", "")
				result += parsedAgain3 + "\n"
			}
		} else {
			result += line + "\n"
		}
	}
	return result
}

// getClientIP attempts to return the real client IP using X-Forwarded-For when present,
// otherwise falls back to RemoteAddr.
func getClientIP(r *http.Request) string {
	// X-Forwarded-For may contain a comma-separated list; take the first
	if xf := r.Header.Get("X-Forwarded-For"); xf != "" {
		parts := strings.Split(xf, ",")
		if len(parts) > 0 {
			return strings.TrimSpace(parts[0])
		}
	}
	// Fallback to RemoteAddr (host:port)
	host := r.RemoteAddr
	if idx := strings.LastIndex(host, ":"); idx != -1 {
		return host[:idx]
	}
	return host
}

func handler(w http.ResponseWriter, r *http.Request) {
	fullUrl := r.URL.String()
	suffix := strings.Split(fullUrl, ".")

	//Handle files
	if len(suffix) > 1 {
		if strings.HasSuffix(fullUrl, "."+suffix[1]) {
			data, err := os.ReadFile(fmt.Sprintf("public%s", fullUrl))
			if err != nil {
				fmt.Println(err)
			} else {
				w.Header().Set("Content-Type", extensions[suffix[1]])
				w.Write(data)
			}
			return
		}
	}

	//Handle API endpoints
	urlSplit := strings.Split(fullUrl, "/")
	if len(urlSplit) > 1 && strings.TrimSpace(urlSplit[1]) == "api" {
		clientIP := getClientIP(r)
		const maxReq = 30
		const windowDur = time.Minute
		now := time.Now()
		rateMu.Lock()
		lim, ok := rateLimiters[clientIP]
		if !ok || now.Sub(lim.windowStart) >= windowDur {
			rateLimiters[clientIP] = &clientLimit{count: 1, windowStart: now}
		} else {
			if lim.count >= maxReq {
				rateMu.Unlock()
				http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
				return
			}
			lim.count++
		}
		rateMu.Unlock()
		if len(urlSplit) > 2 {
			apiEndpoint := urlSplit[2]
			params := strings.Split(apiEndpoint, "?")
			if _, ok := endpoints[params[0]]; ok {
				endpointName := params[0]

				if endpointName == "decode" || endpointName == "fallPeople" {
					apiKey := ""
					if len(params) > 1 {
						for _, param := range params[1:] {
							if strings.HasPrefix(param, "apiKey=") {
								apiKey = strings.TrimPrefix(param, "apiKey=")
								if idx := strings.Index(apiKey, "&"); idx != -1 {
									apiKey = apiKey[:idx]
								}
								break
							}
							if idx := strings.Index(param, "apiKey="); idx != -1 {
								rest := param[idx+len("apiKey="):]
								if amp := strings.Index(rest, "&"); amp != -1 {
									rest = rest[:amp]
								}
								apiKey = rest
								break
							}
						}
					}
					if apiKey == "" {
						http.Error(w, "API key required", http.StatusUnauthorized)
						return
					}
					if _, err := api.ValidateAPIKey(apiKey); err != nil {
						http.Error(w, "Invalid API key", http.StatusUnauthorized)
						return
					}
				}

				data := endpoints[endpointName](params[1:])
				w.Header().Set("Content-Type", "application/json")
				w.Write(data)
				return
			}
		}
		http.Error(w, "API endpoint not found", http.StatusNotFound)
		return
	}

	//Handle auth endpoints
	if len(urlSplit) > 1 && strings.TrimSpace(urlSplit[1]) == "auth" {
		if len(urlSplit) > 2 {
			authEndpoint := urlSplit[2]
			if handler, ok := authRoutes[authEndpoint]; ok {
				handler(w, r)
				return
			}
		}
		http.Error(w, "Auth endpoint not found", http.StatusNotFound)
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
		content, _ = os.ReadFile("public/404.html")
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
	var websiteVersion string = os.Getenv("websiteVersion")
	var test string = "placeholder"
	vars["websiteName"] = &websiteName
	vars["websiteVersion"] = &websiteVersion
	vars["test"] = &test
	http.HandleFunc("/", handler)

	fmt.Println("Server started")
	http.ListenAndServe(":8081", nil) //Steam webhelper uses port 8080 >:(
}

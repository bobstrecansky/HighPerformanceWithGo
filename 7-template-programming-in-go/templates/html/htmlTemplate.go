package main

import (
	"html/template"
	"log"
	"net/http"
)

type UserFields struct {
	Name  string
	URL   string
	Email string
}

var userResponse = `
<html>
<head></head>
<body>
<h1>Hello {{.Name}}</h1>
<p>You visited {{.URL}}</p>
<p>Hope you're enjoying this book!</p>
<p>We have your email recorded as {{.Email}}</p>
</body>
</html>
`

func rootHandler(w http.ResponseWriter, r *http.Request) {
	requestedURL := string(r.URL.Path)
	userfields := UserFields{"Bob", requestedURL, "bob@example.com"}
	t := template.Must(template.New("HTML Body").Parse(userResponse))
	t.Execute(w, userfields)
	log.Printf("User " + userfields.Name + " Visited : " + requestedURL)
}

func main() {
	s := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/", rootHandler)
	s.ListenAndServe()
}

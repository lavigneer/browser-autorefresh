package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	autorefresh "github.com/lavigneer/browser-autorefresh"
)

var MainTemplate = template.Must(template.New("main").Parse(`
<!DOCTYPE html>
<html>
	<head>
		<meta charset="UTF-8">
		<title>Hello</title>
		{{ template "autorefresh" . }}
	</head>
	<body>
		{{ . }}
	</body>
</html>
`))

func main() {
	// Create a template off of the MainTemplate and pass it in so it can be rendered in the MainTemplate
	t := MainTemplate.New("autorefresh")
	a, _ := autorefresh.New(t, "/__dev/auto-refresh", 100)

	mux := http.NewServeMux()
	mux.Handle(a.Path, a)
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		MainTemplate.Execute(w, time.Now().String())
	})

	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Printf("%v", err)
	}
}

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/AankTia/learning-golang/pkg/config"
	"github.com/AankTia/learning-golang/pkg/handlers"
	"github.com/AankTia/learning-golang/pkg/render"
)

var portNumber = ":8080"

// main is the main application function
func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}

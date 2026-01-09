package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"project-app-portfolio-golang-fathoni/database"
	"project-app-portfolio-golang-fathoni/handler"
	"project-app-portfolio-golang-fathoni/repository"
	"project-app-portfolio-golang-fathoni/router"
	"project-app-portfolio-golang-fathoni/service"
)

func main() {
	var templates = template.Must(template.New("").ParseGlob("views/**/*.html"))

	db, err := database.InitDB()
	if err != nil {
		panic(err)
	}

	for _, t := range templates.Templates() {
		fmt.Println("template:", t.Name())
	}

	repo := repository.NewRepository(db)
	service := service.NewService(repo)
	handler := handler.NewHandler(service, templates)

	r := router.NewRouter(handler)

	fmt.Println("server running on port 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("error server")
	}
}

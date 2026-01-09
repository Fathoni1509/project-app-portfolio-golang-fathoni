package handler

import (
	"html/template"
	"project-app-portfolio-golang-fathoni/service"
)

type Handler struct {
	HandlerMenu     MenuHandler
	PersonalHandler PersonalHandler
}

func NewHandler(service service.Service, templates *template.Template) Handler {
	return Handler{
		HandlerMenu: NewMenuHandler(templates, service.PersonalService),
		PersonalHandler: NewPersonalHandler(templates, service.PersonalService),
	}
}

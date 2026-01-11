package handler

import (
	"html/template"
	"project-app-portfolio-golang-fathoni/service"
)

type Handler struct {
	HandlerMenu     MenuHandler
	PersonalHandler PersonalHandler
	ActivityHandler ActivityHandler
}

func NewHandler(service service.Service, templates *template.Template) Handler {
	return Handler{
		HandlerMenu: NewMenuHandler(templates, service.PersonalService, service.ActivityService),
		PersonalHandler: NewPersonalHandler(templates, service.PersonalService),
		ActivityHandler: NewActivityHandler(templates, service.ActivityService),
	}
}

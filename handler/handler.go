package handler

import (
	"html/template"
	"project-app-portfolio-golang-fathoni/service"
)

type Handler struct {
	HandlerMenu     MenuHandler
	PersonalHandler PersonalHandler
	ActivityHandler ActivityHandler
	WorkHandler     WorkHandler
}

func NewHandler(service *service.Service, templates *template.Template) Handler {
	return Handler{
		// HandlerMenu:     NewMenuHandler(templates, service.PersonalService, service.ActivityService, service.WorkService),
		HandlerMenu:     NewMenuHandler(templates, service),
		PersonalHandler: NewPersonalHandler(templates, service.PersonalService),
		ActivityHandler: NewActivityHandler(templates, service.ActivityService),
		WorkHandler:     NewWorkHandler(templates, service.WorkService),
	}
}

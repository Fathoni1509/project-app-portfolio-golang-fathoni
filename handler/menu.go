package handler

import (
	"fmt"
	"html/template"
	"net/http"
	"project-app-portfolio-golang-fathoni/dto"
	"project-app-portfolio-golang-fathoni/service"
)

type MenuHandler struct {
	Templates       *template.Template
	PersonalService service.PersonalService
	ActivityService service.ActivityService
	WorkService     service.WorkService
	ProjectService  service.ProjectService
	ContactService  service.ContactService
	Services        *service.Service
}

// func NewMenuHandler(templates *template.Template, personalService service.PersonalService, activityService service.ActivityService, workService service.WorkService) MenuHandler {
// 	return MenuHandler{
// 		Templates: templates,
// 		PersonalService: personalService,
// 		ActivityService: activityService,
// 		WorkService: workService,
// 	}
// }

func NewMenuHandler(templates *template.Template, service *service.Service) MenuHandler {
	return MenuHandler{
		Templates: templates,
		Services:  service,
	}
}

func (h *MenuHandler) PortfolioView(w http.ResponseWriter, r *http.Request) {
	personal, err := h.Services.PersonalService.GetDataPersonal()
	if err != nil {
		fmt.Println("error get data personal:", err)
	}

	// activity, err := h.ActivityService.GetDataActivity()
	activity, err := h.Services.ActivityService.GetDataActivity()
	if err != nil {
		fmt.Println("error get data activity:", err)
	}

	work, err := h.Services.WorkService.GetDataWork()
	if err != nil {
		fmt.Println("error get data work:", err)
	}

	project, err := h.Services.ProjectService.GetDataProject()
	if err != nil {
		fmt.Println("error get data work:", err)
	}

	contact, err := h.Services.ContactService.GetDataContact()
	if err != nil {
		fmt.Println("error get data contact:", err)
	}

	pageData := dto.PortfolioPage{
		Personal: personal,
		Activity: activity,
		Work:     work,
		Project:  project,
		Contact:  contact,
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := h.Templates.ExecuteTemplate(w, "main_portfolio", pageData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *MenuHandler) EditView(w http.ResponseWriter, r *http.Request) {
	personal, err := h.Services.PersonalService.GetDataPersonal()
	if err != nil {
		fmt.Println("error get data personal:", err)
	}

	activity, err := h.Services.ActivityService.GetDataActivity()
	if err != nil {
		fmt.Println("error get data activity:", err)
	}

	work, err := h.Services.WorkService.GetDataWork()
	if err != nil {
		fmt.Println("error get data work:", err)
	}

	project, err := h.Services.ProjectService.GetDataProject()
	if err != nil {
		fmt.Println("error get data work:", err)
	}

	contact, err := h.Services.ContactService.GetDataContact()
	if err != nil {
		fmt.Println("error get data contact:", err)
	}

	existingData := dto.PortfolioPage{
		Personal: personal,
		Activity: activity,
		Work:     work,
		Project:  project,
		Contact:  contact,
		Path:     r.URL.Path,
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := h.Templates.ExecuteTemplate(w, "main_edit", existingData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// ga guna
func (h *MenuHandler) PersonalView(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := h.Templates.ExecuteTemplate(w, "personal", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

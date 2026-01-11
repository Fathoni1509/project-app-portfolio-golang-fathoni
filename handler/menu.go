package handler

import (
	"fmt"
	"html/template"
	"net/http"
	"project-app-portfolio-golang-fathoni/dto"
	"project-app-portfolio-golang-fathoni/service"
)

type MenuHandler struct {
	Templates *template.Template
	PersonalService service.PersonalService
	ActivityService service.ActivityService
}

func NewMenuHandler(templates *template.Template, personalService service.PersonalService, activityService service.ActivityService) MenuHandler {
	return MenuHandler{
		Templates: templates,
		PersonalService: personalService,
		ActivityService: activityService,
	}
}

func (h *MenuHandler) PortfolioView(w http.ResponseWriter, r *http.Request) {
	personal, err := h.PersonalService.GetDataPersonal()
	if err != nil {
		fmt.Println("error get data personal:", err)
	}

	activity, err := h.ActivityService.GetDataActivity()
	if err != nil {
		fmt.Println("error get data activity:", err)
	}

	pageData := dto.PortfolioPage{
		Personal: personal,
		Activity: activity,
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := h.Templates.ExecuteTemplate(w, "main_portfolio", pageData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *MenuHandler) EditView(w http.ResponseWriter, r *http.Request) {
	personal, err := h.PersonalService.GetDataPersonal()
	if err != nil {
		fmt.Println("error get data personal:", err)
	}

	activity, err := h.ActivityService.GetDataActivity()
	if err != nil {
		fmt.Println("error get data activity:", err)
	}

	existingData := dto.PortfolioPage{
		Personal: personal,
		Activity: activity,
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := h.Templates.ExecuteTemplate(w, "main_edit", existingData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *MenuHandler) PersonalView(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := h.Templates.ExecuteTemplate(w, "personal", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
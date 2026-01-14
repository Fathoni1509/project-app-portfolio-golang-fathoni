package handler

import (
	"html/template"
	"net/http"
	"project-app-portfolio-golang-fathoni/dto"
	"project-app-portfolio-golang-fathoni/service"
	"project-app-portfolio-golang-fathoni/utils"
)

type PersonalHandler struct {
	PersonalService service.PersonalService
	Templates *template.Template
}

func NewPersonalHandler(templates *template.Template,personalService service.PersonalService) PersonalHandler {
	return PersonalHandler{
		PersonalService: personalService,
		Templates: templates,
	}
}

// update data
func (personalHandler *PersonalHandler) UpdatePersonal(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		http.Error(w, "failed process form", http.StatusBadRequest)
		return
	}

	nameStr := r.FormValue("personal_name") 
    descStr := r.FormValue("personal_description")

	// parsing to model personal
	personal := dto.PersonalUpdateRequest{
		Name: &nameStr,
		Description: &descStr,
	}

	err := personalHandler.PersonalService.UpdatePersonal(&personal)
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusInternalServerError, "error update:"+err.Error(), nil)
		return
	}

	http.Redirect(w, r, "/edit#personal", http.StatusSeeOther)
}
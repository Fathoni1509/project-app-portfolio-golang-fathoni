package handler

import (
	// "encoding/json"
	"html/template"
	"net/http"
	"project-app-portfolio-golang-fathoni/dto"
	"project-app-portfolio-golang-fathoni/service"
	"project-app-portfolio-golang-fathoni/utils"
	"strconv"
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

// get data personal
func (personalHandler *PersonalHandler) GetDataPersonal(w http.ResponseWriter, r *http.Request) {
	personal, err := personalHandler.PersonalService.GetDataPersonal()
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusNotFound, "data not found:"+err.Error(), nil)
		return
	}

	// utils.ResponseSuccess(w, http.StatusOK, "success get data personal", personal)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := personalHandler.Templates.ExecuteTemplate(w, "personal", personal); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// update data
func (personalHandler *PersonalHandler) UpdatePersonal(w http.ResponseWriter, r *http.Request) {
	// var req dto.PersonalUpdateRequest
	// if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
	// 	utils.ResponseBadRequest(w, http.StatusBadRequest, "error data:"+err.Error(), nil)
	// 	return
	// }

	if err := r.ParseForm(); err != nil {
		http.Error(w, "failed process form", http.StatusBadRequest)
		return
	}

	nameStr := r.FormValue("personal_name") 
    ageStr  := r.FormValue("personal_age")
    descStr := r.FormValue("personal_description")

	ageInt, err := strconv.Atoi(ageStr)
    if err != nil {
        http.Error(w, "age must be integer", http.StatusBadRequest)
        return
    }

	// if err := r.FormValue("personal"); err != nil {
	// 	utils.ResponseBadRequest(w, http.StatusBadRequest, "error data:"+err.Error(), nil)
	// 	return
	// }

	// existing, err := personalHandler.PersonalService.GetDataPersonal()
	// if err != nil {
	// 	utils.ResponseBadRequest(w, http.StatusNotFound, "data not found:"+err.Error(), nil)
	// 	return
	// }

	// if req.Name != nil {
	// 	existing.Name = *req.Name
	// }

	// if req.Age != nil {
	// 	existing.Age = *req.Age
	// }

	// if req.Description != nil {
	// 	existing.Description = *req.Description
	// }

	// validation
	// messages, err := utils.ValidateErrors(req)
	// if err != nil {
	// 	utils.ResponseBadRequest(w, http.StatusBadRequest, err.Error(), messages)
	// 	return
	// }

	// parsing to model personal
	personal := dto.PersonalUpdateRequest{
		Name: &nameStr,
		Age: &ageInt,
		Description: &descStr,
	}

	err = personalHandler.PersonalService.UpdatePersonal(&personal)
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusInternalServerError, "error update:"+err.Error(), nil)
		return
	}

	// utils.ResponseSuccess(w, http.StatusOK, "updated success", nil)
	// w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// if err := personalHandler.Templates.ExecuteTemplate(w, "personal_edit", personal); err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// }

	http.Redirect(w, r, "/edit#personal", http.StatusSeeOther)
}
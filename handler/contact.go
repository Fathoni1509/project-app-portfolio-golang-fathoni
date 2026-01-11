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

type ContactHandler struct {
	ContactService service.ContactService
	Templates      *template.Template
}

func NewContactHandler(templates *template.Template, contactService service.ContactService) ContactHandler {
	return ContactHandler{
		ContactService: contactService,
		Templates:      templates,
	}
}

// abaikan get data, get data lakukan di menu
// get data contact
func (contactHandler *ContactHandler) GetDataContact(w http.ResponseWriter, r *http.Request) {
	contact, err := contactHandler.ContactService.GetDataContact()
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusNotFound, "data not found:"+err.Error(), nil)
		return
	}

	// utils.ResponseSuccess(w, http.StatusOK, "success get data contact", contact)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := contactHandler.Templates.ExecuteTemplate(w, "contact", contact); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// create data
func (contactHandler *ContactHandler) CreateContact(w http.ResponseWriter, r *http.Request) {
	// var req dto.ContactUpdateRequest
	// if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
	// 	utils.ResponseBadRequest(w, http.StatusBadRequest, "error data:"+err.Error(), nil)
	// 	return
	// }

	if err := r.ParseForm(); err != nil {
		http.Error(w, "failed process form", http.StatusBadRequest)
		return
	}

	nameStr := r.FormValue("contact_name")
	typeStr := r.FormValue("contact_type")
	linkStr := r.FormValue("contact_link")

	// if err := r.FormValue("contact"); err != nil {
	// 	utils.ResponseBadRequest(w, http.StatusBadRequest, "error data:"+err.Error(), nil)
	// 	return
	// }

	// existing, err := contactHandler.ContactService.GetDataContact()
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

	// parsing to model contact
	contact := dto.ContactCreateRequest{
		Name: nameStr,
		Type: typeStr,
		Link: linkStr,
	}

	err := contactHandler.ContactService.CreateContact(&contact)
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusInternalServerError, "error:"+err.Error(), nil)
		return
	}

	// utils.ResponseSuccess(w, http.StatusOK, "updated success", nil)
	// w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// if err := contactHandler.Templates.ExecuteTemplate(w, "contact_edit", contact); err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// }

	http.Redirect(w, r, "/edit", http.StatusSeeOther)
}

// update data
func (contactHandler *ContactHandler) UpdateContact(w http.ResponseWriter, r *http.Request) {
	// var req dto.ContactUpdateRequest
	// if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
	// 	utils.ResponseBadRequest(w, http.StatusBadRequest, "error data:"+err.Error(), nil)
	// 	return
	// }

	if err := r.ParseForm(); err != nil {
		http.Error(w, "failed process form", http.StatusBadRequest)
		return
	}

	idStr := r.FormValue("contact_id")
	nameStr := r.FormValue("contact_name")
	typeStr := r.FormValue("contact_type")
	linkStr := r.FormValue("contact_link")

	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "id must be integer", http.StatusBadRequest)
		return
	}

	// if err := r.FormValue("contact"); err != nil {
	// 	utils.ResponseBadRequest(w, http.StatusBadRequest, "error data:"+err.Error(), nil)
	// 	return
	// }

	// existing, err := contactHandler.ContactService.GetDataContact()
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

	// parsing to model contact
	contact := dto.ContactUpdateRequest{
		Name: &nameStr,
		Type: &typeStr,
		Link: &linkStr,
	}

	err = contactHandler.ContactService.UpdateContact(idInt, &contact)
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusInternalServerError, "error update:"+err.Error(), nil)
		return
	}

	// utils.ResponseSuccess(w, http.StatusOK, "updated success", nil)
	// w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// if err := contactHandler.Templates.ExecuteTemplate(w, "contact_edit", contact); err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// }

	http.Redirect(w, r, "/edit", http.StatusSeeOther)
}

// delete data
func (contactHandler *ContactHandler) DeleteContact(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		http.Error(w, "failed process form", http.StatusBadRequest)
		return
	}

	idStr := r.FormValue("contact_id")

	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "id must be integer", http.StatusBadRequest)
		return
	}

	err = contactHandler.ContactService.DeleteContact(idInt)
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusInternalServerError, "error delete:"+err.Error(), nil)
		return
	}

	// utils.ResponseSuccess(w, http.StatusOK, "updated success", nil)
	// w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// if err := contactHandler.Templates.ExecuteTemplate(w, "contact_edit", contact); err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// }

	http.Redirect(w, r, "/edit", http.StatusSeeOther)
}

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

// create data
func (contactHandler *ContactHandler) CreateContact(w http.ResponseWriter, r *http.Request) {
	
	if err := r.ParseForm(); err != nil {
		http.Error(w, "failed process form", http.StatusBadRequest)
		return
	}

	nameStr := r.FormValue("contact_name")
	typeStr := r.FormValue("contact_type")
	linkStr := r.FormValue("contact_link")

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

	http.Redirect(w, r, "/edit#contact", http.StatusSeeOther)
}

// update data
func (contactHandler *ContactHandler) UpdateContact(w http.ResponseWriter, r *http.Request) {

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

	http.Redirect(w, r, "/edit#contact", http.StatusSeeOther)
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

	http.Redirect(w, r, "/edit#contact", http.StatusSeeOther)
}

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

type WorkHandler struct {
	WorkService service.WorkService
	Templates   *template.Template
}

func NewWorkHandler(templates *template.Template, workService service.WorkService) WorkHandler {
	return WorkHandler{
		WorkService: workService,
		Templates:   templates,
	}
}

// create data
func (workHandler *WorkHandler) CreateWork(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		http.Error(w, "failed process form", http.StatusBadRequest)
		return
	}

	nameStr := r.FormValue("work_name")
	descStr := r.FormValue("work_description")
	yearStr := r.FormValue("work_year")

	yearInt, err := strconv.Atoi(yearStr)
	if err != nil {
		http.Error(w, "year must be integer", http.StatusBadRequest)
		return
	}

	// parsing to model work
	work := dto.WorkCreateRequest{
		Name:        nameStr,
		Description: descStr,
		Year:        yearInt,
	}

	err = workHandler.WorkService.CreateWork(&work)
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusInternalServerError, "error:"+err.Error(), nil)
		return
	}

	http.Redirect(w, r, "/edit#work", http.StatusSeeOther)
}

// update data
func (workHandler *WorkHandler) UpdateWork(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		http.Error(w, "failed process form", http.StatusBadRequest)
		return
	}

	idStr := r.FormValue("work_id")
	nameStr := r.FormValue("work_name")
	descStr := r.FormValue("work_description")
	yearStr := r.FormValue("work_year")

	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "id must be integer", http.StatusBadRequest)
		return
	}

	yearInt, err := strconv.Atoi(yearStr)
	if err != nil {
		http.Error(w, "year must be integer", http.StatusBadRequest)
		return
	}

	// parsing to model work
	work := dto.WorkUpdateRequest{
		Name:        &nameStr,
		Description: &descStr,
		Year:        &yearInt,
	}

	err = workHandler.WorkService.UpdateWork(idInt, &work)
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusInternalServerError, "error update:"+err.Error(), nil)
		return
	}

	http.Redirect(w, r, "/edit#work", http.StatusSeeOther)
}

// delete data
func (workHandler *WorkHandler) DeleteWork(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		http.Error(w, "failed process form", http.StatusBadRequest)
		return
	}

	idStr := r.FormValue("work_id")

	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "id must be integer", http.StatusBadRequest)
		return
	}

	err = workHandler.WorkService.DeleteWork(idInt)
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusInternalServerError, "error delete:"+err.Error(), nil)
		return
	}

	http.Redirect(w, r, "/edit#work", http.StatusSeeOther)
}

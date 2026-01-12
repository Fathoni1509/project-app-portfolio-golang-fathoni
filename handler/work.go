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

// abaikan get data, get data lakukan di menu
// get data work
func (workHandler *WorkHandler) GetDataWork(w http.ResponseWriter, r *http.Request) {
	work, err := workHandler.WorkService.GetDataWork()
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusNotFound, "data not found:"+err.Error(), nil)
		return
	}

	// utils.ResponseSuccess(w, http.StatusOK, "success get data work", work)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := workHandler.Templates.ExecuteTemplate(w, "work", work); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// create data
func (workHandler *WorkHandler) CreateWork(w http.ResponseWriter, r *http.Request) {
	// var req dto.WorkUpdateRequest
	// if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
	// 	utils.ResponseBadRequest(w, http.StatusBadRequest, "error data:"+err.Error(), nil)
	// 	return
	// }

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

	// if err := r.FormValue("work"); err != nil {
	// 	utils.ResponseBadRequest(w, http.StatusBadRequest, "error data:"+err.Error(), nil)
	// 	return
	// }

	// existing, err := workHandler.WorkService.GetDataWork()
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

	// utils.ResponseSuccess(w, http.StatusOK, "updated success", nil)
	// w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// if err := workHandler.Templates.ExecuteTemplate(w, "work_edit", work); err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// }

	http.Redirect(w, r, "/edit#work", http.StatusSeeOther)
}

// update data
func (workHandler *WorkHandler) UpdateWork(w http.ResponseWriter, r *http.Request) {
	// var req dto.WorkUpdateRequest
	// if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
	// 	utils.ResponseBadRequest(w, http.StatusBadRequest, "error data:"+err.Error(), nil)
	// 	return
	// }

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

	// if err := r.FormValue("work"); err != nil {
	// 	utils.ResponseBadRequest(w, http.StatusBadRequest, "error data:"+err.Error(), nil)
	// 	return
	// }

	// existing, err := workHandler.WorkService.GetDataWork()
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

	// utils.ResponseSuccess(w, http.StatusOK, "updated success", nil)
	// w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// if err := workHandler.Templates.ExecuteTemplate(w, "work_edit", work); err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// }

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

	// utils.ResponseSuccess(w, http.StatusOK, "updated success", nil)
	// w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// if err := workHandler.Templates.ExecuteTemplate(w, "work_edit", work); err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// }

	http.Redirect(w, r, "/edit#work", http.StatusSeeOther)
}

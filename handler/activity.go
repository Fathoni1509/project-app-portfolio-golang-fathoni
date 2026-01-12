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

type ActivityHandler struct {
	ActivityService service.ActivityService
	Templates *template.Template
}

func NewActivityHandler(templates *template.Template,activityService service.ActivityService) ActivityHandler {
	return ActivityHandler{
		ActivityService: activityService,
		Templates: templates,
	}
}

// abaikan get data, get data lakukan di menu
// get data activity
func (activityHandler *ActivityHandler) GetDataActivity(w http.ResponseWriter, r *http.Request) {
	activity, err := activityHandler.ActivityService.GetDataActivity()
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusNotFound, "data not found:"+err.Error(), nil)
		return
	}

	// utils.ResponseSuccess(w, http.StatusOK, "success get data activity", activity)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := activityHandler.Templates.ExecuteTemplate(w, "activity", activity); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// create data
func (activityHandler *ActivityHandler) CreateActivity(w http.ResponseWriter, r *http.Request) {
	// var req dto.ActivityUpdateRequest
	// if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
	// 	utils.ResponseBadRequest(w, http.StatusBadRequest, "error data:"+err.Error(), nil)
	// 	return
	// }

	if err := r.ParseForm(); err != nil {
		http.Error(w, "failed process form", http.StatusBadRequest)
		return
	}

	nameStr := r.FormValue("activity_name") 
    yearStr  := r.FormValue("activity_year")
    catIdStr := r.FormValue("activity_category_id")

	yearInt, err := strconv.Atoi(yearStr)
    if err != nil {
        http.Error(w, "year must be integer", http.StatusBadRequest)
        return
    }

	catIdInt, err := strconv.Atoi(catIdStr)
    if err != nil {
        http.Error(w, "category id must be integer", http.StatusBadRequest)
        return
    }

	// if err := r.FormValue("activity"); err != nil {
	// 	utils.ResponseBadRequest(w, http.StatusBadRequest, "error data:"+err.Error(), nil)
	// 	return
	// }

	// existing, err := activityHandler.ActivityService.GetDataActivity()
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

	// parsing to model activity
	activity := dto.ActivityCreateRequest{
		Name: nameStr,
		Year: yearInt,
		CategoryId: catIdInt,
	}

	err = activityHandler.ActivityService.CreateActivity(&activity)
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusInternalServerError, "error:"+err.Error(), nil)
		return
	}

	// utils.ResponseSuccess(w, http.StatusOK, "updated success", nil)
	// w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// if err := activityHandler.Templates.ExecuteTemplate(w, "activity_edit", activity); err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// }

	http.Redirect(w, r, "/edit#activity", http.StatusSeeOther)
}

// update data
func (activityHandler *ActivityHandler) UpdateActivity(w http.ResponseWriter, r *http.Request) {
	// var req dto.ActivityUpdateRequest
	// if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
	// 	utils.ResponseBadRequest(w, http.StatusBadRequest, "error data:"+err.Error(), nil)
	// 	return
	// }

	if err := r.ParseForm(); err != nil {
		http.Error(w, "failed process form", http.StatusBadRequest)
		return
	}

	idStr := r.FormValue("activity_id")
	nameStr := r.FormValue("activity_name") 
    yearStr  := r.FormValue("activity_year")
    catIdStr := r.FormValue("activity_category_id")

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

	catIdInt, err := strconv.Atoi(catIdStr)
    if err != nil {
        http.Error(w, "category id must be integer", http.StatusBadRequest)
        return
    }

	// if err := r.FormValue("activity"); err != nil {
	// 	utils.ResponseBadRequest(w, http.StatusBadRequest, "error data:"+err.Error(), nil)
	// 	return
	// }

	// existing, err := activityHandler.ActivityService.GetDataActivity()
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

	// parsing to model activity
	activity := dto.ActivityUpdateRequest{
		Name: &nameStr,
		Year: &yearInt,
		CategoryId: &catIdInt,
	}

	err = activityHandler.ActivityService.UpdateActivity(idInt, &activity)
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusInternalServerError, "error update:"+err.Error(), nil)
		return
	}

	// utils.ResponseSuccess(w, http.StatusOK, "updated success", nil)
	// w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// if err := activityHandler.Templates.ExecuteTemplate(w, "activity_edit", activity); err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// }

	http.Redirect(w, r, "/edit#activity", http.StatusSeeOther)
}

// delete data
func (activityHandler *ActivityHandler) DeleteActivity(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		http.Error(w, "failed process form", http.StatusBadRequest)
		return
	}

	idStr := r.FormValue("activity_id")

	idInt, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "id must be integer", http.StatusBadRequest)
        return
    }



	err = activityHandler.ActivityService.DeleteActivity(idInt)
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusInternalServerError, "error delete:"+err.Error(), nil)
		return
	}

	// utils.ResponseSuccess(w, http.StatusOK, "updated success", nil)
	// w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// if err := activityHandler.Templates.ExecuteTemplate(w, "activity_edit", activity); err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// }

	http.Redirect(w, r, "/edit#activity", http.StatusSeeOther)
}
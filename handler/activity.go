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

// create data
func (activityHandler *ActivityHandler) CreateActivity(w http.ResponseWriter, r *http.Request) {
	
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

	http.Redirect(w, r, "/edit#activity", http.StatusSeeOther)
}

// update data
func (activityHandler *ActivityHandler) UpdateActivity(w http.ResponseWriter, r *http.Request) {

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

	http.Redirect(w, r, "/edit#activity", http.StatusSeeOther)
}
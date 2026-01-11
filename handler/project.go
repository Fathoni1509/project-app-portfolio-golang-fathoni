package handler

import (
	// "encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"project-app-portfolio-golang-fathoni/dto"
	"project-app-portfolio-golang-fathoni/service"
	"project-app-portfolio-golang-fathoni/utils"
	"strconv"
)

type ProjectHandler struct {
	ProjectService service.ProjectService
	Templates      *template.Template
}

func NewProjectHandler(templates *template.Template, projectService service.ProjectService) ProjectHandler {
	return ProjectHandler{
		ProjectService: projectService,
		Templates:      templates,
	}
}

// abaikan get data, get data lakukan di menu
// get data project
func (projectHandler *ProjectHandler) GetDataProject(w http.ResponseWriter, r *http.Request) {
	project, err := projectHandler.ProjectService.GetDataProject()
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusNotFound, "data not found:"+err.Error(), nil)
		return
	}

	// utils.ResponseSuccess(w, http.StatusOK, "success get data project", project)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := projectHandler.Templates.ExecuteTemplate(w, "project", project); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// create data
func (projectHandler *ProjectHandler) CreateProject(w http.ResponseWriter, r *http.Request) {
	// var req dto.ProjectUpdateRequest
	// if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
	// 	utils.ResponseBadRequest(w, http.StatusBadRequest, "error data:"+err.Error(), nil)
	// 	return
	// }

	// if err := r.ParseForm(); err != nil {
	// 	http.Error(w, "failed process form", http.StatusBadRequest)
	// 	return
	// }

	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "file too large", http.StatusBadRequest)
		return
	}

	nameStr := r.FormValue("project_name")
	descStr := r.FormValue("project_description")
	yearStr := r.FormValue("project_year")
	linkStr := r.FormValue("project_link")
	// imgStr := r.FormValue("project_image_data")

	yearInt, err := strconv.Atoi(yearStr)
	if err != nil {
		http.Error(w, "year must be integer", http.StatusBadRequest)
		return
	}

	var imagePath string
	file, header, err := r.FormFile("project_image_data")

	if err == nil {
		defer file.Close()

		filename := fmt.Sprintf("project-%s", header.Filename)

		uploadDir := "./public/uploads"
		os.MkdirAll(uploadDir, os.ModePerm)

		filePath := filepath.Join(uploadDir, filename)

		dst, err := os.Create(filePath)
		if err != nil {
			http.Error(w, "error save file", http.StatusInternalServerError)
			return
		}
		defer dst.Close()

		if _, err := io.Copy(dst, file); err != nil {
			http.Error(w, "error copy file", http.StatusInternalServerError)
			return
		}

		imagePath = "/public/uploads/" + filename
	} else {
		utils.ResponseBadRequest(w, http.StatusInternalServerError, "error:"+err.Error(), nil)
		return
	}

	// if err := r.FormValue("project"); err != nil {
	// 	utils.ResponseBadRequest(w, http.StatusBadRequest, "error data:"+err.Error(), nil)
	// 	return
	// }

	// existing, err := projectHandler.ProjectService.GetDataProject()
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

	// parsing to model project
	project := dto.ProjectCreateRequest{
		Name:        nameStr,
		Description: descStr,
		Year:        yearInt,
		Link:        linkStr,
		ImageData:   imagePath,
	}

	err = projectHandler.ProjectService.CreateProject(&project)
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusInternalServerError, "error:"+err.Error(), nil)
		return
	}

	// utils.ResponseSuccess(w, http.StatusOK, "updated success", nil)
	// w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// if err := projectHandler.Templates.ExecuteTemplate(w, "project_edit", project); err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// }

	http.Redirect(w, r, "/edit", http.StatusSeeOther)
}

// update data
func (projectHandler *ProjectHandler) UpdateProject(w http.ResponseWriter, r *http.Request) {
	// var req dto.ProjectUpdateRequest
	// if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
	// 	utils.ResponseBadRequest(w, http.StatusBadRequest, "error data:"+err.Error(), nil)
	// 	return
	// }

	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "file too large", http.StatusBadRequest)
		return
	}

	idStr := r.FormValue("project_id")
	nameStr := r.FormValue("project_name")
	descStr := r.FormValue("project_description")
	yearStr := r.FormValue("project_year")
	linkStr := r.FormValue("project_link")

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

	var imagePath string
	file, header, err := r.FormFile("project_image_data")

	if err == nil {
		defer file.Close()

		filename := fmt.Sprintf("project-%s", header.Filename)

		uploadDir := "./public/uploads"
		os.MkdirAll(uploadDir, os.ModePerm)

		filePath := filepath.Join(uploadDir, filename)

		dst, err := os.Create(filePath)
		if err != nil {
			http.Error(w, "error save file", http.StatusInternalServerError)
			return
		}
		defer dst.Close()

		if _, err := io.Copy(dst, file); err != nil {
			http.Error(w, "error copy file", http.StatusInternalServerError)
			return
		}

		imagePath = "/public/uploads/" + filename
	} 
	// else {
	// 	utils.ResponseBadRequest(w, http.StatusInternalServerError, "error:"+err.Error(), nil)
	// 	return
	// }

	// if err := r.FormValue("project"); err != nil {
	// 	utils.ResponseBadRequest(w, http.StatusBadRequest, "error data:"+err.Error(), nil)
	// 	return
	// }

	// existing, err := projectHandler.ProjectService.GetDataProject()
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

	// parsing to model project
	project := dto.ProjectUpdateRequest{
		Name:        &nameStr,
		Description: &descStr,
		Year:        &yearInt,
		Link:        &linkStr,
		// ImageData:   &imagePath,
	}

	if imagePath != "" {
        project.ImageData = &imagePath
    }

	err = projectHandler.ProjectService.UpdateProject(idInt, &project)
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusInternalServerError, "error update:"+err.Error(), nil)
		return
	}

	// utils.ResponseSuccess(w, http.StatusOK, "updated success", nil)
	// w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// if err := projectHandler.Templates.ExecuteTemplate(w, "project_edit", project); err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// }

	http.Redirect(w, r, "/edit", http.StatusSeeOther)
}

// delete data
func (projectHandler *ProjectHandler) DeleteProject(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "file too large", http.StatusBadRequest)
		return
	}

	idStr := r.FormValue("project_id")

	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "id must be integer", http.StatusBadRequest)
		return
	}

	err = projectHandler.ProjectService.DeleteProject(idInt)
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusInternalServerError, "error delete:"+err.Error(), nil)
		return
	}

	// utils.ResponseSuccess(w, http.StatusOK, "updated success", nil)
	// w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// if err := projectHandler.Templates.ExecuteTemplate(w, "project_edit", project); err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// }

	http.Redirect(w, r, "/edit", http.StatusSeeOther)
}

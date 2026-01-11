package service

import (
	"project-app-portfolio-golang-fathoni/dto"
	"project-app-portfolio-golang-fathoni/repository"
)

type ProjectService interface {
	GetDataProject() ([]dto.ProjectResponse, error)
	CreateProject(project *dto.ProjectCreateRequest) error
	UpdateProject(project_id int, project *dto.ProjectUpdateRequest) error
	DeleteProject(project_id int) error
}

type projectService struct {
	Repo repository.Repository
}

func NewProjectService(repo repository.Repository) ProjectService {
	return &projectService{Repo: repo}
}

// service get data project
func (srv *projectService) GetDataProject() ([]dto.ProjectResponse, error) {
	return srv.Repo.ProjectRepo.GetDataProject()
}

// service create data
func (srv *projectService) CreateProject(project *dto.ProjectCreateRequest) error {
	return srv.Repo.ProjectRepo.CreateProject(project)
}

// service update data
func (srv *projectService) UpdateProject(project_id int, project *dto.ProjectUpdateRequest) error {
	return srv.Repo.ProjectRepo.UpdateProject(project_id, project)
}

// service delete data
func (srv *projectService) DeleteProject(project_id int) error {
	return srv.Repo.ProjectRepo.DeleteProject(project_id)
}

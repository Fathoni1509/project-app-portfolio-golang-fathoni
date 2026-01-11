package service

import (
	"project-app-portfolio-golang-fathoni/dto"
	"project-app-portfolio-golang-fathoni/repository"
)

type WorkService interface {
	GetDataWork() ([]dto.WorkResponse, error)
	CreateWork(work *dto.WorkCreateRequest) error
	UpdateWork(work_id int, work *dto.WorkUpdateRequest) error
	DeleteWork(work_id int) error
}

type workService struct {
	Repo repository.Repository
}

func NewWorkService(repo repository.Repository) WorkService {
	return &workService{Repo: repo}
}

// service get data work
func (srv *workService) GetDataWork() ([]dto.WorkResponse, error) {
	return srv.Repo.WorkRepo.GetDataWork()
}

// service create data
func (srv *workService) CreateWork(work *dto.WorkCreateRequest) error {
	return srv.Repo.WorkRepo.CreateWork(work)
}

// service update data
func (srv *workService) UpdateWork(work_id int, work *dto.WorkUpdateRequest) error {
	return srv.Repo.WorkRepo.UpdateWork(work_id, work)
}

// service delete data
func (srv *workService) DeleteWork(work_id int) error {
	return srv.Repo.WorkRepo.DeleteWork(work_id)
}

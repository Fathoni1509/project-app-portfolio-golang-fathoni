package service

import (
	"project-app-portfolio-golang-fathoni/dto"
	"project-app-portfolio-golang-fathoni/repository"
)

type ActivityService interface {
	GetDataActivity() ([]dto.ActivityResponse, error)
	CreateActivity(activity *dto.ActivityCreateRequest) error
	UpdateActivity(activity_id int,activity *dto.ActivityUpdateRequest) error
	DeleteActivity(activity_id int) error
}

type activityService struct {
	Repo repository.Repository
}

func NewActivityService(repo repository.Repository) ActivityService {
	return &activityService{Repo: repo}
}

// service get data activity
func (srv *activityService) GetDataActivity() ([]dto.ActivityResponse, error) {
	return srv.Repo.ActivityRepo.GetDataActivity()
}

// service create data
func (srv *activityService) CreateActivity(activity *dto.ActivityCreateRequest) error {
	return srv.Repo.ActivityRepo.CreateActivity(activity)
}

// service update data
func (srv *activityService) UpdateActivity(activity_id int,activity *dto.ActivityUpdateRequest) error {
	return srv.Repo.ActivityRepo.UpdateActivity(activity_id, activity)
}

// service delete data
func (srv *activityService) DeleteActivity(activity_id int) error {
	return srv.Repo.ActivityRepo.DeleteActivity(activity_id)
}
package service

import (
	"project-app-portfolio-golang-fathoni/dto"
	"project-app-portfolio-golang-fathoni/repository"
)

type PersonalService interface {
	GetDataPersonal() (dto.PersonalResponse, error)
	UpdatePersonal(personal *dto.PersonalUpdateRequest) error
}

type personalService struct {
	Repo repository.Repository
}

func NewPersonalService(repo repository.Repository) PersonalService {
	return &personalService{Repo: repo}
}

// service get data personal
func (srv *personalService) GetDataPersonal() (dto.PersonalResponse, error) {
	return srv.Repo.PersonalRepo.GetDataPersonal()
}

// service update data
func (srv *personalService) UpdatePersonal(personal *dto.PersonalUpdateRequest) error {
	return srv.Repo.PersonalRepo.UpdatePersonal(personal)
}
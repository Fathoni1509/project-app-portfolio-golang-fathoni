package service

import (
	"project-app-portfolio-golang-fathoni/dto"
	"project-app-portfolio-golang-fathoni/repository"
)

type ContactService interface {
	GetDataContact() ([]dto.ContactResponse, error)
	CreateContact(contact *dto.ContactCreateRequest) error
	UpdateContact(contact_id int, contact *dto.ContactUpdateRequest) error
	DeleteContact(contact_id int) error
}

type contactService struct {
	Repo repository.Repository
}

func NewContactService(repo repository.Repository) ContactService {
	return &contactService{Repo: repo}
}

// service get data contact
func (srv *contactService) GetDataContact() ([]dto.ContactResponse, error) {
	return srv.Repo.ContactRepo.GetDataContact()
}

// service create data
func (srv *contactService) CreateContact(contact *dto.ContactCreateRequest) error {
	return srv.Repo.ContactRepo.CreateContact(contact)
}

// service update data
func (srv *contactService) UpdateContact(contact_id int, contact *dto.ContactUpdateRequest) error {
	return srv.Repo.ContactRepo.UpdateContact(contact_id, contact)
}

// service delete data
func (srv *contactService) DeleteContact(contact_id int) error {
	return srv.Repo.ContactRepo.DeleteContact(contact_id)
}

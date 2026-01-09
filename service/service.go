package service

import "project-app-portfolio-golang-fathoni/repository"

type Service struct {
	PersonalService PersonalService
}

func NewService(repo repository.Repository) Service {
	return Service{
		PersonalService: NewPersonalService(repo),
	}
}
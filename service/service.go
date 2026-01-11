package service

import "project-app-portfolio-golang-fathoni/repository"

type Service struct {
	PersonalService PersonalService
	ActivityService ActivityService
	WorkService     WorkService
	ProjectService  ProjectService
	ContactService  ContactService
}

func NewService(repo repository.Repository) Service {
	return Service{
		PersonalService: NewPersonalService(repo),
		ActivityService: NewActivityService(repo),
		WorkService:     NewWorkService(repo),
		ProjectService:  NewProjectService(repo),
		ContactService:  NewContactService(repo),
	}
}

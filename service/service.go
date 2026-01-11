package service

import "project-app-portfolio-golang-fathoni/repository"

type Service struct {
	PersonalService PersonalService
	ActivityService ActivityService
	WorkService     WorkService
}

func NewService(repo repository.Repository) Service {
	return Service{
		PersonalService: NewPersonalService(repo),
		ActivityService: NewActivityService(repo),
		WorkService:     NewWorkService(repo),
	}
}

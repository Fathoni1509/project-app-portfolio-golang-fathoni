package repository

import "project-app-portfolio-golang-fathoni/database"

type Repository struct {
	PersonalRepo PersonalRepository
	ActivityRepo ActivityRepository
	WorkRepo     WorkRepository
	ProjectRepo  ProjectRepository
}

func NewRepository(db database.PgxIface) Repository {
	return Repository{
		PersonalRepo: NewPersonalRepository(db),
		ActivityRepo: NewActivityRepository(db),
		WorkRepo:     NewWorkRepository(db),
		ProjectRepo:  NewProjectRepository(db),
	}
}

package repository

import "project-app-portfolio-golang-fathoni/database"

type Repository struct {
	PersonalRepo PersonalRepository
	ActivityRepo ActivityRepository
}

func NewRepository(db database.PgxIface) Repository {
	return Repository{
		PersonalRepo: NewPersonalRepository(db),
		ActivityRepo: NewActivityRepository(db),
	}
}
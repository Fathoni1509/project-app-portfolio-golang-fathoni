package repository

import (
	"context"
	"errors"
	"project-app-portfolio-golang-fathoni/database"
	"project-app-portfolio-golang-fathoni/dto"
)

type ProjectRepository interface {
	GetDataProject() ([]dto.ProjectResponse, error)
	CreateProject(project *dto.ProjectCreateRequest) error
	UpdateProject(project_id int, project *dto.ProjectUpdateRequest) error
	DeleteProject(project_id int) error
}

type projectRepository struct {
	db database.PgxIface
}

func NewProjectRepository(db database.PgxIface) ProjectRepository {
	return &projectRepository{db: db}
}

// get data project
func (repo *projectRepository) GetDataProject() ([]dto.ProjectResponse, error) {
	query := `SELECT
		project_id,
		name,
		description,
		year,
		link,
		image_data
	FROM project
	WHERE deleted_at IS NULL
	ORDER BY project_id ASC`

	rows, err := repo.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var project []dto.ProjectResponse
	var list dto.ProjectResponse

	for rows.Next() {
		err := rows.Scan(&list.ProjectId, &list.Name, &list.Description, &list.Year, &list.Link, &list.ImageData)
		if err != nil {
			return nil, err
		}
		project = append(project, list)
	}

	return project, err
}

// create data project
func (repo *projectRepository) CreateProject(project *dto.ProjectCreateRequest) error {
	query := `INSERT INTO project (name, description, year, link, image_data, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, NOW(), NOW()) RETURNING project_id`

	_, err := repo.db.Exec(context.Background(), query,
		project.Name,
		project.Description,
		project.Year,
		project.Link,
		project.ImageData,
	)

	if err != nil {
		return err
	}

	return nil
}

// update data project
func (repo *projectRepository) UpdateProject(project_id int, project *dto.ProjectUpdateRequest) error {
	query := `UPDATE project
		SET name=$1, description=$2, year=$3, link=$4, image_data= COALESCE($5, image_data),
		updated_at=NOW()
		WHERE deleted_at IS NULL AND project_id=$6`

	commandTag, err := repo.db.Exec(context.Background(), query,
		project.Name,
		project.Description,
		project.Year,
		project.Link,
		project.ImageData,
		project_id,
	)

	if err != nil {
		return err
	}

	if commandTag.RowsAffected() == 0 {
		return errors.New("data project not found")
	}

	return nil
}

// delete data project
func (repo *projectRepository) DeleteProject(project_id int) error {
	query := `UPDATE project
		SET deleted_at=NOW()
		WHERE deleted_at IS NULL AND project_id=$1`

	commandTag, err := repo.db.Exec(context.Background(), query,
		project_id,
	)

	if err != nil {
		return err
	}

	if commandTag.RowsAffected() == 0 {
		return errors.New("data project not found")
	}

	return nil
}
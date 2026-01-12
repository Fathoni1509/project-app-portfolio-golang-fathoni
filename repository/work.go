package repository

import (
	"context"
	"errors"
	"project-app-portfolio-golang-fathoni/database"
	"project-app-portfolio-golang-fathoni/dto"
)

type WorkRepository interface {
	GetDataWork() ([]dto.WorkResponse, error)
	CreateWork(work *dto.WorkCreateRequest) error
	UpdateWork(work_id int, work *dto.WorkUpdateRequest) error
	DeleteWork(work_id int) error
}

type workRepository struct {
	db database.PgxIface
}

func NewWorkRepository(db database.PgxIface) WorkRepository {
	return &workRepository{db: db}
}

// get data work
func (repo *workRepository) GetDataWork() ([]dto.WorkResponse, error) {
	query := `SELECT
		work_id,
		name,
		description,
		year
	FROM work
	WHERE deleted_at IS NULL
	ORDER BY year DESC`

	rows, err := repo.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var work []dto.WorkResponse
	var list dto.WorkResponse

	for rows.Next() {
		err := rows.Scan(&list.WorkId, &list.Name, &list.Description, &list.Year)
		if err != nil {
			return nil, err
		}
		work = append(work, list)
	}

	return work, err
}

// create data work
func (repo *workRepository) CreateWork(work *dto.WorkCreateRequest) error {
	query := `INSERT INTO work (name, description, year, created_at, updated_at)
		VALUES ($1, $2, $3, NOW(), NOW()) RETURNING work_id`

	_, err := repo.db.Exec(context.Background(), query,
		work.Name,
		work.Description,
		work.Year,
	)

	if err != nil {
		return err
	}

	return nil
}

// update data work
func (repo *workRepository) UpdateWork(work_id int, work *dto.WorkUpdateRequest) error {
	query := `UPDATE work
		SET name=$1, description=$2, year=$3, updated_at=NOW()
		WHERE deleted_at IS NULL AND work_id=$4`

	commandTag, err := repo.db.Exec(context.Background(), query,
		work.Name,
		work.Description,
		work.Year,
		work_id,
	)

	if err != nil {
		return err
	}

	if commandTag.RowsAffected() == 0 {
		return errors.New("data work not found")
	}

	return nil
}

// delete data work
func (repo *workRepository) DeleteWork(work_id int) error {
	query := `UPDATE work
		SET deleted_at=NOW()
		WHERE deleted_at IS NULL AND work_id=$1`

	commandTag, err := repo.db.Exec(context.Background(), query,
		work_id,
	)

	if err != nil {
		return err
	}

	if commandTag.RowsAffected() == 0 {
		return errors.New("data work not found")
	}

	return nil
}
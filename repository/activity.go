package repository

import (
	"context"
	"errors"
	"project-app-portfolio-golang-fathoni/database"
	"project-app-portfolio-golang-fathoni/dto"
)

type ActivityRepository interface {
	GetDataActivity() ([]dto.ActivityResponse, error)
	CreateActivity(activity *dto.ActivityCreateRequest) error
	UpdateActivity(activity_id int, activity *dto.ActivityUpdateRequest) error
	DeleteActivity(activity_id int) error
}

type activityRepository struct {
	db database.PgxIface
}

func NewActivityRepository(db database.PgxIface) ActivityRepository {
	return &activityRepository{db: db}
}

// get data activity
func (repo *activityRepository) GetDataActivity() ([]dto.ActivityResponse, error) {
	query := `SELECT
		a.activity_id,
		a.name,
		a.year,
		a.category_id,
		ac.name
	FROM activity a
	JOIN activity_category ac ON a.category_id = ac.activity_category_id
	WHERE a.deleted_at IS NULL
	ORDER BY a.activity_id ASC`

	rows, err := repo.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var activity []dto.ActivityResponse
	var list dto.ActivityResponse

	for rows.Next() {
		err := rows.Scan(&list.ActivityId, &list.Name, &list.Year, &list.CategoryId, &list.Category)
		if err != nil {
			return nil, err
		}
		activity = append(activity, list)
	}

	return activity, err
}

// create data activity
func (repo *activityRepository) CreateActivity(activity *dto.ActivityCreateRequest) error {
	query := `INSERT INTO activity (name, year, category_id, created_at, updated_at)
		VALUES ($1, $2, $3, NOW(), NOW()) RETURNING activity_id`

	_, err := repo.db.Exec(context.Background(), query,
		activity.Name,
		activity.Year,
		activity.CategoryId,
	)

	if err != nil {
		return err
	}

	return nil
}

// update data activity
func (repo *activityRepository) UpdateActivity(activity_id int, activity *dto.ActivityUpdateRequest) error {
	query := `UPDATE activity
		SET name=$1, year=$2, category_id=$3, updated_at=NOW()
		WHERE deleted_at IS NULL AND activity_id=$4`

	commandTag, err := repo.db.Exec(context.Background(), query,
		activity.Name,
		activity.Year,
		activity.CategoryId,
		activity_id,
	)

	if err != nil {
		return err
	}

	if commandTag.RowsAffected() == 0 {
		return errors.New("data activity not found")
	}

	return nil
}

// delete data activity
func (repo *activityRepository) DeleteActivity(activity_id int) error {
	query := `UPDATE activity
		SET deleted_at=NOW()
		WHERE deleted_at IS NULL AND activity_id=$1`

	commandTag, err := repo.db.Exec(context.Background(), query,
		activity_id,
	)

	if err != nil {
		return err
	}

	if commandTag.RowsAffected() == 0 {
		return errors.New("data activity not found")
	}

	return nil
}
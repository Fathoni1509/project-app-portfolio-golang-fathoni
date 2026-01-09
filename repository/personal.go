package repository

import (
	"context"
	"errors"
	"project-app-portfolio-golang-fathoni/database"
	"project-app-portfolio-golang-fathoni/dto"
)

type PersonalRepository interface {
	GetDataPersonal() (dto.PersonalResponse, error)
	UpdatePersonal(personal *dto.PersonalUpdateRequest) error
}

type personalRepository struct {
	db database.PgxIface
}

func NewPersonalRepository(db database.PgxIface) PersonalRepository {
	return &personalRepository{db: db}
}

// get data personal
func (repo *personalRepository) GetDataPersonal() (dto.PersonalResponse, error) {
	query := `SELECT
		personal_id,
		name,
		age,
		description
	FROM personal
	WHERE deleted_at IS NULL`

	var personal dto.PersonalResponse

	err := repo.db.QueryRow(context.Background(), query).Scan(&personal.PersonalId, &personal.Name, &personal.Age, &personal.Description)

	if err != nil {
		return dto.PersonalResponse{}, err
	}

	return personal, err
}

// update data personal
func (repo *personalRepository) UpdatePersonal(personal *dto.PersonalUpdateRequest) error {
	query := `UPDATE personal
		SET name=$1, age=$2, description=$3, updated_at=NOW()
		WHERE deleted_at IS NULL AND personal_id=1`

	commandTag, err := repo.db.Exec(context.Background(), query,
		personal.Name,
		personal.Age,
		personal.Description,
	)

	if err != nil {
		return err
	}

	if commandTag.RowsAffected() == 0 {
		return errors.New("data personal not found")
	}

	return nil
}
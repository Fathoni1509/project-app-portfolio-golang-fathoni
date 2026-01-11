package repository

import (
	"context"
	"errors"
	"project-app-portfolio-golang-fathoni/database"
	"project-app-portfolio-golang-fathoni/dto"
)

type ContactRepository interface {
	GetDataContact() ([]dto.ContactResponse, error)
	CreateContact(contact *dto.ContactCreateRequest) error
	UpdateContact(contact_id int, contact *dto.ContactUpdateRequest) error
	DeleteContact(contact_id int) error
}

type contactRepository struct {
	db database.PgxIface
}

func NewContactRepository(db database.PgxIface) ContactRepository {
	return &contactRepository{db: db}
}

// get data contact
func (repo *contactRepository) GetDataContact() ([]dto.ContactResponse, error) {
	query := `SELECT
		contact_id,
		name,
		type,
		link
	FROM contact
	WHERE deleted_at IS NULL
	ORDER BY contact_id ASC`

	rows, err := repo.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var contact []dto.ContactResponse
	var list dto.ContactResponse

	for rows.Next() {
		err := rows.Scan(&list.ContactId, &list.Name, &list.Type, &list.Link)
		if err != nil {
			return nil, err
		}
		contact = append(contact, list)
	}

	return contact, err
}

// create data contact
func (repo *contactRepository) CreateContact(contact *dto.ContactCreateRequest) error {
	query := `INSERT INTO contact (name, type, link, created_at, updated_at)
		VALUES ($1, $2, $3, NOW(), NOW()) RETURNING contact_id`

	_, err := repo.db.Exec(context.Background(), query,
		contact.Name,
		contact.Type,
		contact.Link,
	)

	if err != nil {
		return err
	}

	return nil
}

// update data contact
func (repo *contactRepository) UpdateContact(contact_id int, contact *dto.ContactUpdateRequest) error {
	query := `UPDATE contact
		SET name=$1, type=$2, link=$3, updated_at=NOW()
		WHERE deleted_at IS NULL AND contact_id=$4`

	commandTag, err := repo.db.Exec(context.Background(), query,
		contact.Name,
		contact.Type,
		contact.Link,
		contact_id,
	)

	if err != nil {
		return err
	}

	if commandTag.RowsAffected() == 0 {
		return errors.New("data contact not found")
	}

	return nil
}

// delete data contact
func (repo *contactRepository) DeleteContact(contact_id int) error {
	query := `UPDATE contact
		SET deleted_at=NOW()
		WHERE deleted_at IS NULL AND contact_id=$1`

	commandTag, err := repo.db.Exec(context.Background(), query,
		contact_id,
	)

	if err != nil {
		return err
	}

	if commandTag.RowsAffected() == 0 {
		return errors.New("data contact not found")
	}

	return nil
}
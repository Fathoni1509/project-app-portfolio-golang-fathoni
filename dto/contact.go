package dto

type ContactCreateRequest struct {
	Name      string `json:"name" validate:"required,min=3"`
	Type      string `json:"type" validate:"required"`
	Link      string `json:"link" validate:"required"`
}

type ContactUpdateRequest struct {
	Name      *string `json:"name"`
	Type      *string `json:"type"`
	Link      *string `json:"link"`
}

type ContactResponse struct {
	ContactId int    `json:"contact_id"`
	Name      string `json:"name"`
	Type      string `json:"type"`
	Link      string `json:"link"`
}

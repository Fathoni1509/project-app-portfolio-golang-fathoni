package dto

type ProjectCreateRequest struct {
	Name        string `json:"name" validate:"required,min=3"`
	Description string `json:"description" validate:"required,min=3"`
	Year        int    `json:"year" validate:"required,min=1000,max=2999"`
	Link        string `json:"link" validate:"required"`
	ImageData   byte   `json:"image_data" validate:"required"`
}

type ProjectUpdateRequest struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
	Year        *int    `json:"year"`
	Link        *string `json:"link"`
	ImageData   *byte   `json:"image_data"`
}

type ProjectResponse struct {
	ProjectId   int    `json:"project_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Year        int    `json:"year"`
	Link        string `json:"link"`
	ImageData   byte   `json:"image_data"`
}

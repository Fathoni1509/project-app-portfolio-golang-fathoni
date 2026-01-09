package dto

type WorkCreateRequest struct {
	Name        string `json:"name" validate:"required,min=3"`
	Description string `json:"description" validate:"required,min=3"`
	Year        int    `json:"year" validate:"required,min=1000,max=2999"`
}

type WorkUpdateRequest struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
	Year        *int    `json:"year"`
}

type WorkResponse struct {
	WorkId      int    `json:"work_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Year        int    `json:"year"`
}

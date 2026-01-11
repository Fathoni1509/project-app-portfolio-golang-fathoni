package dto

type ActivityCreateRequest struct {
	Name       string `json:"name" validate:"required,min=3"`
	Year       int    `json:"year" validate:"required,min=1000,max=2999"`
	CategoryId int    `json:"category_id" validate:"required,gt=0"`
}

type ActivityUpdateRequest struct {
	Name       *string `json:"name"`
	Year       *int    `json:"year"`
	CategoryId *int    `json:"category_id"`
}

type ActivityResponse struct {
	ActivityId int    `json:"activity_id"`
	Name       string `json:"name"`
	Year       int    `json:"year"`
	CategoryId int    `json:"category_id"`
	Category   string `json:"category"`
}

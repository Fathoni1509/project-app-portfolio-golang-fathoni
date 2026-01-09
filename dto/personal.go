package dto

// type PersonalCreateRequest struct {
// 	Name        string `json:"name" validate:"required,min=3"`
// 	Age         int    `json:"age" validate:"required,gt=0"`
// 	Description string `json:"description" validate:"required,min=3"`
// }

type PersonalUpdateRequest struct {
	Name        *string `json:"name"`
	Age         *int    `json:"age"`
	Description *string `json:"description"`
}

type PersonalResponse struct {
	PersonalId  int    `json:"personal_id"`
	Name        string `json:"name"`
	Age         int    `json:"age"`
	Description string `json:"description"`
}

package dto

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

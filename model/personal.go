package model

type Personal struct {
	PersonalId  int    `json:"personal_id"`
	Name        string `json:"name"`
	Age         int    `json:"age"`
	Description string `json:"description"`
	Model
}

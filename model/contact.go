package model

type Contact struct {
	ContactId int    `json:"contact_id"`
	Name      string `json:"name"`
	Type      string `json:"type"`
	Link      string `json:"link"`
	Model
}

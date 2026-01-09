package model

type Activity struct {
	ActivityId int    `json:"activity_id"`
	Name       string `json:"name"`
	Year       int    `json:"year"`
	CategoryId int    `json:"category_id"`
	Model
}

package model

type Work struct {
	WorkId      int    `json:"work_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Year        int    `json:"year"`
	Model
}

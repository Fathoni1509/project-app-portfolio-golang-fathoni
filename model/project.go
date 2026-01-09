package model

type Project struct {
	ProjectId   int    `json:"project_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Year        int    `json:"year"`
	Link        string `json:"link"`
	ImageData   byte   `json:"image_data"`
	Model
}

package models

type Message struct {
	BaseModel
	Role    string `json:"role"`
	Content string `json:"content"`
}

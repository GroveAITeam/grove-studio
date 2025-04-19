package models

type Message struct {
	BaseModel
	ConversationID uint   `json:"conversation_id"`
	Role           string `json:"role"`
	Content        string `json:"content"`
}

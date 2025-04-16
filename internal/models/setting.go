package models

// Setting 应用程序设置模型
type Setting struct {
	BaseModel
	Key   string `gorm:"uniqueIndex" json:"key"`
	Value string `json:"value"`
}

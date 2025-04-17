package models

// Setting 应用程序设置
type Setting struct {
	BaseModel
	Key   string `gorm:"uniqueIndex" json:"key"`
	Value string `json:"value"`
}

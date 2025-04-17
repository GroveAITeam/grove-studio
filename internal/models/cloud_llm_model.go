package models

// CloudLLMModel 云端模型设置
type CloudLLMModel struct {
	BaseModel
	Name     string `json:"name"`
	Provider string `json:"provider"`
	EndPoint string `json:"endpoint"`
	ApiKey   string `json:"api_key"`
	Enabled  bool   `json:"enabled"`
}

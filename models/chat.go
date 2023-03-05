package models

type Mode string

const (
	ModeChat        Mode = "chat"
	ModeSetToken    Mode = "set_token"
	ModeSetLanguage Mode = "set_language"
)

type Chat struct {
	ChatId       int64    `json:"chat_id" bson:"chat_id"`
	Language     Language `json:"language" bson:"language"`
	Mode         Mode     `json:"mode" bson:"mode"`
	User         User     `json:"user" bson:"user"`
	ChatGPTToken *string  `json:"chat_gpt_token" bson:"chat_gpt_token"`
}

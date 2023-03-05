package models

// Account is a model of an account.
type Account struct {
	AccountId string `json:"accountId" bson:"accountId"`
	ChatId    string `json:"chatId" bson:"chatId"`
	Name      string `json:"name" bson:"name"`
}

package models

type Message struct {
	MessageId int64  `json:"message_id" bson:"message_id"`
	ChatId    int64  `json:"chat_id" bson:"chat_id"`
	From      int64  `json:"from" bson:"from"`
	Text      string `json:"text" bson:"text"`
	Timestamp int    `json:"timestamp" bson:"timestamp"`
	Archived  bool   `json:"archived" bson:"archived"`
}

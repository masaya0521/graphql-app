package model

type Todo struct {
	ID   string `json:"id"`
	Text string `json:"text"`
	User *User  `json:"user"`
}

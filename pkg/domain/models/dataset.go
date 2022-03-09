package models

type Dataset struct {
	ID        string `json:"id"`
	Timestamp int    `json:"timestamp"`
	Todos     []Todo `json:"todos"`
}

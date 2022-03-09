package models

type RawTodo struct {
	UserID    int    `json:"userId"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type Todo struct {
	ID          string `json:"id"`
	CreatedAt   string `json:"createdAt"`
	Description string `json:"description"`
	RawTodo
}

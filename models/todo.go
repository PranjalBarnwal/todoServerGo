package models

type Todo struct {
	ID        int    `json:"id" db:"id"`
	Title     string `json:"title" db:"title"`
	Completed bool   `json:"completed" db:"completed"`
}

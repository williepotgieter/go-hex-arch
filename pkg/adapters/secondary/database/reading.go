package database

import "github.com/williepotgieter/go-hex-arch/pkg/domain/models"

func (a *dbadapter) GetAllTodos() []models.Todo {
	data := a.db[len(a.db)-1]

	return data.Todos
}

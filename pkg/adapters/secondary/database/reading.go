package database

import "github.com/williepotgieter/go-hex-arch/pkg/domain/models"

func (a *dbadapter) GetAllTodos() []models.Todo {

	if len(a.db) > 0 {
		return a.db[len(a.db)-1].Todos
	}

	return []models.Todo{}
}

func (a *dbadapter) GetAllDatasets() []models.Dataset {
	return a.db
}

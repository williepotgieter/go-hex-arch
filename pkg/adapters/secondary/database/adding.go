package database

import "github.com/williepotgieter/go-hex-arch/pkg/domain/models"

func (a *dbadapter) AddTodos(dataset models.Dataset) error {
	a.db = append(a.db, dataset)

	return nil
}

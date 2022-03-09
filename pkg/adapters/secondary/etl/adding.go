package etl

import "github.com/williepotgieter/go-hex-arch/pkg/domain/models"

func (a *etladapter) LoadTodos(todos []models.Todo) error {

	return a.db.Add.Todos(todos)
}

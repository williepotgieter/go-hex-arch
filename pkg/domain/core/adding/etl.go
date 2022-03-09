package adding

import "github.com/williepotgieter/go-hex-arch/pkg/domain/models"

func (s *etlservice) LoadTodos(todos []models.Todo) error {
	return s.etl.LoadTodos(todos)
}

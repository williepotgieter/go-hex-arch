package reading

import "github.com/williepotgieter/go-hex-arch/pkg/domain/models"

func (s *dbservice) Todos() []models.Todo {
	return s.db.GetAllTodos()
}

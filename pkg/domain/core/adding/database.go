package adding

import (
	"time"

	"github.com/google/uuid"
	"github.com/williepotgieter/go-hex-arch/pkg/domain/models"
)

func (s *dbservice) Todos(todos []models.Todo) error {
	return s.db.AddTodos(models.Dataset{
		ID:        uuid.New().String(),
		Timestamp: int(time.Now().Unix()),
		Todos:     todos,
	})
}

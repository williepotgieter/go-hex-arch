package updating

import (
	"strings"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"github.com/williepotgieter/go-hex-arch/pkg/domain/models"
)

func (s *etlservice) TransformTodos(rawTodos []models.RawTodo) (todos []models.Todo) {
	for _, todo := range rawTodos {
		todos = append(todos, models.Todo{
			ID:          uuid.New().String(),
			CreatedAt:   time.Now().Local().Format(time.RFC3339),
			Description: gofakeit.Sentence(25),
			RawTodo: models.RawTodo{
				UserID:    todo.UserID,
				Title:     strings.Title(todo.Title),
				Completed: todo.Completed,
			},
		})
	}

	return
}

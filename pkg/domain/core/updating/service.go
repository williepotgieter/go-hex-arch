package updating

import "github.com/williepotgieter/go-hex-arch/pkg/domain/models"

type ETL interface{}

type ETLService interface {
	TransformTodos(rawTodos []models.RawTodo) (todos []models.Todo)
}

type etlservice struct {
	etl ETL
}

func NewETLService(etl ETL) ETLService {
	return &etlservice{etl}
}

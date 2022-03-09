package adding

import "github.com/williepotgieter/go-hex-arch/pkg/domain/models"

// ETL Service setup
type ETL interface {
	LoadTodos(todos []models.Todo) error
}

type ETLService interface {
	LoadTodos(todos []models.Todo) error
}

type etlservice struct {
	etl ETL
}

func NewETLService(etl ETL) ETLService {
	return &etlservice{etl}
}

// Database service setup
type Database interface {
	AddTodos(dataset models.Dataset) error
}

type DatabaseService interface {
	Todos(todos []models.Todo) error
}

type dbservice struct {
	db Database
}

func NewDBService(db Database) DatabaseService {
	return &dbservice{db}
}

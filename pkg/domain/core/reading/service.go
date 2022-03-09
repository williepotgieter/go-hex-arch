package reading

import "github.com/williepotgieter/go-hex-arch/pkg/domain/models"

// ETL Service setup
type ETL interface {
	ExtractTodos() (rawTodos []byte, err error)
}

type ETLService interface {
	ExtractTodos() (todos []models.RawTodo, err error)
}

type etlservice struct {
	etl ETL
}

func NewETLService(etl ETL) ETLService {
	return &etlservice{etl}
}

// Database service setup
type Database interface {
	GetAllTodos() []models.Todo
	GetAllDatasets() []models.Dataset
}

type DatabaseService interface {
	Todos() []models.Todo
	Datasets() []models.Dataset
}

type dbservice struct {
	db Database
}

func NewDBService(db Database) DatabaseService {
	return &dbservice{db}
}

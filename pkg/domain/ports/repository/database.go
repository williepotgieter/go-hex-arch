package repository

import (
	"github.com/williepotgieter/go-hex-arch/pkg/domain/core/adding"
	"github.com/williepotgieter/go-hex-arch/pkg/domain/core/reading"
)

type Database interface {
	adding.Database
	reading.Database
}

type DBPort struct {
	Add  adding.DatabaseService
	Read reading.DatabaseService
}

func NewDBPort(db Database) *DBPort {
	return &DBPort{
		Add:  adding.NewDBService(db),
		Read: reading.NewDBService(db),
	}
}

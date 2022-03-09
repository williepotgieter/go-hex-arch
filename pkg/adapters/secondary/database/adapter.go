package database

import "github.com/williepotgieter/go-hex-arch/pkg/domain/models"

type MemDB []models.Dataset

type dbadapter struct {
	db MemDB
}

func NewDBAdapter() *dbadapter {
	return &dbadapter{MemDB{}}
}

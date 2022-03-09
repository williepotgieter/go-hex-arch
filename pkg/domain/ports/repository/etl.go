package repository

import (
	"github.com/williepotgieter/go-hex-arch/pkg/domain/core/adding"
	"github.com/williepotgieter/go-hex-arch/pkg/domain/core/reading"
	"github.com/williepotgieter/go-hex-arch/pkg/domain/core/updating"
)

type ETL interface {
	adding.ETL
	reading.ETL
	updating.ETL
}

type ETLPort struct {
	Add    adding.ETLService
	Read   reading.ETLService
	Update updating.ETLService
}

func NewETLPort(etl ETL) *ETLPort {
	return &ETLPort{
		Add:    adding.NewETLService(etl),
		Read:   reading.NewETLService(etl),
		Update: updating.NewETLService(etl),
	}
}

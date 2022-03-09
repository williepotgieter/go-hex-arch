package etl

import (
	"net/http"

	"github.com/williepotgieter/go-hex-arch/pkg/domain/ports/repository"
)

type etladapter struct {
	client *http.Client
	db     *repository.DBPort
}

func NewETLAdapter(db *repository.DBPort) *etladapter {
	return &etladapter{http.DefaultClient, db}
}

package reading

import (
	"encoding/json"

	"github.com/williepotgieter/go-hex-arch/pkg/domain/models"
)

func (s *etlservice) ExtractTodos() (rawTodos []models.RawTodo, err error) {
	var rawData []byte

	rawData, err = s.etl.ExtractTodos()
	if err != nil {
		return
	}

	err = json.Unmarshal(rawData, &rawTodos)

	return
}

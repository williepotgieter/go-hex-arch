package scheduler

import (
	"log"

	"github.com/williepotgieter/go-hex-arch/pkg/domain/models"
)

func (a *cronadapter) RunCronJob() {
	var (
		rawTodos []models.RawTodo
		todos    []models.Todo
		err      error
	)

	// EXTRACT todo data from external API
	log.Println(">>>>> SCHEDULED DATA LOAD <<<<<")
	rawTodos, err = a.etl.Read.ExtractTodos()
	if err != nil {
		log.Printf("failed to perform scheduled data load: %s\n", err.Error())
		return
	}

	// TRANSFORM todo data
	log.Println("----- TRANSFORM DATA -----")
	todos = a.etl.Update.TransformTodos(rawTodos)

	// LOAD todo data into database
	err = a.etl.Add.LoadTodos(todos)
	if err != nil {
		log.Printf("failed to load data into database: %s\n", err.Error())
		a.err <- err
		return
	}
	log.Println("<<<<< DATASET SAVED TO DATABASE >>>>>")
}

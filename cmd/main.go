package main

import (
	"fmt"

	"github.com/williepotgieter/go-hex-arch/pkg/adapters/secondary/database"
	"github.com/williepotgieter/go-hex-arch/pkg/adapters/secondary/etl"
	"github.com/williepotgieter/go-hex-arch/pkg/domain/ports/repository"
)

func main() {
	dbAdapter := database.NewDBAdapter()
	dbPort := repository.NewDBPort(dbAdapter)

	etlAdapter := etl.NewETLAdapter(dbPort)
	etlPort := repository.NewETLPort(etlAdapter)

	rawTodos, _ := etlPort.Read.ExtractTodos()
	todos := etlPort.Update.TransformTodos(rawTodos)

	etlPort.Add.LoadTodos(todos)

	fmt.Println(dbPort.Read.Todos())
}

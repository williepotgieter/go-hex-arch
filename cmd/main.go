package main

import (
	"sync"

	"github.com/williepotgieter/go-hex-arch/pkg/adapters/primary/rest"
	"github.com/williepotgieter/go-hex-arch/pkg/adapters/primary/scheduler"
	"github.com/williepotgieter/go-hex-arch/pkg/adapters/secondary/database"
	"github.com/williepotgieter/go-hex-arch/pkg/adapters/secondary/etl"
	"github.com/williepotgieter/go-hex-arch/pkg/domain/ports/repository"
)

const (
	PORT          = 3000
	LOAD_INTERVAL = "* * * * *"
)

func main() {
	// Setup database adapter
	dbAdapter := database.NewDBAdapter()
	dbPort := repository.NewDBPort(dbAdapter)

	// Setup ETL adapter
	etlAdapter := etl.NewETLAdapter(dbPort)
	etlPort := repository.NewETLPort(etlAdapter)

	// Setup primary adapters
	cronService := scheduler.NewCRONAdapter(LOAD_INTERVAL, etlPort)
	apiService := rest.NewRESTAdapter(dbPort)

	// Setup waitgroup for running services (primary adapters) as goroutines
	wg := sync.WaitGroup{}
	wg.Add(2)

	go cronService.Run(&wg)
	go apiService.Run(PORT, &wg)

	wg.Wait()
}

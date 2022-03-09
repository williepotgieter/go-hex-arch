package scheduler

import (
	"log"
	"sync"

	"github.com/robfig/cron/v3"
	"github.com/williepotgieter/go-hex-arch/pkg/domain/ports/repository"
)

type cronadapter struct {
	c        *cron.Cron
	etl      *repository.ETLPort
	schedule string
	err      chan error
}

func NewCRONAdapter(schedule string, etl *repository.ETLPort) *cronadapter {
	return &cronadapter{
		c:        cron.New(),
		etl:      etl,
		schedule: schedule,
		err:      make(chan error),
	}
}

func (a *cronadapter) Run(wg *sync.WaitGroup) {

	defer wg.Done()

	// Setup CRON job
	a.c.AddFunc(a.schedule, func() {
		a.RunCronJob()
	})

	// Start CRON service
	a.c.Start()

	// Listen for database load errors and exit application
	// in the event of any errors with the CRON service
	log.Fatalf("failed to load data into database: %s\n", <-a.err)
}

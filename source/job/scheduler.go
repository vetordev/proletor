package job

import (
	c "github.com/robfig/cron/v3"
	"time"
)

type Scheduler struct {
	cron *c.Cron
}

func New(utc *time.Location) *Scheduler {
	cron := c.New(c.WithLocation(utc))

	return &Scheduler{cron}
}

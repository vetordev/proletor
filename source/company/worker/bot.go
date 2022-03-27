package worker

import "job"

type Bot struct {
	jobs []*job.Job
}

func (b Bot) Start() {
	for _, j := range b.jobs {
		(*j).Start()
	}
}

func New(jobs []*job.Job) Bot {
	return Bot{jobs}
}

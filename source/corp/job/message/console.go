package message

import (
	"fmt"
	"proletor/corp/job"
	"time"
)

type Console struct {
	Time    time.Time
	Content string
}

func (j Console) Start() {
	fmt.Printf("An %s, %s", j.Content, j.Time)
}

func NewConsole(time time.Time, content string) job.Job {
	return Console{time, content}
}

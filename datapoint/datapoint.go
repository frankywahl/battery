package datapoint

import (
	"time"
)

type Datapoint struct {
	Percentage int
	startTime  time.Time
}

func New() *Datapoint {
	return &Datapoint{150, time.Now()}
}

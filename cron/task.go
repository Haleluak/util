package cron

import (
	"fmt"
	"time"
)

type Task interface {
	Run(Command) error
	Status() string
}

type Command struct {
	Name string
	Func func() error
}

type Schedule struct {
	Time time.Time
	Interval time.Duration
}

func (c Command) Execute() error {
	return c.Func()
}

func (c Command) String() string {
	return c.Name
}

func (s Schedule) Run() <-chan time.Time {
	d := s.Time.Sub(time.Now())

	ch := make(chan time.Time, 1)

	go func() {
		// wait for start time
		<-time.After(d)

		// zero interval
		if s.Interval == time.Duration(0) {
			ch <- time.Now()
			close(ch)
			return
		}

		// start ticker
		ticker := time.NewTicker(s.Interval)
		defer ticker.Stop()
		for t := range ticker.C {
			ch <- t
		}
	}()

	return ch
}

func (s Schedule) String() string {
	return fmt.Sprintf("%d-%d", s.Time.Unix(), s.Interval)
}
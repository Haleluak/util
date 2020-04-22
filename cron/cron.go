package cron

import "fmt"

type Cron interface {
	Schedule(Schedule, Command) error
}

type syncCron struct {

}

func (c *syncCron) Schedule(s Schedule, t Command) error {
	id := fmt.Sprintf("%s-%s", s.String(), t.String())
	fmt.Println(id)
	go func() {
		tc := s.Run()
	Tick:
		for {
			select {
			case _, ok := <-tc:
				if !ok {
					break Tick
				}

				fmt.Print("[cron] executing command %s", t.Name)

				if err := t.Execute(); err != nil {
					fmt.Print("[cron] error executing command %s: %v", t.Name, err)
				}
			}
		}
	}()
	return nil
}

func NewCron() Cron {
	return &syncCron{
	}
}
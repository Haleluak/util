package main

import (
	"fmt"
	"time"
	"tutorial/cron"
)

func main()  {
	c := cron.NewCron()
	err := c.Schedule(cron.Schedule{Time: time.Now(), Interval: time.Second * 5}, cron.Command{
		Name: "Cron tab",
		Func: func() error {
			fmt.Println("call api")
			return nil
		},
	})

	if err != nil {
		fmt.Println(err.Error())
	}

	time.Sleep(time.Minute * 1)
}
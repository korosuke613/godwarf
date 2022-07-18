package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
)

func makeCron(configs *map[string]Config) (*cron.Cron, error) {
	c := cron.New()

	for name, config := range *configs {
		var schedule string
		if config.Schedule != "" {
			schedule = config.Schedule
		} else {
			schedule = "*/5 * * * *"
		}

		_, err := c.AddFunc(schedule, func() {
			fmt.Printf("Pull %s", name)
			pull(&config)
		})
		if err != nil {
			return nil, err
		}
	}

	return c, nil
}

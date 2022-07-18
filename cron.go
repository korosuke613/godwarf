package main

import (
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
)

func makeCron(configs *Configs, logger *zap.SugaredLogger) (*cron.Cron, error) {
	l := logger.With("action", "cron")
	c := cron.New()

	for name, config := range *configs {
		var schedule string
		if config.Schedule != "" {
			schedule = config.Schedule
		} else {
			schedule = "*/5 * * * *"
		}

		l.Infow("Add func",
			"name", name,
			"schedule", schedule,
		)

		cronFunc := func() {
			l.Infow("Start func",
				"name", name,
				"schedule", schedule,
			)

			gc, err := makeGitClient(&config, logger)
			if err != nil {
				l.Errorw("Failed init git client",
					"name", name,
					"schedule", schedule,
					"reason", err.Error(),
				)
			}

			err = gc.pull(&config)
			if err != nil {
				l.Errorw("Failed pull",
					"name", name,
					"schedule", schedule,
					"reason", err.Error(),
				)
			}

			sc, err := makeScriptClient(&config, logger)
			if err != nil {
				l.Errorw("Failed init script client",
					"name", name,
					"schedule", schedule,
					"reason", err.Error(),
				)
			}

			sc.beforeExec()
			sc.afterExec()

			l.Infow("Finish func",
				"name", name,
				"schedule", schedule,
			)
		}

		cronFunc()
		_, err := c.AddFunc(schedule, cronFunc)
		if err != nil {
			return nil, err
		}
	}

	return c, nil
}

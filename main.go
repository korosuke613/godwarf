package main

import (
	"github.com/robfig/cron/v3"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
	"os"
)

func makeLogger() *zap.SugaredLogger {
	loggerConfig := zap.NewProductionConfig()
	loggerConfig.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	loggerConfig.OutputPaths = []string{"./log.txt", "stdout"}

	logger, err := loggerConfig.Build()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	sugarLogger := logger.Sugar()

	return sugarLogger
}

func main() {
	app := &cli.App{
		Name: "greet",
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "name", Value: "bob", Usage: "a name to say"},
		},
		Action: func(con *cli.Context) error {
			logger := makeLogger()
			filePath := con.Args().Get(0)
			configs, err := readConfig(filePath, logger)
			if err != nil {
				panic(err)
			}

			var c *cron.Cron
			c, err = makeCron(configs, logger)
			c.Run()
			return nil
		},
		UsageText: "app [first_arg] [second_arg]",
		Authors:   []*cli.Author{{Name: "Oliver Allen", Email: "oliver@toyshop.example.com"}},
	}

	app.Run(os.Args)
}

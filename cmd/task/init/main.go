package initTaskApplication

import (
	"orchid-starter/cmd/task/init/handler/task"
	modelInit "orchid-starter/cmd/task/init/model"
	"orchid-starter/internal/bootstrap"

	"github.com/urfave/cli"
)

func NewInitTask(di *bootstrap.DirectInjection) cli.Command {
	return cli.Command{
		Name:    "init-task",
		Aliases: []string{"init-task"},
		Usage:   "Run init-task",
		Flags: []cli.Flag{
			cli.UintFlag{
				Name:  "id",
				Value: 1,
			},
			cli.BoolFlag{
				Name: "count",
			},
			cli.StringFlag{
				Name:     "name",
				Required: false,
			},
		},
		Action: func(c *cli.Context) error {
			return task.NewTask(di, modelInit.Init{
				ID: c.Uint64("id"),
			}).Start()
		},
	}
}

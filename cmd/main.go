package main

import (
	"log"
	"os"

	"github.com/arijitAD/alien-invasion/invasion"
	"github.com/urfave/cli"
)

// alien Invasion flags
var (
	alienCountFlag = cli.IntFlag{
		Name:     "alien-count",
		Usage:    "Number of aliens unleashed upon the World",
		Value:    0,
		Required: true,
	}
	worldMapFlag = cli.StringFlag{
		Name:        "world-map",
		Usage:       "File path of World map",
		Value:       "",
		Required:    true,
	}
)

func main() {
	app := cli.NewApp()
	app.Name = "alien Invaders"
	app.Copyright = "Copyright 2021 None"
	app.Usage = "Official invasion"
	app.Commands = cli.Commands{
		{
			Name:      "game",
			Action:    invadeWorld,
			Flags: []cli.Flag{
				alienCountFlag,
				worldMapFlag,
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatalln(err)
	}
}

func invadeWorld(ctx *cli.Context) error {
	fp := ctx.String(worldMapFlag.Name)

	alienCount := ctx.Int(alienCountFlag.Name)
	log.Println("Aliens to be spawned", alienCount)

	invasion := invasion.NewGame()

	log.Println("Initializing game")
	err := invasion.Init(fp, alienCount)
	if err != nil {
		return err
	}

	log.Println("Starting game")
	invasion.Start()

	// print the final state of the cities.
	invasion.PrintState()

	return nil
}

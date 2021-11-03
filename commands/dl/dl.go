package dl

import (
	"log"

	"github.com/ArjixWasTaken/anime-dl-go/utils"
	"github.com/urfave/cli/v2"
)

func PrintArgs(args cli.Args) {
	arguments := utils.ArgsToStringList(args)

	app := &cli.App{
		Name:        "dl",
		Description: "Search and download an anime.",
		Usage:       "anime dl [command options] [arguments...]",
	}

	err := app.Run(arguments)
	if err != nil {
		log.Fatal(err)
	}
}

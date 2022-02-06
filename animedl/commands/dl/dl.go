package dl

import (
	"errors"
	"fmt"
	"log"

	"github.com/ArjixWasTaken/animedl/animedl/commands"
	"github.com/ArjixWasTaken/animedl/animedl/utils"
	"github.com/urfave/cli/v2"
)

type DL struct {
	commands.Command
}

func RunWithArgs(args cli.Args) error {
	arguments := utils.ArgsToStringList(args)
	if len(arguments) == 0 {
		return errors.New("error: no arguments for `dl` were given")
	}

	query, found, newArgs := utils.ParseQueryFromArgs(arguments)
	if !found {
		return errors.New("error: no query was given for `dl`")
	}

	app := &cli.App{
		Name:            "anime dl",
		Usage:           "Search and download an anime.",
		UsageText:       "anime dl \"overlord\"",
		HideHelp:        true,
		HideHelpCommand: true,
		Action: func(c *cli.Context) error {

			fmt.Println(c.Args()) // should be empty for now
			fmt.Println(query)    // should be "overlord iii" if ran by the Makefile

			return nil
		},
	}

	err := app.Run(newArgs)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

package dl

import (
	"errors"
	"fmt"
	"log"

	"github.com/ArjixWasTaken/animedl/animedl/commands"
	"github.com/ArjixWasTaken/animedl/animedl/providers/allProviders"
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
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "provider",
				Value: "gogoanime",
				Usage: "the provider to use",
			},
		},
		Action: func(c *cli.Context) error {

			provider := allProviders.GetProviderByName(c.String("provider"))

			if provider == nil {
				log.Fatal("That provider does not exist.")
			} else {

				results := provider.Search(query)
				fmt.Println(results)
			}

			return nil
		},
	}

	err := app.Run(newArgs)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

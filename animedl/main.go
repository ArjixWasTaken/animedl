package main

import (
	"log"
	"os"

	"github.com/ArjixWasTaken/animedl/animedl/commands/dl"
	"github.com/urfave/cli/v2"
)

// TODO(Arjix): Create/Parse the config file.
// TODO(Arjix): Load the config file and provide it to the sub-commands.

func main() {
	app := &cli.App{
		Name:  "anime-dl",
		Usage: "A simple but powerful anime downloader and streamer.",
		Commands: []*cli.Command{
			{
				Name:  "dl",
				Usage: "Search and download an anime.",
				Action: func(c *cli.Context) error {
					error := dl.RunWithArgs(c.Args())
					if error != nil {
						log.Fatal(error)
					}
					return nil
				},
				SkipFlagParsing: false,
				HideHelp:        true,
				Hidden:          false,
			}, /*, // TODO(Arjix): Implement more sub-commands.
			{
				Name:  "watch",
				Usage: "Watch anime from your watch list.",
				Action: func(c *cli.Context) error {
					return nil
				},
				SkipFlagParsing: false,
				HideHelp:        false,
				Hidden:          false,
			},
			{
				Name:  "config",
				Usage: "Configure anime-dl with a fast and easy to use interface.",
				Action: func(c *cli.Context) error {
					return nil
				},
				SkipFlagParsing: false,
				HideHelp:        false,
				Hidden:          false,
			},*/
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

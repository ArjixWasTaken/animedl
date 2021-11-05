package main

import (
	"fmt"

	"github.com/ArjixWasTaken/anime-dl-go/providers"
	"github.com/ArjixWasTaken/anime-dl-go/providers/gogoanime"
)

var allProviderAPIs = []providers.ProviderInterface{gogoanime.GogoanimeProvider}

func main() {
	gogo := allProviderAPIs[0]
	fmt.Println(gogo.Search("overlord"))
	// app := &cli.App{
	// 	Name:  "anime-dl",
	// 	Usage: "A simple but powerful anime downloader and streamer.",
	// 	Commands: []*cli.Command{
	// 		{
	// 			Name:  "dl",
	// 			Usage: "Search and download an anime.",
	// 			Action: func(c *cli.Context) error {
	// 				dl.RunWithArgs(c.Args())
	// 				return nil
	// 			},
	// 			SkipFlagParsing: false,
	// 			HideHelp:        false,
	// 			Hidden:          false,
	// 		}, /*,
	// 		{
	// 			Name:  "watch",
	// 			Usage: "Watch anime from your watch list.",
	// 			Action: func(c *cli.Context) error {
	// 				return nil
	// 			},
	// 			SkipFlagParsing: false,
	// 			HideHelp:        false,
	// 			Hidden:          false,
	// 		},
	// 		{
	// 			Name:  "config",
	// 			Usage: "Configure anime-dl with a fast and easy to use interface.",
	// 			Action: func(c *cli.Context) error {
	// 				return nil
	// 			},
	// 			SkipFlagParsing: false,
	// 			HideHelp:        false,
	// 			Hidden:          false,
	// 		},*/
	// 	},
	// }

	// err := app.Run(os.Args)
	// if err != nil {
	// 	log.Fatal(err)
	// }
}

package dl

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func RunWithArgs(args cli.Args) {
	// arguments := utils.ArgsToStringList(args)

	// app := &cli.App{
	// 	Name:  "dl",
	// 	Usage: "Search and download an anime.",
	// }

	// err := app.Run(arguments)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	search("overlord")
}

func search(query string) {
	var choice string
	fmt.Scanln(&choice)
	fmt.Println(choice)
}

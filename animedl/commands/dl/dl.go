package dl

import (
	"fmt"

	"github.com/ArjixWasTaken/animedl/animedl/commands"
	"github.com/ArjixWasTaken/animedl/animedl/providers/allProviders"
	"github.com/urfave/cli/v2"
)

type DL struct {
	commands.Command
}

func RunWithArgs(args cli.Args) {
	// arguments := utils.ArgsToStringList(args)
	// if len(arguments) == 0 {
	// 	return
	// }

	// app := &cli.App{
	// 	Name:  "dl",
	// 	Usage: "Search and download an anime.",
	// }

	// err := app.Run(arguments)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	providers := allProviders.GetProviders()

	fmt.Println(providers)
}

package utils

import "github.com/urfave/cli/v2"

func ArgsToStringList(args cli.Args) []string {
	items := make([]string, args.Len())

	for index, _ := range items {
		items[index] = args.Get(index)
	}

	return items
}

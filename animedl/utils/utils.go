package utils

import "github.com/urfave/cli/v2"

func ArgsToStringList(args cli.Args) []string {
	items := make([]string, args.Len())

	for index, _ := range items {
		items[index] = args.Get(index)
	}

	return items
}

func ParseQueryFromArgs(args []string) (string, bool, []string) {

	var query string = ""
	var found bool = false
	newArgs := make([]string, 0, len(args))
	newArgs = append(newArgs, "")

	var withinParam bool = false

	for index, arg := range args {
		if arg[0] == '-' {
			newArgs = append(newArgs, arg)
			withinParam = true
		} else if !withinParam {
			// not really the best way to deal with this, but the cli library doesnt support it either.
			query = arg
			found = true
			break
		} else {
			newArgs = append(newArgs, arg)
		}
		if withinParam && index > 0 && args[index-1][0] == '-' {
			withinParam = false
		}
	}

	return query, found, newArgs
}

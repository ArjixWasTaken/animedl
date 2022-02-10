package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ArjixWasTaken/animedl/animedl/providers"
	"github.com/jedib0t/go-pretty/v6/list"
	"github.com/ttacon/chalk"
	"github.com/urfave/cli/v2"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

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

type Row struct {
	SlNo  int
	Title string
	Year  int64
}

func TabulateTheSearchResults(results []providers.SearchResult) string {

	var longestStr string = ""

	for _, item := range results {
		if len(item.Title) > len(longestStr) {
			longestStr = item.Title
		}
	}

	padding := len(longestStr)

	l := list.NewWriter()
	l.SetStyle(list.StyleBulletTriangle)

	for index, row := range results {
		indexStr := fmt.Sprint(index + 1)

		if len(indexStr) == 1 {
			indexStr = " " + indexStr
		}

		itemStr := chalk.Magenta.Color(fmt.Sprintf("%s. %s", indexStr, row.Title))

		if row.Year != 0 {
			itemStr += strings.Repeat(" ", padding-len(row.Title))
			itemStr += fmt.Sprintf(" [%s]", chalk.Green.Color(fmt.Sprint(row.Year)))
		}
		l.AppendItem(itemStr)
	}
	return l.Render()
}

func GetUserInput(message string) string {
	fmt.Print(message)
	input, _ := reader.ReadString('\n')

	return input
}

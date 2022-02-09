package dl

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/ArjixWasTaken/animedl/animedl/providers/allProviders"
	"github.com/ArjixWasTaken/animedl/animedl/utils"
	"github.com/briandowns/spinner"
	"github.com/ttacon/chalk"
	"github.com/urfave/cli/v2"
)

func getIndexFromUser(start int, end int) int64 {
	var input string = strings.Trim(utils.GetUserInput("Enter the anime no: [1]: "), " \n\r")

	if input == "" {
		return 0
	}

	inputAsNum, err := strconv.ParseInt(input, 10, 0)

	if err != nil {
		fmt.Println(chalk.Underline.TextStyle("Wrong input! Enter a number from " + strconv.Itoa(start) + " to " + strconv.Itoa(end)))
		fmt.Println(chalk.Red, err.Error(), chalk.Reset)
		return getIndexFromUser(start, end)
	}

	if int(inputAsNum) < start || int(inputAsNum) > end {
		fmt.Println(chalk.Underline.TextStyle("Wrong input! Enter a number from " + strconv.Itoa(start) + " to " + strconv.Itoa(end)))
		return getIndexFromUser(start, end)
	}

	return inputAsNum - 1
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
		Name:                   "anime dl",
		Usage:                  "Search and download an anime.",
		UsageText:              "anime dl \"overlord\"",
		HideHelp:               true,
		HideHelpCommand:        true,
		UseShortOptionHandling: true,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "provider",
				Value: "gogoanime",
				Usage: "the provider to use",
			}, // TODO(Arjix): Add a `--choice` param to skip the search section.
		},
		Action: func(c *cli.Context) error {

			provider := allProviders.GetProviderByName(c.String("provider"))

			if provider == nil {
				log.Fatal("That provider does not exist.")
			} else {
				var index int

				s := spinner.New(spinner.CharSets[9], 200*time.Millisecond)
				s.Suffix = fmt.Sprintf(" Searching for %s", chalk.Italic.TextStyle(query))
				s.Start()

				results := provider.Search(query)

				s.Stop()
				fmt.Println(utils.TabulateTheSearchResults(results))

				index = int(getIndexFromUser(1, len(results)))

				fmt.Println(chalk.Green, fmt.Sprintf("Selected [%d]: %s - %d\n", index+1, results[index].Title, results[index].Year), chalk.Reset)
				fmt.Println(chalk.Italic.TextStyle(fmt.Sprintf("Fetching info for `%s` from `%s`...\n", chalk.Cyan.Color(results[index].Title), chalk.Yellow.Color(provider.Name))))
				anime := provider.Load(results[index].Url)
				fmt.Println(anime)
			}

			return nil
		},
	}

	err := app.Run(newArgs)
	if err != nil {
		log.Fatal(err.Error())
	}

	return nil
}

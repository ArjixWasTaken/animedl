package dl

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/ArjixWasTaken/animedl/animedl/providers/allProviders"
	"github.com/ArjixWasTaken/animedl/animedl/utils"
	"github.com/urfave/cli/v2"
)

func getIndexFromUser(start int, end int) int64 {
	var input string = strings.Trim(utils.GetUserInput("Enter the anime no: [1]: "), " \n\r")

	if input == "" {
		return 0
	}

	inputAsNum, err := strconv.ParseInt(input, 10, 64)

	if err != nil {
		fmt.Println("Wrong input! Enter a number from " + strconv.Itoa(start) + " to " + strconv.Itoa(end))
		fmt.Println(err.Error())
		return getIndexFromUser(start, end)
	}

	if int(inputAsNum) < start || int(inputAsNum) > end {
		fmt.Println("Wrong input! Enter a number from " + strconv.Itoa(start) + " to " + strconv.Itoa(end))
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
				fmt.Println(utils.TabulateTheSearchResults(results))
				index := int(getIndexFromUser(1, len(results)))
				fmt.Println(results[index])
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

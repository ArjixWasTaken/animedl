package gogoanime

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/ArjixWasTaken/animedl/animedl/providers"
	"github.com/ArjixWasTaken/animedl/animedl/utils"
	"github.com/PuerkitoBio/goquery"
)

const mainUrl, apiName = "https://gogoanime.film", "gogoanime"

var GogoanimeProvider = &providers.Provider{
	Name:    apiName,
	MainUrl: mainUrl,
	Search: func(query string) []providers.SearchResult {
		link := mainUrl + "/search.html?keyword=" + query
		response := utils.Get(link, map[string]string{})
		soup, err := utils.Soupify(*response)

		if err != nil {
			log.Fatal(err)
		} else {
			var results = make([]providers.SearchResult, soup.Find(".last_episodes li").Length())

			soup.Find(".last_episodes li").Each(func(i int, s *goquery.Selection) {
				var year int64 = 0

				yearText := s.Find(".released").First()

				if yearText != nil && strings.Contains(yearText.Text(), ":") {
					year, err = strconv.ParseInt(strings.Trim(strings.Split(yearText.Text(), ":")[1], " "), 10, 64)
				}

				results[i] = providers.SearchResult{
					Title:   s.Find(".name").First().Text(),
					Url:     mainUrl + s.Find(".name > a").First().AttrOr("href", "none"),
					ApiName: apiName,
					Poster:  s.Find("img").First().AttrOr("src", "none"),
					Year:    year,
				}
			})

			return results
		}
		return nil
	},
	Load: func(url string) providers.LoadResponse {
		response := utils.Get(url, map[string]string{})
		soup, err := utils.Soupify(*response)

		if err != nil {
			fmt.Println(soup)
		} else {
			log.Fatal(err)
		}

		return providers.LoadResponse{}
	},
	LoadLinks: func(url string) []providers.ExtractorLink {
		return nil
	},
}

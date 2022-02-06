package gogoanime

import (
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

				if s.Find(".released").First() != nil {
					year, err = strconv.ParseInt(strings.Trim(strings.Split(s.Find(".released").First().Text(), ":")[1], " "), 10, 64)
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
		return providers.LoadResponse{}
	},
	LoadLinks: func(url string) []providers.ExtractorLink {
		return nil
	},
}

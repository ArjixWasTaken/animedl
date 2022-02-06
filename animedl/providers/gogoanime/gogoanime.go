package gogoanime

import (
	"log"
	"strings"

	"github.com/ArjixWasTaken/animedl/animedl/providers"
	"github.com/ArjixWasTaken/animedl/animedl/utils"
	"github.com/PuerkitoBio/goquery"
)

var mainUrl string = "https://gogoanime.film"
var apiName string = "gogoanime"

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
				results[i] = providers.SearchResult{
					Title:   strings.Replace(s.Find(".name").First().Text(), " (Dub)", "", 1),
					Url:     s.Find(".name > a").First().AttrOr("href", "none"),
					ApiName: apiName,
					Poster:  s.Find("img").First().AttrOr("src", "none"),
					Year:    0,
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

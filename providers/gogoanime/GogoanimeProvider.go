package gogoanime

import (
	"log"
	"strings"

	"github.com/ArjixWasTaken/anime-dl-go/providers"
	"github.com/ArjixWasTaken/anime-dl-go/utils"
	"github.com/PuerkitoBio/goquery"
)

type Provider struct {
	MainUrl string
	ApiName string
}

func (p Provider) Search(query string) []providers.SearchResult {
	link := p.MainUrl + "/search.html?keyword=" + query
	response := utils.Get(link, map[string]string{})
	soup, err := utils.Soupify(*response)
	if err != nil {
		log.Fatal(err)
	}

	var results = make([]providers.SearchResult, soup.Find(".last_episodes li").Length())

	soup.Find(".last_episodes li").Each(func(i int, s *goquery.Selection) {
		results[i] = providers.SearchResult{
			Title:   strings.Replace(s.Find(".name").First().Text(), " (Dub)", "", 1),
			Url:     s.Find(".name > a").First().AttrOr("href", "none"),
			ApiName: "gogoanime",
			Poster:  s.Find("img").First().AttrOr("src", "none"),
			Year:    0,
		}
	})
	return results
}

func (p Provider) Load(url string) providers.LoadResponse {
	return providers.LoadResponse{}
}

func (p Provider) LoadLinks(url string) []providers.ExtractorLink {
	return []providers.ExtractorLink{}
}

var GogoanimeProvider = Provider{MainUrl: "https://gogoanime.vc", ApiName: "gogoanime"}

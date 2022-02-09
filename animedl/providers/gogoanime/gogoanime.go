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

func getProperAnimeLink(url string) string {
	if strings.Contains(url, "-episode") {
		split := strings.Split(url, "/")
		slug := strings.Split(split[len(split)-1], "-episode")[0]
		return mainUrl + "/category/" + slug
	}
	return url
}

func getTvStatus(status string) providers.TvStatus {
	switch status {
	case "Completed":
		return providers.Completed
	case "Ongoing":
		return providers.Ongoing
	}

	return providers.Completed
}

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
		link := getProperAnimeLink(url)

		response := utils.Get(link, map[string]string{})
		soup, err := utils.Soupify(*response)

		if err != nil {
			log.Fatal(err.Error())
		}

		animeBody := soup.Find(".anime_info_body_bg").First()

		var genres []string = make([]string, 0)
		var description string
		var year int
		var status providers.TvStatus

		animeBody.Find("p.type").Each(func(i int, s *goquery.Selection) {
			switch strings.Trim(s.Find("span").First().Text(), " \n\r\t") {
			case "Genre:":
				s.Find("a").Each(func(i int, s *goquery.Selection) {
					genre := s.AttrOr("title", "")
					if len(genre) > 0 {
						genres = append(genres, genre)
					}
				})
			case "Plot Summary:":
				description = strings.Replace(s.Text(), "Plot Summary:", "", 1)
			case "Released:":
				year, _ = strconv.Atoi(strings.Replace(s.Text(), "Released:", "", 1))
			case "Status:":
				status = getTvStatus(strings.Trim(strings.Replace(s.Text(), "Status:", "", 1), " \n\r\t"))
			}
		})

		var episodes []providers.Episode = make([]providers.Episode, 0)
		// TODO(Arjix): Actually scrape the gogoanime episodes.

		return providers.LoadResponse{
			ApiName:     apiName,
			Title:       animeBody.Find("h1").First().Text(),
			Poster:      animeBody.Find("img").First().AttrOr("src", ""),
			Tags:        genres,
			Description: description,
			Year:        int64(year),
			TvStatus:    status,
			Episodes:    episodes,
		}
	},
	LoadLinks: func(url string) []providers.ExtractorLink {
		// TODO(Arjix): Implement this.
		return nil
	},
}

package gogoanime

import "github.com/ArjixWasTaken/animedl/animedl/providers"

var GogoanimeProvider = &providers.Provider{
	Name:    "Gogoanime",
	MainUrl: "https://gogoanime.wiki",
	Search: func(query string) []providers.SearchResult {
		return nil
	},
	Load: func(url string) providers.LoadResponse {
		return providers.LoadResponse{}
	},
	LoadLinks: func(url string) []providers.ExtractorLink {
		return nil
	},
}

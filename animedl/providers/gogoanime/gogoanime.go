package gogoanime

import "github.com/ArjixWasTaken/animedl/animedl/providers"

func Search(query string) []providers.SearchResult {
	return nil
}

func Load(url string) providers.LoadResponse {
	return providers.LoadResponse{}
}

func LoadLinks(url string) []providers.ExtractorLink {
	return nil
}

var GogoanimeProvider = &providers.Provider{
	Name:      "Gogoanime",
	MainUrl:   "https://gogoanime.wiki",
	Search:    Search,
	Load:      Load,
	LoadLinks: LoadLinks,
}

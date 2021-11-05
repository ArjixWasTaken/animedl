package providers

type SearchResult struct {
	Title   string
	Url     string
	ApiName string
	Poster  string
	Year    int
}

type DubStatus int

const (
	Subbed DubStatus = iota
	Dubbed DubStatus = iota
)

type TvStatus int

const (
	Ongoing   TvStatus = iota
	Completed TvStatus = iota
	Cancelled TvStatus = iota
)

type Episode struct {
	Title       string
	Url         string
	EpisodeNum  float32
	EpisodeType DubStatus
	Description string
	Date        string
}

type LoadResponse struct {
	Title       string
	Url         string
	ApiName     string
	Episodes    []Episode
	TvType      TvStatus
	Poster      string
	Description string
	Year        int
	Tags        []string
}

type ExtractorLink struct {
	Title   string
	Url     string
	Headers map[string]string
	IsM3u8  bool
}

type ProviderInterface interface {
	// Performs a search on the provider with the given query and returns a list of the results.
	Search(query string) []SearchResult

	// Fetches all of the info for an anime on the provider using the given url.
	Load(url string) LoadResponse

	// Fetches all the video streams for the given url.
	LoadLinks(url string) []ExtractorLink
}

type Provider struct {
	MainUrl string
	ApiName string
}

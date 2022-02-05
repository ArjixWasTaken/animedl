package utils

import (
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func Get(url string, headers map[string]string) *http.Response {
	client := http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		//Handle Error
	}

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	res, err := client.Do(req)
	if err != nil {
		//Handle Error
	}
	return res
}

func Post(url string, headers map[string]string) *http.Response {
	client := http.Client{}
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		//Handle Error
	}

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	res, err := client.Do(req)
	if err != nil {
		//Handle Error
	}
	return res
}

// Warning, this closes the response's body.
func Soupify(response http.Response) (*goquery.Document, error) {
	return goquery.NewDocumentFromReader(response.Body)
}

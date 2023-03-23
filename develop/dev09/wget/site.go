package wget

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/YoungGoofy/wb_l2/develop/dev09/parsing"
)

func getHtml(url string) (io.ReadCloser, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("error: status code [%d]", res.StatusCode))
	}
	return res.Body, nil
}

func saveInFile(b io.ReadCloser, filename string) error {
	body, err := io.ReadAll(b)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filename, body, 0644)
	if err != nil {
		return err
	}
	return nil
}

func getLinks(body io.ReadCloser) ([]string, error) {
	res, err := parsing.ParseHTML(body)
	if err != nil {
		return nil, err
	}

	links := make([]string, 0)

	for _, v := range res {
		href := v.Href.Host + v.Href.Path
		links = append(links, "https://"+href)
	}

	return links, nil
}

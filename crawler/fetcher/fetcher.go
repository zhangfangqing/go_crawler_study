package fetcher

import (
	"net/http"
	"golang.org/x/text/transform"
	"golang.org/x/text/encoding/simplifiedchinese"
	"io/ioutil"
	"errors"
)

func FetchWebContent(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("http get error.code")
	}

	uft8Reader := transform.NewReader(resp.Body, simplifiedchinese.GBK.NewDecoder())
	return ioutil.ReadAll(uft8Reader)
}

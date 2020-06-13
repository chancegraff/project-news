package vendors

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func getArticles(baseURL string, token string) (*[]byte, error) {
	url := fmt.Sprint(baseURL, token)

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return &body, err
}

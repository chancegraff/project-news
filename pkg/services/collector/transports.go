package collector

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func decodeGetRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request getRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeAllRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request allRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeArticlesResponse(_ context.Context, r *http.Response) (interface{}, error) {
	var response articlesResponse
	if err := json.NewDecoder(r.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func encodeRequest(_ context.Context, r *http.Request, request interface{}) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(request); err != nil {
		return err
	}
	r.Body = ioutil.NopCloser(&buf)
	return nil
}

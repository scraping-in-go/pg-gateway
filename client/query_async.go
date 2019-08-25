package client

import (
	"encoding/json"
	"github.com/bcicen/jstream"
	"github.com/just1689/pg-gateway/query"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

//GetEntityManyAsync asynchronously fetches 0..many rows by a query
func GetEntityManyAsync(baseURL string, query query.Query) (results chan []byte, err error) {
	url := query.ToURL(baseURL)
	resp, err := http.Get(url)
	if err != nil {
		logrus.Error(err)
		resp.Body.Close()
		return
	}
	results = make(chan []byte)
	go closerToChan(resp.Body, results)
	return
}

func closerToChan(body io.ReadCloser, results chan []byte) {
	decoder := jstream.NewDecoder(body, 1)
	for mv := range decoder.Stream() {
		b, _ := json.Marshal(mv.Value)
		results <- b
	}
	close(results)
}

package client

import (
	"github.com/just1689/json2channel/j2c"
	"github.com/just1689/pg-gateway/query"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

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
	defer body.Close()
	ri := j2c.NewReaderItemFromReadCloser(body)
	in := j2c.BuildInterpreter(ri)
	out := j2c.ReadObjects(in, ".")
	for o := range out {
		results <- []byte(o)
	}
	close(results)

}

package client

import (
	"github.com/just1689/pg-gateway/query"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

//GetEntityMany synchronously fetches 0..many rows by a query
func GetEntityMany(baseURL string, query query.Query) (result []byte, err error) {
	url := query.ToURL(baseURL)
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		logrus.Error(err)
		return
	}
	result, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		logrus.Error(err)
	}
	return

}

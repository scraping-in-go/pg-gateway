package client

import (
	"github.com/just1689/json2channel/j2c"
	"github.com/sirupsen/logrus"
	"net/http"
)

func GetEntityAllAsync(svr string, entity string) (results chan []byte, err error) {
	url := svr + "/" + entity
	resp, err := http.Get(url)
	if err != nil {
		logrus.Error(err)
		resp.Body.Close()
		return
	}

	results = make(chan []byte)
	go func() {
		defer resp.Body.Close()
		ri := j2c.NewReaderItemFromReadCloser(resp.Body)
		in := j2c.BuildInterpreter(ri)
		out := j2c.ReadObjects(in, ".")
		for o := range out {
			results <- []byte(o)
		}
		close(results)
	}()
	return

}

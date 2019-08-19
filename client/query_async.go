package client

import (
	"github.com/just1689/json2channel/j2c"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

func NewReaderItem(body io.ReadCloser) *readerItem {
	result := readerItem{
		Body:  body,
		Bytes: make([]byte, 1),
	}
	return &result
}

type readerItem struct {
	Body  io.ReadCloser
	Bytes []byte
}

func (r *readerItem) Next() (b byte, eof bool) {
	n, err := r.Body.Read(r.Bytes)
	if err != nil || n == 0 {
		return 0, true
	}
	return r.Bytes[0], false
}

func GetEntityAllAsync(svr string, entity string) (results chan []byte, err error) {
	url := svr + "/" + entity
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		logrus.Error(err)
		return
	}

	results = make(chan []byte)
	go func() {
		ri := NewReaderItem(resp.Body)
		in := j2c.BuildInterpreter(ri)
		out := j2c.ReadObjects(in, ".")
		for o := range out {
			results <- []byte(o)
		}
		close(results)
	}()
	return

}

//func GetEntityMany(svr string, entity, field, id string) (result []byte, err error) {
//	url := svr + "/" + entity + "/" + field + "/" + id
//	resp, err := http.Get(url)
//	defer resp.Body.Close()
//	if err != nil {
//		logrus.Error(err)
//		return
//	}
//	result, err = ioutil.ReadAll(resp.Body)
//	if err != nil {
//		logrus.Error(err)
//	}
//	return
//
//}
//
//func GetEntityByID(svr string, entity, id string) (result []byte, err error) {
//	url := svr + "/" + entity + "/" + id
//	resp, err := http.Get(url)
//	defer resp.Body.Close()
//	if err != nil {
//		logrus.Error(err)
//		return
//	}
//	result, err = ioutil.ReadAll(resp.Body)
//	if err != nil {
//		logrus.Error(err)
//	}
//	return
//
//}

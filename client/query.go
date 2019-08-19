package client

import (
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

func GetEntityAll(svr string, entity string) (result []byte, err error) {
	url := svr + "/" + entity
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

func GetEntityMany(svr string, entity, field, id string) (result []byte, err error) {
	url := svr + "/" + entity + "/" + field + "/" + id
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

func GetEntityByID(svr string, entity, id string) (result []byte, err error) {
	url := svr + "/" + entity + "/" + id
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

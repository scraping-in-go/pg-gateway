package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

//Insert accepts a server addr, the tables name and the instance to be inserted
func Insert(svr string, entity string, i interface{}) (err error) {
	url := svr + "/" + entity
	b, err := json.Marshal(i)
	if err != nil {
		logrus.Error(err)
		return
	}
	resp, err := http.Post(url, "Application/json", bytes.NewReader(b))
	if err != nil {
		logrus.Error(err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		err = errors.New(strconv.Itoa(resp.StatusCode))
	}
	return

}

//Update accepts a server addr, the tables name, the id value for row to be updated
func Update(svr string, entity, field, id string, i interface{}) (err error) {
	url := svr + "/" + entity + "/" + field + "/" + id
	b, err := json.Marshal(i)
	if err != nil {
		logrus.Error(err)
		return
	}
	req, err := http.NewRequest(http.MethodPatch, url, bytes.NewBuffer(b))
	if err != nil {
		logrus.Error(err)
		return
	}
	req.Header.Set("Content-type", "Application/json")

	client := http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		logrus.Error(err)
		return
	}

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		err = errors.New(strconv.Itoa(resp.StatusCode))
	}
	return
}

//Delete accepts a server addr, the tables name, the id value for row to be deleted
func Delete(svr string, entity, field, id string) (err error) {
	url := svr + "/" + entity + "/" + field + "/" + id
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		logrus.Error(err)
		return
	}
	req.Header.Set("Content-type", "Application/json")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logrus.Error(err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		err = errors.New(strconv.Itoa(resp.StatusCode))
	}
	return
}

package db

import (
	"github.com/scraping-in-go/svc-db-gateway/model"
	"github.com/sirupsen/logrus"
)

func GetEntityByID(id string) (entity model.Entity, err error) {
	conn, err := Connect()
	defer conn.Close()
	if err != nil {
		logrus.Errorln(err)
		return
	}

	entity = model.Entity{}
	sql := "select * from entities where id=$1"
	err = conn.QueryRow(sql, id).Scan(&entity.Entity, &entity.ID, &entity.V)
	if err != nil {
		logrus.Errorln(err)
		return
	}
	return

}

func Insert(entity model.Entity) (err error) {
	conn, err := Connect()
	defer conn.Close()
	if err != nil {
		logrus.Errorln(err)
		return
	}

	sql := "insert into entities values($1, $2, $3)"
	_, err = conn.Exec(sql, entity.Entity, entity.ID, string(entity.V))
	if err != nil {
		logrus.Errorln(err)
		return
	}
	return

}

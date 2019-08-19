package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/just1689/pg-gateway/client"
)

var svr = "http://localhost:8080"

type user struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func main() {
	//testInsert()
	c, err := client.GetEntityAllAsync(svr, "users")
	if err != nil {
		panic(err)
	}
	for r := range c {
		fmt.Println(string(r))
	}
}

func testInsert() {
	for i := 0; i < 1000; i++ {
		u := user{
			ID:        uuid.New().String(),
			FirstName: "Justin",
			LastName:  "Tamblyn",
			Email:     uuid.New().String(),
			Password:  "some_hash",
		}
		err := client.Insert(svr, "users", u)
		if err != nil {
			panic(err)
		}

	}

}

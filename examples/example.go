package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/just1689/pg-gateway/client"
)

var svr = "http://localhost:8080"

func main() {
	testInsert()
	testReadAsync()
}

var userEntities = "users"

type user struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
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
		if err := client.Insert(svr, userEntities, u); err != nil {
			panic(err)
		}
	}
}

func testReadAsync() {
	c, err := client.GetEntityAllAsync(svr, "users")
	if err != nil {
		panic(err)
	}
	count := 0
	for r := range c {
		count++
		fmt.Println(count, string(r))
	}

}

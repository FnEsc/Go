package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type Person struct {
	NickName string
	Phone    string
}

func main() {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// 切换模式为单调模式
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("test").C("people")
	err = c.Insert(
		&Person{"superWang", "13800138000"},
		&Person{"David", "13800138001"},
	)
	if err != nil {
		log.Fatal(err)
	}

	result := Person{}
	err = c.Find(bson.M{"nickname": "superWang"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Name:", result.NickName)
	fmt.Println("Phone:", result.Phone)
}

package config

import (
	"dblogin"
	"fmt"
	"gopkg.in/mgo.v2"
)

// DB variables
var (
	DB    *mgo.Database
	Books *mgo.Collection
)

func init() {
	// S get a mongo sessions
	s, err := mgo.Dial(dblogin.Bookstore) // mongodb://username:yourpasscode@serverip:27017/database?authSource=admin
	if err != nil {
		panic(err)
	}

	if err = s.Ping(); err != nil {
		panic(err)
	}

	DB = s.DB("bookstore")
	Books = DB.C("books")

	fmt.Println("You connected to your mongo database.")
}

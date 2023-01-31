package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func main() {
	router := httprouter.New()

	router.GET("/users", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Println("GOLANG SERVER!!!!!!")
	})

	http.ListenAndServe("localhost:7070", router)
}

func getSession() *mgo.Session {
	session, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}

	return session
}
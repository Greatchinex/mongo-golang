package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"

	"github.com/Greatchinex/mongo-golang/controllers"
)

func main() {
	router := httprouter.New()

	uc := controllers.NewUserController(getSession())
	router.GET("/user/:id", uc.GetUser)
	router.GET("/users", uc.GetAllUsers)
	router.POST("/user", uc.CreateUser)
	router.DELETE("/user/:id", uc.DeleteUser)

	http.ListenAndServe("localhost:7070", router)
}

func getSession() *mgo.Session {
	session, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}

	return session
}
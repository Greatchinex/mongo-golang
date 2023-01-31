package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()

	router.GET("/users", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Println("GOLANG SERVER!!!!!!")
	})
}
package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/Greatchinex/mongo-golang/models"
)


type UserController struct {
	session *mgo.Session
}

func NewUserController(session *mgo.Session) *UserController {
	return &UserController{session}
}

func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id") // Request params
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	userId := bson.ObjectIdHex(id)
	user := models.User{}

	err := uc.session.DB("mongo-golang").C("users").FindId(userId).One(&user)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	userToJson, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", userToJson)
}

func (uc UserController) GetAllUsers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	users := make([]models.User, 0)

	err := uc.session.DB("mongo-golang").C("users").Find(bson.M{}).All(&users)
	if err != nil {
		fmt.Printf("Something went wrong ===> %v\n", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	userToJson, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", userToJson)
}
 
func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	user := models.User{}
	json.NewDecoder(r.Body).Decode(&user)

	user.Id = bson.NewObjectId()

	// Insert user into DB
	uc.session.DB("mongo-golang").C("users").Insert(user)

	userToJson, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", userToJson)
}

func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
		return
	}
 
	userId := bson.ObjectIdHex(id)

	if err := uc.session.DB("mongo-golang").C("users").RemoveId(userId); err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Deleted user ===> ", userId, "\n")
} 
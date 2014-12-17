package hello

import (
	//"time"
	//"strings"
	//"strconv"
	"log"
	"appengine"
	"appengine/datastore"
	"net/http"
	"encoding/json"
	//"html/template"
	//"github.com/gorilla/mux"
	//"gopkg.in/mgo.v2/bson"
	)

type User struct {
	FirstName	string	`json:"firstname"`
	LastName	string	`json:"lastname"`
	Age			string	`json:"age"`
}

type Users struct {
	UserArray []User
}

type UsersArray []User

func userGetHandler(w http.ResponseWriter, r *http.Request) {
	//Test1 := User{
	//	FirstName:"George",
	//	LastName:"T",
	//	Age:"12"}

	//Test2 := User{
	//	FirstName:"George",
	//	LastName:"T",
	//	Age:"13"}

	//ArrayTest := []User{Test1,Test2}
	
	//log.Println(ArrayTest)
	
	AllUsers, err := usersFindAll(w, r)

	log.Println(AllUsers)

	if err == nil {
		encoder := json.NewEncoder(w)
		
		err1 := encoder.Encode(&AllUsers)
		if err1 == nil {
			log.Println("Returned users list")
		} else {
			log.Println("Unable to encode users JSON", err1)
		}
	} else {
		log.Println("Error retrieving users", err)
	}		
	
	
}

func userPostHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	
	var newUser User
	err1 := decoder.Decode(&newUser)
	if err1 == nil {
		newUser.userSave(w, r)
	} else {
		log.Println("Unable to decode the JSON request", err1);
	}
}

func (p *User) userSave(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	
	key := datastore.NewIncompleteKey(c, "test", nil)

	if _, err := datastore.Put(c, key, p); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	
	datastore.Get(c, key, p)
	
	log.Println("Successfully save", p.FirstName)
}

func usersFindAll(w http.ResponseWriter, r *http.Request) (UsersArray, error) {
	c := appengine.NewContext(r)
		
	q := datastore.NewQuery("test")
		
	var UsersResult []User
	
	_, err := q.GetAll(c, &UsersResult)
		if err != nil {
		log.Println("Error retrieving datastore ", err)
		return nil, err
	}

	return UsersResult, nil
}
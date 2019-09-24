package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm"
	"net/http"
	"time"
	_ "time"
)

var db = connect()

type User struct {
	Name            string `json:"name"`
	First_last_name string `json:"first_last_name"`
	Email           string `json:"email"`
	CreatedAt       time.Time   `json:"created_at"`
	UpdatedAt       time.Time   `json:"updated_at"`
	DeletedAt       *time.Time  `json:"deleted_at"`
}

func index(writer http.ResponseWriter, request *http.Request) {
	var users []User
	db.Find(&users)

	json.NewEncoder(writer).Encode(users)
}

func show(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]
	var user User
	db.Where("id = ?", id).Find(&user)
	json.NewEncoder(writer).Encode(user)
}

func create(writer http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	name := r.FormValue("name")
	firstLastName := r.FormValue("first_last_name")
	email := r.FormValue("email")

	db.Create(&User{
		Name:            name,
		First_last_name: firstLastName,
		Email:           email,
		CreatedAt:       time.Now().UTC(),
		UpdatedAt:       time.Now().UTC(),
		DeletedAt:       nil,
	})

	fmt.Fprintln(writer, "New user created successfully!")
}


func update(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]

	var user User
	db.Where("id = ?", id).Find(&user)

	request.ParseForm()

	name := request.FormValue("name")
	firstLastName := request.FormValue("first_last_name")
	email := request.FormValue("email")

	db.Model(&user).Where("id = ?", id).Update(&User{
		Name:            name,
		First_last_name: firstLastName,
		Email:           email,
		UpdatedAt:       time.Now().UTC(),
	})

	fmt.Fprintln(writer, "User updated successfully!")
}

func delete(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]
	var user User
	db.Where("id = ?", id).Delete(&user)

	fmt.Fprintln(writer, "User delete successfully!")
}

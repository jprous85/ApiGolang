package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm"
	"net/http"
	"time"
	_ "time"
)

var db = connect()

type User struct {
	Id              int        `json:"id"`
	Name            string     `json:"name"`
	First_last_name string     `json:"first_last_name"`
	Email           string     `json:"email"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
	DeletedAt       *time.Time `json:"deleted_at"`
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
	var u User
	err := json.NewDecoder(r.Body).Decode(&u)

	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	u.CreatedAt = time.Now().UTC()
	u.UpdatedAt = time.Now().UTC()
	u.DeletedAt = nil

	result := db.Create(&u)

	if result.Error == nil {
		json.NewEncoder(writer).Encode(map[string]string{"data": "ok"})
	} else {
		json.NewEncoder(writer).Encode(map[string]string{"data": "ko"})
	}
}

func update(writer http.ResponseWriter, request *http.Request) {

	vars := mux.Vars(request)
	id := vars["id"]

	var u User
	err := json.NewDecoder(request.Body).Decode(&u)

	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	u.UpdatedAt = time.Now().UTC()
	result := db.Model(&u).Where("id = ?", id).Update(&u)

	if result.Error == nil {
		json.NewEncoder(writer).Encode(map[string]string{"data": "ok"})
	} else {
		json.NewEncoder(writer).Encode(map[string]string{"data": "ko"})
	}
}

func delete(writer http.ResponseWriter, request *http.Request) {

	vars := mux.Vars(request)
	id := vars["id"]
	var user User

	result := db.Where("id = ?", id).Unscoped().Delete(&user)

	if result.Error == nil {
		json.NewEncoder(writer).Encode(map[string]string{"data": "ok"})
	} else {
		json.NewEncoder(writer).Encode(map[string]string{"data": "ko"})
	}
}

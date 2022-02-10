package handlers

import (
	"encoding/json"
	"net/http"
	"prueba/database"
	"prueba/models"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rs/xid"
	"golang.org/x/crypto/bcrypt"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	db := database.Connection()
	var users []models.User
	db.Find(&users)
	json.NewEncoder(w).Encode(users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	db := database.Connection()
	params := mux.Vars(r)
	var user models.User
	id, _ := strconv.ParseInt(params["id"], 10, 64)
	db.First(&user, id)

	if user.Id == "" {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("User not found")
	} else {
		json.NewEncoder(w).Encode(user)
	}

}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	db := database.Connection()
	var user models.User
	_ = json.NewDecoder(r.Body).Decode(&user)

	user.Id = xid.New().String()
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	user.Password = string(hashedPassword)

	db.Create(&user)
	json.NewEncoder(w).Encode(user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	db := database.Connection()
	params := mux.Vars(r)
	var user models.User
	id, _ := strconv.ParseInt(params["id"], 10, 64)
	db.First(&user, id)
	_ = json.NewDecoder(r.Body).Decode(&user)
	db.Model(&user).Updates(user)
	json.NewEncoder(w).Encode(user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	db := database.Connection()
	params := mux.Vars(r)
	var user models.User
	id, _ := strconv.ParseInt(params["id"], 10, 64)
	db.First(&user, id)
	db.Delete(&user)
	json.NewEncoder(w).Encode(user)
}

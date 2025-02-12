package main

import (
	"encoding/json"
	"net/http"
	"sync"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

var (
	users  = make(map[int]User)
	idUser = 1
	mu     sync.Mutex
)

func userHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	switch r.Method {
	case "GET":
		getAllUsers(w, r)
	case "POST":
		createUser(w, r)
	case "PUT":
		updateUser(w, r)
	default:
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
	}
}

// get all
func getAllUsers(w http.ResponseWriter, r *http.Request) {
	var userList []User

	for _, user := range users {
		userList = append(userList, user)
	}
	json.NewEncoder(w).Encode(userList)
}

// get by id
func getUserById(w http.ResponseWriter, r *http.Request) {

}

// post
func createUser(w http.ResponseWriter, r *http.Request) {
	var user User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	users[idUser] = user
	idUser++
	json.NewEncoder(w).Encode(user)
}

// put
func updateUser(w http.ResponseWriter, r *http.Request) {

}

// delete by id
func deleteUserById(w http.ResponseWriter, r *http.Request) {

}

// delete all
func deleteAllUser(w http.ResponseWriter, r *http.Request) {

}

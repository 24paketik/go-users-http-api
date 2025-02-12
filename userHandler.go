package main

import (
	"encoding/json"
	"net/http"
	"sync"
)

type User struct {
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
	case "DELETE":
		deleteUserById(w, r)
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
	var request struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if user, exists := users[request.ID]; exists {
		user.Name = request.Name
		users[request.ID] = user
		json.NewEncoder(w).Encode(user)
	} else {
		http.Error(w, "Пользователь не найден", http.StatusNotFound)
	}
}

// delete by id
func deleteUserById(w http.ResponseWriter, r *http.Request) {
	var userId int

	if err := json.NewDecoder(r.Body).Decode(&userId); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if _, exists := users[userId]; exists {
		delete(users, userId)
		response := map[string]string{"message": "Пользователь удален"}
		json.NewEncoder(w).Encode(response)
	} else {
		http.Error(w, "Пользователя нет", http.StatusNotFound)
	}
}

// delete all
func deleteAllUser(w http.ResponseWriter, r *http.Request) {

}

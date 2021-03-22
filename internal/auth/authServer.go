package auth

import (
	"encoding/json"
	"log"
	"net/http"
)

var database = make(map[string]string)

func AuthServer() {
	http.HandleFunc("/register", register)
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

type User struct {
	username string
	password string
}

type Msg struct {
	msg      string
	token    string
	imServer string
}

func register(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		w.WriteHeader(405)
		return
	}

	user, err := getUser(req)
	if err != nil {
		log.Fatal(err)
		return
	}

	if _, ok := database[user.username]; ok {
		sendMsg(w, 200, &Msg{"username has been registered", "", ""})
		return
	}

	token, err := GenerateToken(user.username, 1<<20)
	if err != nil {
		sendMsg(w, 409, &Msg{err.Error(), "", ""})
		return
	}

	database[user.username] = user.password

	sendMsg(w, 200, &Msg{"register success. auto login", token, "localhost:8081"})
	// w.Header().Add("x-auth-token", tokenString)
}

func login(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		w.WriteHeader(405)
		return
	}

	user, err := getUser(req)
	if err != nil {
		log.Fatal(err)
		return
	}

	if database[user.username] != user.password {
		sendMsg(w, 401, &Msg{"password incorrect", "", ""})
		return
	}

	token, err := GenerateToken(user.username, 1<<20)
	if err != nil {
		sendMsg(w, 409, &Msg{err.Error(), "", ""})
		return
	}
	sendMsg(w, 200, &Msg{"login success", token, "localhost:8081"})
}

func getUser(req *http.Request) (*User, error) {
	var user User
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func sendMsg(w http.ResponseWriter, statusCode int, msg *Msg) {
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(*msg)
	if err != nil {
		log.Fatal(err)
	}
}

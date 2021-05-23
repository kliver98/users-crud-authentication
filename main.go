package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	db "authentication/database"
	received "authentication/model/received"
	"authentication/service"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

var ENDPOINT string = "/api/latest/"

func (app *Main) indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "users-crud-authentication service for authentication")
}

func (app *Main) authenticate(w http.ResponseWriter, r *http.Request) {
	client, err := db.GetConnection()
	if err != nil {
		return
	}
	var userInfo received.User
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(reqBody, &userInfo)

	found, err := service.AuthenticateUser(client, received.User{Username: userInfo.Username, Password: userInfo.Password})
	if err != nil {
		setResponse(w, http.StatusNotAcceptable, err)
		return
	}
	setResponse(w, http.StatusOK, found)
}

func setResponse(w http.ResponseWriter, statusCode int, content interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(content)
}

func Init() *Main {
	main := &Main{}
	main.CreateRoutes()
	return main
}

func (main *Main) CreateRoutes() {
	router := mux.NewRouter().StrictSlash(true)
	main.router = router

	router.HandleFunc(ENDPOINT, main.indexRoute)
	router.HandleFunc(ENDPOINT+"users/auth/", main.authenticate).Methods("POST")
}

type Main struct {
	router *mux.Router
}

func (main *Main) Router() *mux.Router {
	return main.router
}

func main() {
	godotenv.Load(".env.development.local")
	app := Init()
	port := os.Getenv("PORT")
	if port == "" {
		port = "3001"
	}
	log.Fatal(http.ListenAndServe(":"+port, app.router))

}

package models

import (
	"net/http"
	"time"
)

// a forum user has an unique ID, username and password, and reflects time of creation and updating of user account
type user struct {
	ID        uint32 `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// create some sample user accounts in an array named "seed"
var seed = []user{
	{ID: 1, Username: "Spino", Password: "Orcatide"},
	{ID: 2, Username: "Flagman Carolla", Password: "Who15Two"},
	{ID: 3, Username: "Raptos", Password: "AkagamiNOShanks123"},
}

// function for a new user creating an account on the forum
func createUser(writer http.ResponseWriter, request *http.Request) {

}

func userLogin(writer http.ResponseWriter, request *http.Request) {

}

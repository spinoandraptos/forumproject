package models

import "time"

// a forum user has an unique ID, username and password, and reflects time of creation and updating of user account
type user struct {
	ID        uint32 `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

//create some sample user accounts in an array users
var users = []user{
	{ID: 1, Username: "Spino", Password: "Orcatide"},
	{ID: 2, Username: "Flagman Carolla", Password: "Who15Two"},
	{ID: 3, Username: "Raptos", Password: "AkagamiNOShanks123"},
}

//function for a new user creating an account on the forum
func createacc(User *user) {

}

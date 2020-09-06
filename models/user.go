package models

// UserList For this demo, we're storing the user list in memory
var UserList = []user{
	user{Username: "user1", Password: "pass1"},
	user{Username: "user2", Password: "pass2"},
	user{Username: "user3", Password: "pass3"},
}

type user struct {
	Username string `json:"username"`
	Password string `json:"-"`
}

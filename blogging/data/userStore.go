package data

type User struct {
	Username string
	Password string
	Token    string
}

var UsersStore = []User{
	{
		Username: "matwa",
		Password: "password",
		Token:    "matwa:password",
	},
}

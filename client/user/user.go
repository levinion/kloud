package user

type User struct {
	ID             string
	HashedPassword string
}

var Me = &User{
	ID: "maruka",
}

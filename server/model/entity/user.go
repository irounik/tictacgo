package entity

type User struct {
	Username       string `json:"username"`
	HashedPassword string `json:"hashedPassword"`
}

package main

// User is a struct that includes the
// attributes that a user carries
type User struct {
	ID                int    `json:"id"`
	Username          string `json:"username"`
	EncryptedPassword string `json:"encryptedPassword"`
}

package main

import (
	"fmt"
)

type User struct {
	ID         string
	SecretCode string
	Name       string
	Email      string
	Complaints []Complaint
}

type Complaint struct {
	ID       string
	UserID   string
	Title    string
	Summary  string
	Rating   int
	Resolved bool
}

func GenerateID() string {
	return fmt.Sprintf("%d", len(Users)+1)
}

func GenerateComplaintID() string {
	return fmt.Sprintf("%d", len(Complaints)+1)
}

func GenerateSecretCode() string {
	return fmt.Sprintf("%d", len(Users)+1000)
}

package main

import (
	"errors"
	"net/http"
)

func GetUserBySecretCode(secretCode string) (*User, error) {
	for _, user := range Users {
		if user.SecretCode == secretCode {
			return &user, nil
		}
	}

	return nil, errors.New("user not found")
}

func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		secretCode := r.Header.Get("Secret-Code")

		if _, err := GetUserBySecretCode(secretCode); err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	}
}

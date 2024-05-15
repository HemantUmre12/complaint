package main

import (
	"encoding/json"
	"net/http"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	user.ID = GenerateID()
	user.SecretCode = GenerateSecretCode()
	Users[user.ID] = user

	json.NewEncoder(w).Encode(user)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		SecretCode string `json:"secret_code"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	user, err := GetUserBySecretCode(input.SecretCode)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func SubmitComplaintHandler(w http.ResponseWriter, r *http.Request) {
	secretCode := r.Header.Get("Secret-Code")
	user, err := GetUserBySecretCode(secretCode)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var complaint Complaint
	if err := json.NewDecoder(r.Body).Decode(&complaint); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	complaint.ID = GenerateComplaintID()
	complaint.UserID = user.ID
	Complaints[complaint.ID] = complaint
	user.Complaints = append(user.Complaints, complaint)
	Users[user.ID] = *user

	json.NewEncoder(w).Encode(complaint)
}

func GetAllComplaintsForUserHandler(w http.ResponseWriter, r *http.Request) {
	secretCode := r.Header.Get("Secret-Code")
	user, err := GetUserBySecretCode(secretCode)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	json.NewEncoder(w).Encode(user.Complaints)
}

func GetAllComplaintsForAdminHandler(w http.ResponseWriter, r *http.Request) {
	// Work under progress...
}

func ViewComplaintHandler(w http.ResponseWriter, r *http.Request) {
	// Work under progress...
}

func ResolveComplaintHandler(w http.ResponseWriter, r *http.Request) {
	secretCode := r.Header.Get("Secret-Code")
	_, err := GetUserBySecretCode(secretCode)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	complaintID := r.URL.Query().Get("id")
	complaint, exists := Complaints[complaintID]
	if !exists {
		http.Error(w, "Complaint not found", http.StatusNotFound)
		return
	}

	complaint.Resolved = true
	Complaints[complaint.ID] = complaint

	json.NewEncoder(w).Encode(complaint)
}

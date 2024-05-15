package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Example request
	// 	curl -X POST \
	//   http://localhost:8080/register \
	//   -H 'Content-Type: application/json' \
	//   -d '{
	//     "name": "John",
	//     "email": "john@example.com"
	// }'

	// Example response
	// {"ID":"1","SecretCode":"1000","Name":"John","Email":"john@example.com","Complaints":null}
	http.HandleFunc("/register", RegisterHandler)

	// Example request
	// curl -X GET \
	// http://localhost:8080/login \
	// -H 'Content-Type: application/json' \
	// -d '{
	// 	"secret_code": "1000"
	// }'

	// Example response
	// {"ID":"1","SecretCode":"1000","Name":"John","Email":"john@example.com","Complaints":null}
	http.HandleFunc("/login", LoginHandler)

	// !! IMPORTANT: All the following endpoints expects 'Secret-Code' in the header
	// for authentication

	// Example Request
	// curl -X POST \
	// http://localhost:8080/submitComplaint \
	// -H 'Content-Type: application/json' \
	// -H 'Secret-Code: 1000' \
	// -d '{
	// 	"title": "Complaint title",
	// 	"summary": "Brief summary of the complaint",
	// 	"rating": 3
	// }'

	// Example Response
	// {"ID":"1","UserID":"1","Title":"Complaint title","Summary":"Brief summary of the complaint","Rating":3,"Resolved":false}
	http.HandleFunc("/submitComplaint", SubmitComplaintHandler)

	// curl -X GET \
	// http://localhost:8080/getAllComplaintsForAdmin \
	// -H 'Content-Type: application/json' \
	// -H 'Secret-Code: 1000'

	// {"1":{"ID":"1","UserID":"1","Title":"Complaint title","Summary":"Brief summary of the complaint","Rating":3,"Resolved":false}}
	http.HandleFunc("/getAllComplaintsForUser", GetAllComplaintsForUserHandler)

	// Yet to be implemented
	http.HandleFunc("/getAllComplaintsForAdmin", GetAllComplaintsForAdminHandler)

	// Yet to be implemented
	http.HandleFunc("/viewComplaint", ViewComplaintHandler)

	http.HandleFunc("/resolveComplaint", ResolveComplaintHandler)

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}

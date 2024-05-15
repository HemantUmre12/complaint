package main

// The Users and Complaints map acts as a temporary database
// for the purpose of this task
var (
	Users      = make(map[string]User)
	Complaints = make(map[string]Complaint)
)

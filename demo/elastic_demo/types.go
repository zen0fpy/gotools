package main

type User struct {
	Name    string
	Age     int64
	Married bool
}

type Employee struct {
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	Age       int      `json:"age"`
	About     string   `json"about"`
	Interests []string `json"interests"`
}

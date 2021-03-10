package model

import "fmt"

type User struct {
	Id       int    `db:"id" json:"id"`
	RealName string `db:"real_name" json:"real_name"`
	City     string `db:"city" json:"city"`
}

func GenerateUsers() []User {
	var Users []User

	for i := 0; i < 10000; i++ {
		item := User{
			Id:       i,
			RealName: fmt.Sprint("real_name_", i),
			City:     "test_city",
		}
		Users = append(Users, item)
	}
	return Users
}

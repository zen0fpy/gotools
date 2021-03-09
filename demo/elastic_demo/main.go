package main

import (
	"fmt"
	"github.com/olivere/elastic/v7"
)

func main() {

	serverConf := ServerConfig{
		Elastic: Elastic{
			Host: "localhost",
			Port: 9200,
		},
	}

	client := NewClient(serverConf)

	employee := Employee{
		FirstName: "Zhang",
		LastName:  "SAN",
		Age:       18,
		About:     "This is a body.",
		Interests: []string{"program", "author"},
	}

	var err error
	err = Add(client, "employees", employee)
	if err != nil {
		fmt.Print(err)
	}

	err = Delete(client, "employees", "QX4DFXgBHCru_1f7fmAi")
	if err != nil {
		fmt.Print(err)
	}

	doc := map[string]interface{}{
		"first_name": "Zhong",
		"age":        100,
		"About":      "good body.",
	}
	err = Update(client, "employees", "SH4aFXgBHCru_1f7RGAd", doc)
	if err != nil {
		fmt.Print(err)
	}

	err = Get(client, "employees", "SH4aFXgBHCru_1f7RGAd")
	if err != nil {
		fmt.Print(err)
	}

	var emp Employee
	err = List(client, "employees", emp, 0, 50)
	if err != nil {
		fmt.Print(err)
	}

	filterQuery := elastic.NewTermQuery("first_name", "Zhang")
	var emp1 Employee
	err = Search(client, "employees", emp1, filterQuery)
	if err != nil {
		fmt.Print(err)
	}

	err = DeleteIndex(client, "user")
	if err != nil {
		fmt.Print(err)
	}

}

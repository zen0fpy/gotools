package main

import (
	"context"
	"github.com/olivere/elastic/v7"
	"log"
	"math/rand"
)

const (
	Address = "http://localhost:9200"
)

type Person struct {
	Name    string `json:"name"`
	Age     int64  `json:"age"`
	Married bool   `json:"married"`
}

func main() {

	client, err := elastic.NewClient(elastic.SetURL(Address))
	if err != nil {
		log.Fatalln(err)
	}

	p := Person{
		Name:    "Li San",
		Age:     int64(rand.Int()),
		Married: true,
	}

	resp, err := client.Index().Index("user").BodyJson(p).Do(context.Background())
	if err != nil {
		log.Fatalf("add persion to es failed, %s\n", err.Error())
	}

	log.Printf("Index user %s, index %s, type %s\n", resp.Id, resp.Index, resp.Type)
}

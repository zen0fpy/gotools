package main

import (
	"fmt"
	"github.com/tal-tech/cds/pkg/ckgroup"
	"github.com/tal-tech/cds/pkg/ckgroup/config"
	"log"
)

type user struct {
	Id       int    `db:int`
	RealName string `db:"real_name"`
	City     string `db:"city"`
}

func generateUsers() []user {
	var users []user

	for i := 0; i < 10000; i++ {
		item := user{
			Id:       i,
			RealName: fmt.Sprint("real_name_", i),
			City:     "test_city",
		}
		users = append(users, item)
	}
	return users
}

func main() {

	var c = config.Config{
		ShardGroups: []config.ShardGroupConfig{
			{ShardNode: "tcp://127.0.0.1:9000"},
		},
	}

	group := ckgroup.MustCKGroup(c)
	users := generateUsers()
	err := group.InsertAuto(`insert into user (id,real_name,city) values (#{id},#{real_name},#{city})`, `id`, users)
	if err != nil {
		log.Fatal(err)

	}

}

package main

import (
	"context"
	"fmt"
	"github.com/globalsign/mgo/bson"
	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/core/stores/cache"
	"log"
)

const (
	ConfigFile = `H:\zen0fpy\gotools\demo\mongo_demo\config.yml`
)

type Config struct {
	Cache      cache.CacheConf
	Datasource string
}

func main() {
	var c Config
	logx.Disable()
	conf.MustLoad(ConfigFile, &c)

	// insert
	m := NewUserModel(c.Datasource, "user", c.Cache)
	for i := 0; i < 1; i++ {
		user := &User{
			ID:   bson.NewObjectId(),
			Name: fmt.Sprintf("Lssssisan ge%d", i),
			Age:  int64(i),
		}

		err := m.Insert(context.Background(), user)
		if err != nil {
			fmt.Println(err)
		}
	}

	// find
	objId := "6049f27c6f40cbb74f3c031b"
	var aUser *User
	aUser, err := m.FindOne(context.Background(), objId)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("user id: %s, name: %s\n", aUser.ID, aUser.Name)

	uUser := User{
		Age: 20,
		ID:  bson.ObjectIdHex(objId),
	}

	// update
	err = m.Update(context.Background(), &uUser)
	if err != nil {
		log.Fatalf("update err: %s\n", err.Error())
	}

	// delete
	err = m.Delete(context.Background(), objId)
	if err != nil {
		log.Fatalf("delete %s\n", err.Error())
	}

	// batch query

	// batch update

	// batch delete
}

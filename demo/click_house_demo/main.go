package main

import (
	"encoding/json"
	"fmt"
	"gotools/demo/click_house_demo/model"
	"log"
)

const (
	configFile = `H:\zen0fpy\gotools\demo\click_house_demo\config.yml`
)

func main() {

	model.InitDB(configFile)

	var err error
	// 查询
	querySql := `select id, real_name, city from user where city = ?`
	filterValue := "test_city"

	//多行查询
	data := &[]*model.User{}
	err = model.DB().GetQueryNode().QueryRows(data, querySql, filterValue)
	if err != nil {
		log.Fatal(err)
	}
	for _, user := range *data {
		fmt.Printf("%d %s, %s\n", user.Id, user.RealName, user.City)
	}

	//QueryRowsNoType
	result, err := model.DB().GetQueryNode().QueryRowsNoType(querySql, filterValue)
	checkErr(err)

	encodeObjs, err := json.Marshal(result)
	checkErr(err)

	err = json.Unmarshal(encodeObjs, data)
	checkErr(err)
	for _, u := range *data {
		fmt.Printf("id: %d, name: %s, city: %s\n", u.Id, u.RealName, u.City)
	}

	// 单行查询
	aUser := &model.User{}
	err = model.DB().GetQueryNode().QueryRow(aUser, querySql, filterValue)
	checkErr(err)
	fmt.Printf("id: %d, name: %s, city: %s\n", aUser.Id, aUser.RealName, aUser.City)

	//QueryRowNoType
	result1, err := model.DB().GetQueryNode().QueryRowNoType(querySql, filterValue)
	checkErr(err)

	enObj, err := json.Marshal(result1)
	checkErr(err)

	err = json.Unmarshal(enObj, aUser)
	checkErr(err)
	fmt.Printf("id: %d. name: %s, city: %s\n", aUser.Id, aUser.RealName, aUser.City)

	//流查询
	c := make(chan *model.User)
	err = model.DB().GetQueryNode().QueryStream(c, querySql, filterValue)
	checkErr(err)
	for u := range c {
		fmt.Printf("%d %s\n", u.Id, u.RealName)
	}

	//插入
	insertSql := `insert into user (id,real_name,city) values (#{id},#{real_name},#{city})`
	err = model.DB().InsertAuto(insertSql, `id`, model.GenerateUsers())
	checkErr(err)

	//全部节点

	for i, n := range model.DB().GetAllNodes() {
		fmt.Printf("node: %d, %s\n", i, n.GetHost())
	}

	//切片
	for i, s := range model.DB().GetAllShard() {
		fmt.Printf("shard: %d %s %s\n", i, s.GetShardConn().GetHost(), s.GetShardConn().GetUser())
	}

	// Exec
	updateSql := ` ALTER TABLE user UPDATE city = ? where id = ?`
	err = model.DB().GetQueryNode().Exec(updateSql, "广州", 9994)
	checkErr(err)

	// Exec
	deleteSql := `ALTER TABLE user DELETE where city = ? and id in(?, ?, ?)`
	err = model.DB().GetQueryNode().Exec(deleteSql, filterValue, 9997, 9996, 9995)
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

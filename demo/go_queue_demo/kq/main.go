package main

import (
	"encoding/json"
	"fmt"
	"github.com/tal-tech/go-queue/kq"
	"github.com/tal-tech/go-zero/core/conf"
	"log"
	"math/rand"
	"strconv"
	"time"
)

const (
	configFile = `H:\zen0fpy\gotools\demo\go_queue_demo\kq\config.yml`
)

func Consume() {
	var c kq.KqConf
	conf.MustLoad(configFile, &c)
	q := kq.MustNewQueue(c, kq.WithHandle(func(key, value string) error {
		fmt.Printf("%s=>%s\n", key, value)
		return nil
	}))

	defer q.Stop()
	q.Start()
}

type Message struct {
	Key     string `json:"key"`
	Value   string `json:"value`
	Payload string `json:"message"`
}

func Produce() {

	pusher := kq.NewPusher([]string{
		"localhost:9092",
	}, "kq")

	ticker := time.NewTicker(3 * time.Second)

	for i := 0; i < 3; i++ {
		select {
		case <-ticker.C:
			count := rand.Intn(100)
			m := Message{
				Key:     strconv.FormatInt(time.Now().UnixNano(), 10),
				Value:   fmt.Sprintf("%d,%d", i, count),
				Payload: fmt.Sprintf("%d,%d", i, count),
			}
			body, err := json.Marshal(m)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(string(body))
			if err := pusher.Push(string(body)); err != nil {
				log.Fatalf("push message failed, err: %s\n", err.Error())
			}
		}
	}
}

func main() {
	go Produce()
	go Consume()

	time.Sleep(15 * time.Second)

}

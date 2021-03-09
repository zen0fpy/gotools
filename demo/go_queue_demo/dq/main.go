package main

import (
	"fmt"
	"github.com/tal-tech/go-queue/dq"
	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/core/stores/redis"
	"log"
	"strconv"
	"time"
)

const (
	configFile = `H:\zen0fpy\gotools\demo\go_queue_demo\dq\config.yml`
)

type Config struct {
	Producer []dq.Beanstalk
	Consumer struct {
		Beanstalks []dq.Beanstalk
		Redis      redis.RedisConf
	}
}

func Produce(c Config) {
	producer := dq.NewProducer(c.Producer)

	for i := 1001; i < 1005; i++ {
		_, err := producer.Delay([]byte(strconv.Itoa(i)), 5*time.Second)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("send: %d, time: %s\n", i, time.Now().Format("20006-01-02 15:04:05"))
	}

}

func Consumer(c Config) {
	consumer := dq.NewConsumer(dq.DqConf{
		Beanstalks: c.Consumer.Beanstalks,
		Redis:      c.Consumer.Redis,
	})

	consumer.Consume(func(body []byte) {
		fmt.Printf("consume: %s, time: %s\n", string(body), time.Now().Format("20006-01-02 15:04:05"))

	})
}

func main() {

	var c Config
	conf.MustLoad(configFile, &c)

	go func() {
		Produce(c)
	}()

	go func() {
		Consumer(c)
	}()

	time.Sleep(10 * time.Second)

}

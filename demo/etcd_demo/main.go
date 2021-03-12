package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/clientv3/concurrency"
	"github.com/tal-tech/cds/cmd/dm/cmd/sync/config"
	"github.com/tal-tech/go-zero/core/conf"
	"log"
	"time"
)

const (
	confFile = `H:\zen0fpy\gotools\demo\etcd_demo\config.yml`
)

type Config struct {
	Etcd config.EtcdConf
}

func main() {

	var c Config
	conf.MustLoad(confFile, &c)

	// connect
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   c.Etcd.Hosts,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	key := "key"
	val := "value1"

	// watch
	go func() {
		rch := cli.Watch(context.Background(), key)
		for response := range rch {
			for _, ev := range response.Events {
				fmt.Printf("Type: %s, key:%s, value: %s\n", ev.Type, ev.Kv.Key, ev.Kv.Value)

			}

		}
	}()

	time.Sleep(time.Second)
	// put
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	_, err = cli.Put(ctx, key, val)
	cancel()

	cli.Put(context.Background(), key, "VALUE1XXXXXXXXXXXXxxxXXXXXXXXX")

	// lease for auto delete
	aLease, err := cli.Grant(context.TODO(), 1)
	if err != nil {
		log.Fatal(err)
	}

	cli.Put(context.Background(), key, "xxxxxxxxxxx", clientv3.WithLease(aLease.ID))
	//time.Sleep(time.Second* 7)

	// get key after lease
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	value1, err := cli.Get(ctx, key)
	if err != nil {
		log.Fatalf("fail to get %s\n", err.Error())
	}
	cancel()
	for _, ev := range value1.Kvs {
		fmt.Printf("key: %s , vaule: %s\n", ev.Key, ev.Value)
	}

	time.Sleep(10 * time.Second)

	// etd keepAlive
	key2 := "key2"
	value2 := "value2"
	cli.Put(context.Background(), key2, value2, clientv3.WithLease(aLease.ID))
	ch, kerr := cli.KeepAlive(context.Background(), aLease.ID)
	if kerr != nil {
		log.Fatal(kerr)
	}

	for {
		ka := <-ch
		fmt.Printf("ttl: %d\n", ka)
		break
	}

	// dislock
	// 创建两个单独会话用来演示
	s1, err := concurrency.NewSession(cli)
	if err != nil {
		log.Fatal(err)
	}
	defer s1.Close()
	m1 := concurrency.NewMutex(s1, "/my-lock/")

	s2, err := concurrency.NewSession(cli)
	if err != nil {
		log.Fatal(err)
	}
	defer s2.Close()
	m2 := concurrency.NewMutex(s2, "/my-lock")

	// 会话1获取锁
	if err := m1.Lock(context.TODO()); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("acquired lock for s1\n")

	m2Locked := make(chan struct{})
	go func() {
		defer close(m2Locked)

		if err := m2.Lock(context.TODO()); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("acquired lock for s2\n")
	}()

	if err := m1.Unlock(context.TODO()); err != nil {
		log.Fatal(err)
	}
	fmt.Println("released lock for s1")

	<-m2Locked
	fmt.Println("acquired lock for s2")

}

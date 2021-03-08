package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"log"
)

const (
	Address = "localhost:9092"
)

func main() {

	client, err := sarama.NewClient([]string{Address}, nil)
	if err != nil {
		log.Fatalf("failed to create client, err: %s\n", err.Error())
	}
	defer client.Close()

	topics, err := client.Topics()
	if err != nil {
		log.Fatalln(err)
	}

	for i, topic := range topics {
		fmt.Printf("topic: %d  name: %s\n", i, topic)
	}

	for _, topic := range topics {
		partitions, _ := client.Partitions(topic)
		for _, p := range partitions {
			leader, err := client.Leader(topic, p)
			if err != nil {
				log.Fatalf("query failed for leader")
			}
			fmt.Printf("leader add: %s, topic: %s, partition: %d\n", leader.Addr(), topic, p)

		}
	}

	brokers := client.Brokers()
	for i, b := range brokers {
		fmt.Printf("broker: %d addr: %s\n", i, b.Addr())
	}

}

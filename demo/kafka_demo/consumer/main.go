package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"log"
)

const (
	Address = "localhost:9092"
)

func SyncConsumeMessage() {

	fmt.Printf("Consume...\n")
	consumer, err := sarama.NewConsumer([]string{Address}, nil)
	if err != nil {
		log.Fatalf("fail to create consumer, %s\n", err.Error())
	}

	topic := "test"
	partitionList, err := consumer.Partitions(topic)
	if err != nil {
		log.Fatalf("failed to get partitions of topic %s\n", topic)
	}

	for {
		for _, p := range partitionList {
			pc, err := consumer.ConsumePartition(topic, p, sarama.OffsetNewest)
			if err != nil {
				log.Printf("consumer failed, topic %s, partition %d\n", topic, p)
				break
			}

			go func(pc sarama.PartitionConsumer) {
				defer pc.AsyncClose()
				for msg := range pc.Messages() {
					fmt.Printf("partition: %d, offset: %d, key: %s, value: %s\n",
						msg.Partition, msg.Offset, msg.Key, msg.Value)
				}
			}(pc)
		}
	}

}
func main() {
	SyncConsumeMessage()
}

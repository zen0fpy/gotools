package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"log"
	"time"
)

const (
	Address = "127.0.0.1:9092"
)

func AsyncProduceMessage() {

	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
	config.Producer.Timeout = time.Second
	config.Producer.Retry.Max = 1
	//config.Producer.MaxMessageBytes = 10

	producer, err := sarama.NewAsyncProducer([]string{Address}, config)
	if err != nil {
		log.Fatalf("failed to create producer， %s\n", err)
	}
	defer producer.Close()

	message := &sarama.ProducerMessage{
		Topic: "test",
		Key:   sarama.StringEncoder("hello"),
	}
	var value string
	for {
		fmt.Scanln(&value)
		message.Value = sarama.ByteEncoder(value)
		fmt.Printf("input [%s]\n", value)

		// send to chain
		producer.Input() <- message

		select {
		case succ := <-producer.Successes():
			fmt.Printf("offset: %d, timestamp: %s, partition: %d\n",
				succ.Offset, succ.Timestamp.Format("2006-01-02 15:04:05"), succ.Partition)
		case fail := <-producer.Errors():
			fmt.Printf("error: %s\n", fail.Err.Error())
		}
	}
}

func SyncProducMessage() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	config.Producer.Partitioner = sarama.NewRandomPartitioner

	msg := &sarama.ProducerMessage{
		Topic: "test",
		Key:   sarama.StringEncoder("sync_producer_key"),
		Value: sarama.StringEncoder("sync_producer_value_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
	}

	producer, err := sarama.NewSyncProducer([]string{Address}, config)
	if err != nil {
		log.Fatalln("failed to create producer.")
	}
	defer producer.Close()

	partitionId, offset, err := producer.SendMessage(msg)
	if err != nil {
		log.Fatalf("send message failed, err: %s\n", err.Error())
	}
	fmt.Printf("partition: %d, offset: %d\n", partitionId, offset)
}

func main() {
	AsyncProduceMessage()
	//SyncProducMessage()
}

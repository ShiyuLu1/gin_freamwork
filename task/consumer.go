package task

import (
	"context"
	"fmt"
	kafka "github.com/segmentio/kafka-go"
	"go_mvc/utils"
	"log"
)

func main() {
	// make a new reader that consumes from topic-A
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{"ip:port"},
		GroupID:  "consumer_group_id",
		Topic:    "consumer_group_topic",
		MinBytes: 10e3, // 10KB
		MaxBytes: 1e6,  // 1MB
	})

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			break
		}
		//fmt.Printf("message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
		msg := string(m.Value)
		data := utils.FormatMessage(msg)
		fmt.Println(data)
	}

	if err := r.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}
}

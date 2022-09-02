package messagebroker

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

func Publish(conn []string, topic string, key []byte, value []byte) error {

	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: conn,
		Topic:   topic,
		// assign the logger to the writer
		//Logger: mb.Logger,
	})

	err := w.WriteMessages(context.TODO(), kafka.Message{
		Key: key,
		// create an arbitrary message payload for the value
		Value: value,
	})
	if err != nil {
		return err
	}

	return nil
}

func Subscribe(conn []string, topic string) <-chan []byte {
	chValues := make(chan []byte)
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: conn,
		Topic:   topic,
		// assign the logger to the writer
		//Logger: mb.Logger,
	})
	//defer r.Close()
	go func() {
		for {
			msg, err := r.ReadMessage(context.TODO())
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(string(msg.Value))
				chValues <- msg.Value
			}
		}
		//close(chValues)
	}()
	return chValues
	//return msg.Key, msg.Value, nil
}

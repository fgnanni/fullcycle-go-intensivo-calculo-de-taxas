package kafka

import ckafka "github.com/confluentinc/confluent-kafka-go/kafka"

func Consume(topics []string, servers string, msgChan chan *ckafka.Message) {
	kfkaConsumer, err := ckafka.NewConsumer(&ckafka.ConfigMap{
		"bootstrap.servers": servers,
		"group.id":          "goapp",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		panic(err)
	}

	err = kfkaConsumer.SubscribeTopics(topics, nil)
	if err != nil {
		panic(err)
	}

	for {
		msg, err := kfkaConsumer.ReadMessage(-1)
		if err == nil {
			msgChan <- msg
		}
	}
}

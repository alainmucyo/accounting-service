package producer

import (
	"accounting-service/core/environment"
	"encoding/json"
	"errors"
	"github.com/Shopify/sarama"
	"log"
)

type Producer struct {
	env      *environment.Environment
	producer sarama.SyncProducer
}

func New(env *environment.Environment) *Producer {
	println("Starting Kafka producer")
	brokers := []string{env.KafkaBroker}
	config := sarama.NewConfig()
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true // Optional
	producer, err := sarama.NewSyncProducer(brokers, config)
	//defer producer.Close()
	if err != nil {
		println("Producer failed")
		log.Fatal(err)

	}
	return &Producer{env: env, producer: producer}
}

func (p *Producer) Produce(reqJSON interface{}, topic string) error {
	object, err := json.Marshal(reqJSON)
	if err != nil {

		return errors.New("invalid object")
	}
	message := sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(object),
	}
	_, _, err = p.producer.SendMessage(&message)
	if err != nil {
		return err
	}
	return nil
}

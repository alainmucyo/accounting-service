package consumer

import (
	"accounting-service/core/environment"
	"accounting-service/store/kafka/topics"
	"github.com/Shopify/sarama"
	"github.com/mistsys/sarama-consumer/offsets"
	"github.com/mistsys/sarama-consumer/stable"
	"sync"
	"time"
)
import consumer "github.com/mistsys/sarama-consumer"

type Consumer struct {
	client *consumer.Config
	env    *environment.Environment
	topics *topics.Topics
}

func New(env *environment.Environment, topics *topics.Topics) *Consumer {
	sconfig := sarama.NewConfig()
	sconfig.Version = consumer.MinVersion // consumer requires at least 0.9
	sconfig.Consumer.Return.Errors = true // needed if asynchronous ErrOffsetOutOfRange handling is desired (it's a good idea)

	// from that, create a consumer.Config
	config := consumer.NewConfig()
	config.Partitioner = stable.New(false)
	config.StartingOffset, config.OffsetOutOfRange = offsets.NoOlderThan(time.Second * 30)
	return &Consumer{env: env, client: config, topics: topics}
}

func (c *Consumer) Consume() {
	var wg sync.WaitGroup
	sclient, _ := sarama.NewClient([]string{c.env.KafkaBroker}, nil)
	// and finally a consumer Client
	client, _ := consumer.NewClient(c.env.KafkaGroupId, c.client, sclient)
	defer client.Close() // not strictly necessary, since we don't exit

	importedTopics := make([]string, 0, len(c.topics.List))
	for k := range c.topics.List {
		importedTopics = append(importedTopics, k)
	}
	// consume a topic
	topicsConsumers, _ := client.ConsumeMany(importedTopics)
	wg.Add(len(topicsConsumers))
	for _, topicConsumer := range topicsConsumers {
		go func(topicConsumer consumer.Consumer) {
			for msg := range topicConsumer.Messages() {
				topicConsumer.Done(msg) // required
				go c.topics.List[msg.Topic](msg.Value, msg.Topic)
			}
			wg.Done()
		}(topicConsumer)
	}
	wg.Wait()
}

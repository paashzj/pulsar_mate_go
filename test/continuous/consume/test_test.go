package main

import (
	"context"
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	topic   = "ci-consume"
	subName = "sub"
	content = "hello"
)

func TestConsume(t *testing.T) {
	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL: "pulsar://localhost:6650",
	})
	if err != nil {
		panic(err)
	}
	producer, err := client.CreateProducer(pulsar.ProducerOptions{
		Topic: topic,
	})
	if err != nil {
		panic(err)
	}
	_, err = producer.Send(context.Background(), &pulsar.ProducerMessage{
		Payload: []byte(content),
	})
	if err != nil {
		panic(err)
	}
	consumer, err := client.Subscribe(pulsar.ConsumerOptions{
		Topic:                       topic,
		SubscriptionName:            subName,
		SubscriptionInitialPosition: pulsar.SubscriptionPositionEarliest,
	})
	if err != nil {
		panic(err)
	}
	message, err := consumer.Receive(context.Background())
	if err != nil {
		panic(err)
	}
	assert.Equal(t, content, string(message.Payload()))
}

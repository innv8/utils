package utils

import (
	"encoding/json"
	"log"

	"github.com/streadway/amqp"
)

// QConnect
// Connects to rabbitmq
// Takes qURI
// Returns conn, error
func QConnect(qURI string) (conn *amqp.Connection, err error) {
	log.Printf("[q] connecting to rabbitmq at %s", qURI)

	conn, err = amqp.Dial(qURI)
	if err != nil {
		log.Printf("[q] unable to connect to rabbitmq because %v", err)
		return
	}
	log.Printf("[q] connected to rabbitmq")
	return
}

// ChannelConnect
// Connects to a channel
// Takes conn
// Returns channel, error
func ChannelConnect(conn *amqp.Connection) (channel *amqp.Channel, err error) {
	log.Printf("[q] connecting to rabbitmq channel")
	channel, err = conn.Channel()
	if err != nil {
		log.Printf("[q] unable to connect to channel because %v", err)
		return
	}
	log.Printf("[q] created channel")
	return
}

// QConsumer
// Start Consumer
// Takes prefetchCount, q, ack, channel
// Returns <- chan Delivery, error
func QConsumer(prefetchCount int, q string, ack bool, channel *amqp.Channel) (msgChan <-chan amqp.Delivery, err error) {
	log.Printf("[q] starting consuming from %s", q)
	if err = channel.Qos(prefetchCount, 0, false); err != nil {
		log.Printf("[q] unable to set QoS because %v", err)
		return
	}

	msgChan, err = channel.Consume(q, "", ack, false, false, false, nil)
	if err != nil {
		log.Printf("[q] unable to consume from queue because %v", err)
		return
	}
	log.Printf("[q] consuming from rabbitmq queue")
	return
}

// QPublish
// Publishes data to exchange
// Takes channel, exchange, routingKey, data,
// Returns error
func QPublish(channel *amqp.Channel, exchange, routingKey string, data interface{}) (err error) {
	log.Printf("[q] publish data to exchange %s, routing key %s", exchange, routingKey)
	byteBody, _ := json.Marshal(data)
	if err = channel.Publish(exchange, routingKey, false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        byteBody,
	}); err != nil {
		log.Printf("[q] unable to publish to exchange %s, routing key %s because %v", exchange, routingKey, err)
		return
	}
	log.Printf("[q] published to exchange %s, routing key %s", exchange, routingKey)
	return nil
}

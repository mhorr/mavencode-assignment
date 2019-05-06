package shared

import (
	"encoding/json"
	"log"

	"github.com/streadway/amqp"
)

type byteProcessor func([]byte)

// RabbitQueue handles setup and cleanup of the Rabbit queue
// for reading/writing Person objects
type RabbitQueue struct {
	conn  *amqp.Connection
	ch    *amqp.Channel
	queue amqp.Queue
}

func (r RabbitQueue) queueName() string {
	return r.queue.Name
}

// Cleanup closes connections, channels, etc.
func (r *RabbitQueue) Cleanup() {
	if r.ch != nil {
		r.ch.Close()
		r.ch = nil
	}
	if r.conn != nil {
		r.conn.Close()
		r.conn = nil
	}
}

func (r *RabbitQueue) setUpChannel() error {
	log.Printf("called setUpChannel()")
	var err error
	r.conn, err = amqp.Dial(GetConfig().Rabbit)
	if err != nil {
		log.Printf("%s: %s. Rabbit connection string: %s\n", err, "Failed to dial rabbit", GetConfig().Rabbit)
		return err
	}

	r.ch, err = r.conn.Channel()
	if err != nil {
		log.Printf("%s: %s\n", err, "Failed to open channel on connection.")
		return err
	}

	r.queue, err = r.ch.QueueDeclare(
		GetConfig().PersonChannel, // name
		false,                     // durable
		false,                     // delete when unused
		false,                     // exclusive
		false,                     // no-wait
		nil,                       // arguments
	)
	if err != nil {
		log.Printf("%s: %s. PersonChannel: %s\n", err, "Failed in QueueDeclare", GetConfig().PersonChannel)
		return err
	}
	return err // will be nil if we get to here
}

// RabbitListener is an object that will listen on the
// configured rabbitmq server and channel, and call the byteProcessor
// function on all received messages.
type RabbitListener struct {
	rq      RabbitQueue
	handler byteProcessor
}

// NewRabbitListener returns an object that will listen on the
// configured rabbitmq server and channel, and call the byteProcessor
// function on all received messages.
func NewRabbitListener(p byteProcessor) (RabbitListener, error) {
	l := RabbitListener{
		handler: p,
	}
	err := l.rq.setUpChannel()
	return l, err
}

// Listen begins listening on the channel.
// It does not return unless there is an error.
func (l *RabbitListener) Listen() error {
	err := l.rq.setUpChannel()
	if err != nil {
		return err
	}
	// now we should have a properly set up channel
	forever := make(chan bool)

	msgs, err := l.rq.ch.Consume(
		l.rq.queueName(), // queue
		"",               //consumer
		true,             // auto-ack
		false,            // exclusive
		false,            // no-local
		false,            // no-wait
		nil,              // args
	)
	go func() {
		for d := range msgs {
			l.handler(d.Body)
		}
	}()

	log.Printf(" Waiting for messages. To exit, press CTRL+C")
	<-forever
	l.rq.Cleanup() // try to cleanup before exiting
	return nil
}

//RabbitPersonWriter writes Person objects to a Rabbit queue
type RabbitPersonWriter struct {
	rq RabbitQueue
}

// GetRabbitPersonWriter returns an object that can write
// person objects to the rabbit queue.
func GetRabbitPersonWriter() (RabbitPersonWriter, error) {
	var w RabbitPersonWriter
	err := w.rq.setUpChannel()
	if err != nil {
		return w, err
	}
	return w, err
}

func (w *RabbitPersonWriter) Write(p Person) error {
	b, err := json.Marshal(p)
	if err != nil {
		return err
	}
	err = w.rq.ch.Publish(
		"",               // exchange
		w.rq.queueName(), // routing key
		false,            // mandatory
		false,            // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        b,
		},
	)
	return err
}

// Cleanup Cleans up the queue.
func (w *RabbitPersonWriter) Cleanup() {
	w.rq.Cleanup()
}

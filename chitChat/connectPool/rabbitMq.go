package connectPool

import (
	"context"
	"github.com/streadway/amqp"
	"io"
	"log"
	"time"
)

type Pool interface {
	Get(ctx context.Context) (io.Closer, error)
	Put(ctx context.Context, c io.Closer, forceClose bool) error
	Close() error
}

type Service struct {
}

var AmqbService Service

const (
	Exchange  = "test"
	QueueName = "test"
	Host      = "127.0.0.1:5672"
)

func Init() {
	conn, err := amqp.Dial("amqp://guest:guest@127.0.0.1:5672/")
	mqFail(err, "connect")
	defer conn.Close()

	//定义mq channel
	ch, err := conn.Channel()
	mqFail(err, "channel")
	defer ch.Close()

	qu, err := ch.QueueDeclare("hello", false, false, false, false, nil)
	mqFail(err, "queueDeclare")

	body := "{\"xxxxx\":\"xxxxxx\"}"
	err = ch.Publish("", qu.Name, false, false, amqp.Publishing{
		Headers:         nil,
		ContentType:     "application/json",
		ContentEncoding: "",
		DeliveryMode:    0,
		Priority:        0,
		CorrelationId:   "",
		ReplyTo:         "",
		Expiration:      "",
		MessageId:       "",
		Timestamp:       time.Time{},
		Type:            "",
		UserId:          "",
		AppId:           "",
		Body:            []byte(body),
	})
	mqFail(err, "Publish")
}

func mqFail(err error, msg string) {
	if err != nil {
		log.Fatalf("%s:%s", msg, err)
	}
}

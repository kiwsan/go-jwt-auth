package data

import (
	"fmt"
	"github.com/kiwsan/go-jwt-auth/utils"
	"github.com/streadway/amqp"
)

func RabbitMqConnection() (*amqp.Connection, *amqp.Channel, *amqp.Queue, error) {

	env := utils.Config.RabbitMq

	connectionString := fmt.Sprintf("amqp://%s:%s@%s:%s/", env.User, env.Password, env.Host, env.Port)

	conn, err := amqp.Dial(connectionString)
	if err != nil {
		return nil, nil, nil, err
	}
	//defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		return nil, nil, nil, err
	}
	//defer ch.Close()

	q, err := ch.QueueDeclare(
		"identityService", // name
		false,             // durable
		false,             // delete when unused
		false,             // exclusive
		false,             // no-wait
		nil,               // arguments
	)
	if err != nil {
		return nil, nil, nil, err
	}

	return conn, ch, &q, nil
}

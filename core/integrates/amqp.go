package integrates

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"practice/auth/core/config"
	"practice/auth/core/constants"
	"practice/auth/core/utils"
	"time"

	"github.com/streadway/amqp"
)

type BindingQueue struct {
	RoutingKey string
	Queues     []string
}

type Rabbit struct {
	Connection    *amqp.Connection
	Channel       *amqp.Channel
	BindingRoutes map[string]BindingQueue
}

var RMQConnection *amqp.Connection = nil

func ConnectRabbit(config config.AppConfig) *Rabbit {
	connectString := fmt.Sprintf(
		"amqp://%s:%s@%s:%s",
		config.RabbitSetting.User,
		config.RabbitSetting.Pass,
		config.RabbitSetting.Host,
		config.RabbitSetting.Port,
	)

	conn, err := amqp.Dial(connectString)
	if err != nil {
		slog.Info(err.Error())
		return nil
	} else {
		log.Println("🚀 Connected Successfully to RabbitMQ")
		RMQConnection = conn
	}

	ch, err := conn.Channel()
	if err != nil {
		slog.Info(err.Error())
		return nil
	}

	rqmInstance := &Rabbit{
		Connection: conn,
		Channel:    ch,
	}

	return rqmInstance
}

func (r *Rabbit) InitExchange(exchanges []string) {
	for _, value := range exchanges {
		r.DeclareTopicExchange(value)
	}
}

func (r *Rabbit) InitQueue(queues []string) {

	for _, value := range queues {
		r.QueueDeClare(value)
	}
}

func (r *Rabbit) BindingQueuesToExchange(queues []string, exchange, routingKey string) (err error) {
	err = r.BindingQueue(queues, exchange, routingKey, true)
	return
}

func (r *Rabbit) DeclareTopicExchange(name string) error {
	return r.Channel.ExchangeDeclare(
		name,                // name
		constants.RQMFanout, // type
		true,                // durable
		false,               // auto-deleted
		false,               // internal
		false,               // no-wait
		nil,                 // arguments
	)
}

func (r *Rabbit) QueueDeClare(name string) (err error) {
	_, err = r.Channel.QueueDeclare(
		name,  // name
		false, // durable
		false, // auto delete
		false, // exclusive
		false, // no wait
		nil,   // args
	)

	return
}

func (r *Rabbit) PublishMessage(exch, rKey string, message []byte) error {
	return r.Channel.Publish(
		exch,  // exchange name
		rKey,  // routing key
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			ContentType:  "application/json",
			DeliveryMode: amqp.Persistent,
			Timestamp:    time.Now(),
			Body:         message,
		},
	)
}

func (r *Rabbit) BindingQueue(queue []string, exchange, routingKey string, noWait bool) (err error) {
	if len(queue) <= 0 || exchange == "" {
		return
	}

	if r.BindingRoutes == nil {
		r.BindingRoutes = make(map[string]BindingQueue)
	}

	for _, q := range queue {

		if q == "" {
			continue
		}

		err = r.Channel.QueueBind(
			q,          // queue name
			routingKey, // routing key
			exchange,   // exchange name
			noWait,     // noWait
			nil,        // args
		)

		if err != nil {
			return
		}
	}

	bindingQueue := BindingQueue{}
	existedBindingQueue := []string{}
	existedExchange, ok := r.BindingRoutes[exchange]

	if ok {
		existedExchange.Queues = append(existedExchange.Queues, queue...)
	} else {
		bindingQueue.RoutingKey = routingKey
		existedBindingQueue = append(existedBindingQueue, queue...)
		r.BindingRoutes[exchange] = bindingQueue
	}

	return
}

func (r *Rabbit) StartConsume(queue string) (data any, err error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	msgs, err := r.Channel.Consume(
		queue, // queue
		"",    // consumer
		true,  // auto ack
		false, // exclusive
		false, // no local
		false, // no wait
		nil,   //args
	)

	if err != nil {
		return nil, err
	}

	go r.ConsumeData(ctx, msgs)

	return
}

func (r *Rabbit) ConsumeData(ctx context.Context, messages <-chan amqp.Delivery) {
	for msg := range messages {
		fmt.Println("message:", utils.ConvertByteToString(msg.Body))
	}
}

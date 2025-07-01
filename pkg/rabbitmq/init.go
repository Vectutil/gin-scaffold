package rabbitmq

import (
	"fmt"
	"gin-scaffold/internal/config"
	"github.com/streadway/amqp"
	"sync"
)

type RabbitMQ struct {
	ch   *amqp.Channel
	conn *amqp.Connection
}

var RabbitMQClient *RabbitMQ

func InitRabbitMQ() {
	once := sync.Once{}
	once.Do(func() {
		RabbitMQClient = &RabbitMQ{}
		rq := config.Cfg.RabbitMQ
		link := fmt.Sprintf("amqp://%s:%s@%s:%s/", rq.User, rq.Password, rq.Host, rq.Port)

		// 创建连接和通道
		var err error
		RabbitMQClient.conn, err = amqp.Dial(link)
		if err != nil {
			panic(err)
		}
		RabbitMQClient.ch, err = RabbitMQClient.conn.Channel()
		if err != nil {
			panic(err)
		}
	})
}

// DeferClose 关闭连接
func (r *RabbitMQ) DeferClose() {
	r.ch.Close()
	r.conn.Close()
}

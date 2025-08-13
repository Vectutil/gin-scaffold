package rabbitmq

import (
	"fmt"
	"gin-scaffold/internal/config"
	"gin-scaffold/pkg/logger"
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

// CheckAndReconnect 检查连接状态并在必要时重新连接
func (r *RabbitMQ) CheckAndReconnect() error {
	// 检查连接是否关闭
	if r.conn == nil || r.conn.IsClosed() {
		rq := config.Cfg.RabbitMQ
		link := fmt.Sprintf("amqp://%s:%s@%s:%s/", rq.User, rq.Password, rq.Host, rq.Port)

		// 重新创建连接
		var err error
		r.conn, err = amqp.Dial(link)
		if err != nil {
			return fmt.Errorf("重新连接失败: %v", err)
		}

		// 重新创建通道
		r.ch, err = r.conn.Channel()
		if err != nil {
			r.conn.Close()
			return fmt.Errorf("创建通道失败: %v", err)
		}

		logger.Logger.Info("RabbitMQ 连接已重新建立")
	}
	return nil
}

// DeferClose 关闭连接
func (r *RabbitMQ) DeferClose() {
	r.ch.Close()
	r.conn.Close()
}

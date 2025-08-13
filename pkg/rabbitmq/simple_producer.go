package rabbitmq

import (
	"fmt"
	"gin-scaffold/pkg/logger"
	"github.com/streadway/amqp"
)

// CreateQueue 创建普通队列
func (r *RabbitMQ) CreateQueue(qName string) (err error) {
	_, err = r.ch.QueueDeclare(
		qName, // 队列名称
		true,  // 非持久化（根据需求修改）
		false, // 不自动删除
		false, // 排他性
		false, // 不等待响应
		nil,   // 无额外参数
	)
	return
}

// SimplePush 推送消息
// SimplePush 推送消息到指定队列
func (r *RabbitMQ) SimplePush(qName string, msg []byte) (err error) {
	err = r.ch.Publish(
		"",    // 默认交换机（队列名即路由键）
		qName, // 路由键，与队列名称一致
		false, // 不强制要求路由（mandatory）
		false, // 不延迟发布
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        msg,
		})
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("推送消息失败: %v", err))
		return err
	}
	return nil
}

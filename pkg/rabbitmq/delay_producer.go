package rabbitmq

import (
	"fmt"
	"gin-scaffold/pkg/logger"
	"github.com/streadway/amqp"
	"strconv"
)

// CreateDelayQueue 创建延迟队列结构（死信交换机+目标队列+延迟队列）
func (r *RabbitMQ) CreateDelayQueue(qName string) (err error) {
	// 1. 创建死信交换机
	err = r.ch.ExchangeDeclare(
		qName+"_dlx", // 死信交换机名称
		"direct",     // 直连交换机
		true,         // 持久化
		false,        // 不自动删除
		false,        // 非内部使用
		false,        // 不等待响应
		nil,          // 无额外参数
	)
	if err != nil {
		return fmt.Errorf("创建死信交换机失败: %w", err)
	}

	// 2. 创建目标队列（最终消费的队列）
	_, err = r.ch.QueueDeclare(
		qName, // 目标队列名称
		true,  // 持久化
		false, // 不自动删除
		false, // 非排他性
		false, // 不等待响应
		nil,   // 无额外参数
	)
	if err != nil {
		return fmt.Errorf("创建目标队列失败: %w", err)
	}

	// 3. 绑定目标队列到死信交换机（路由键设为延迟队列名称）
	err = r.ch.QueueBind(
		qName,          // 目标队列名称
		qName+"_delay", // ✅ 关键修正：路由键为延迟队列名称
		qName+"_dlx",   // 死信交换机名称
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("绑定队列到交换机失败: %w", err)
	}

	// 4. 创建延迟队列（设置死信交换机）
	args := amqp.Table{
		"x-dead-letter-exchange":    qName + "_dlx",   // 死信交换机
		"x-dead-letter-routing-key": qName + "_delay", // 死信路由键（与自身队列名一致）
	}
	_, err = r.ch.QueueDeclare(
		qName+"_delay", // 延迟队列名称
		true,           // 持久化
		false,          // 不自动删除
		false,          // 非排他性
		false,          // 不等待响应
		args,           // 额外参数
	)
	if err != nil {
		return fmt.Errorf("创建延迟队列失败: %w", err)
	}

	logger.Logger.Info("延迟队列结构创建成功")
	return nil
}

// PublishDelayMessage 发送延迟消息（设置消息TTL）
func (r *RabbitMQ) PublishDelayMessage(qName string, body []byte, ttl int) error {
	return r.ch.Publish(
		"",             // 使用默认交换机
		qName+"_delay", // 路由到延迟队列
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
			Expiration:  strconv.Itoa(ttl), // TTL（毫秒）
		},
	)
}

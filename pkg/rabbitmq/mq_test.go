package rabbitmq

import (
	"context"
	"gin-scaffold/pkg/robot"
	"log"
	"testing"
	"time"

	"gin-scaffold/internal/config"
	"gin-scaffold/pkg/logger"
)

func TestDelay(t *testing.T) {
	// 初始化配置和日志
	config.InitConfig("E:\\workspace\\src\\gin-scaffold\\config.yaml")
	logger.InitLogger()
	InitRabbitMQ()

	rabbitMQ := RabbitMQClient
	queueName := "order_delay_queue"

	// 创建延迟队列结构
	err := rabbitMQ.CreateDelayQueue(queueName)
	if err != nil {
		log.Fatalf("创建延迟队列失败: %v", err)
		return
	}

	// 发送延迟消息（TTL设为正确的毫秒值）
	messages := []struct {
		body []byte
		ttl  int
	}{
		{[]byte("订单3001超时未支付，5秒后关闭"), 5000},
		{[]byte("订单3002超时未支付，10秒后关闭"), 10000},
		{[]byte("订单3003超时未支付，30秒后关闭"), 30000},
	}

	// 发送消息
	for _, msg := range messages {
		err = rabbitMQ.PublishDelayMessage(queueName, msg.body, msg.ttl)
		if err != nil {
			log.Printf("发送消息失败: %v", err)
		} else {
			log.Printf("发送延迟消息成功，TTL: %dms, 内容: %s", msg.ttl, string(msg.body))
		}
	}

	// 启动消费者（从目标队列消费）
	err = rabbitMQ.ConsumeDelayQueue(queueName, abcd)
	if err != nil {
		log.Fatalf("启动消费者失败: %v", err)
	}

	// 保持主线程运行（至少等待最长TTL时间）
	log.Println("主线程运行中，等待消息消费...")
	time.Sleep(40 * time.Second) // 等待30秒以上，确保所有消息过期
}

func abcd(bs []byte) error {
	robot.CallQWAssistant(context.Background(), string(bs)+time.Now().Format("2006-01-02 15:04:05"), robot.QWRobotMsgTypeText)
	//return errors.New("假装错误")
	return nil
}

func TestSimple(t *testing.T) {
	queueName := "test_simple_queue"
	// 初始化配置和日志
	config.InitConfig("E:\\workspace\\src\\gin-scaffold\\config.yaml")
	logger.InitLogger()
	InitRabbitMQ()
	rabbitMQ := RabbitMQClient
	rabbitMQ.CreateQueue(queueName)
	rabbitMQ.Consumer(queueName, abcd)
	for {
		time.Sleep(2 * time.Second)
		rabbitMQ.SimplePush(queueName, []byte("hello world "))
	}
}

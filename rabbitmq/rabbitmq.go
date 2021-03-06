package rabbitmq

import (
	"encoding/json"
	"fmt"
	"log"
	"rabbitmq-miaosha/datamodels"

	"github.com/streadway/amqp"
)

const MQURL = "amqp://admin:123456@192.168.1.175:5672/blue"

type RabbitMQ struct {
	conn      *amqp.Connection
	channel   *amqp.Channel
	QueueName string
	Exchange  string
	Key       string
	Mqurl     string
}

// 创建实例
func NewRabbitMQ(queueName string, exchange string, key string) *RabbitMQ {
	rabbitmq := &RabbitMQ{
		QueueName: queueName,
		Exchange:  exchange,
		Key:       key,
		Mqurl:     MQURL,
	}
	var err error
	rabbitmq.conn, err = amqp.Dial(rabbitmq.Mqurl)
	rabbitmq.failOnErr(err, "创建连接错误")
	rabbitmq.channel, err = rabbitmq.conn.Channel()
	rabbitmq.failOnErr(err, "获取channel错误")

	return rabbitmq
}

// 断开channel 和 connection
func (r *RabbitMQ) Destory() {
	r.channel.Close()
	r.conn.Close()
}

// 错误处理函数
func (r *RabbitMQ) failOnErr(err error, message string) {
	if err != nil {
		log.Fatalf("%s:%s", message, err)
		panic(fmt.Sprintf("%s:%s", message, err))
	}
}

// 创建简单模式下的RabbitMQ实例
func NewRabbitMQSimple(queueName string) *RabbitMQ {
	return NewRabbitMQ(queueName, "", "")
}

// 简单模式下生产代码
func (r *RabbitMQ) PublishSimple(message string) {
	// 申请队列 如果队列不存在则自动创建 如果存在则跳过创建  可以保证队列存在  消息可以发送到队列里
	_, err := r.channel.QueueDeclare(
		r.QueueName,
		false, // 是否持久化
		false, // 是否自动删除
		false, // 是否独占
		false, // 是否阻塞
		nil,
	)
	if err != nil {
		fmt.Println(err)
	}

	// 发送消息到队列
	r.channel.Publish(
		r.Exchange,
		r.QueueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
}

// 简单模式下消费消息
func (r *RabbitMQ) ConsumeSimple() {
	_, err := r.channel.QueueDeclare(
		r.QueueName,
		false, // 是否持久化
		false, // 是否自动删除
		false, // 是否独占
		false, // 是否阻塞
		nil,
	)
	if err != nil {
		fmt.Println(err)
	}
	//消费者流控
	r.channel.Qos(
		1,     //当前消费者一次能接受的最大消息数量
		0,     //服务器传递的最大容量（以八位字节为单位）
		false, //如果设置为true 对全局channel可用
	)
	msgs, err := r.channel.Consume(
		r.QueueName,
		"",    //区分多个消费者
		false, //是否自动应答处理完成消息
		false, //是否独占
		false, // 同一个connection 中发送的消息 可传递给 同一个connection中的消费者
		false, // 是否阻塞
		nil,   // 额外参数
	)

	if err != nil {
		fmt.Println(err)
	}

	forever := make(chan bool)
	// 启用协程处理消息
	go func() {
		for d := range msgs {
			log.Printf("receive a msg: %s", d.Body)
			message := &datamodels.Message{}
			err := json.Unmarshal([]byte(d.Body), message)
			if err != nil {
				fmt.Println(err)
			}
			// 之后再执行数据库 插入订单 和 扣除商品数量 的逻辑

			//如果为true表示确认所有未确认的消息， 为false表示确认当前消息
			d.Ack(false)
		}
	}()

	log.Printf("waiting for msgs, to exit press ctrl+c")
	<-forever
}

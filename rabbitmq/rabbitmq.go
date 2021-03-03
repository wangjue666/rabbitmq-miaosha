package rabbitmq

import "github.com/streadway/amqp"

const MQURL = "amqp://admin:admin@192.168.1.175:5672/blue"

type RabbitMq struct {
	conn      *amqp.Connection
	channel   *amqp.channel
	QueueName string
	Exchange  string
	Key       string
	Mqurl     string
}

// 创建实例
func NewRabbitMQ(queueName string, exchange string, key string) *RabbitMQ {
	rabbitmq := &RabbitMQ{
		QueueName: queueName,
		Exchange: exchange,
		Key: key,
		Mqurl: MQURL
	}
	var err error
	rabbitmq.conn, err = amqp.Dial(rabbitmq.Mqurl)
	rabbitmq.failOnErr(err, "创建连接错误")
	rabbitmq.channel, err = rabbitmq.conn.Channel()
	rabbitmq.failOnErr(err, "获取channel错误")

	return rabbitmq
}

// 断开channel 和 connection
func (r *RabbitMq)Destory(){
	r.channel.Close()
	r.conn.Close()
}

// 错误处理函数
func (r *RabbitMQ)failOnErr(err error, message string){
	if(err != nil){
		log.Fatalf("%s:%s", message, err)
		panic(fmt.Sprintf("%s:%s", message, er))
	}
}

// 创建简单模式下的RabbitMQ实例
func NewRabbitMQSimple(queueName string) *RabbitMQ{
	return NewRabbitMQ(queueName, "", "")
}
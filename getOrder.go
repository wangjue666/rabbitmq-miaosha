package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"rabbitmq-miaosha/datamodels"
	RabbitMQ "rabbitmq-miaosha/rabbitmq"
	"sync"
	"time"
)

var sum int64 = 0

//预存商品数量
var productNum int64 = 1000000

//互斥锁
var mutex sync.Mutex

//计数
var count int64 = 0

//获取秒杀商品
func GetOneProduct() bool {
	//加锁
	mutex.Lock()
	defer mutex.Unlock()
	count += 1
	//判断数据是否超限
	if count%100 == 0 {
		if sum < productNum {
			sum += 1
			fmt.Println("预存商品总量", productNum, "已消费数量", sum)
			return true
		}
	}
	return false

}

func GetProduct(w http.ResponseWriter, req *http.Request) {
	if GetOneProduct() {
		w.Write([]byte("true"))
		return
	}
	w.Write([]byte("false"))
	return
}

var rabbitmq = RabbitMQ.NewRabbitMQSimple("blue")

func GetOrder(w http.ResponseWriter, req *http.Request) {
	const userId = 9527
	const productId = 10086
	timeUnix := time.Now().Unix()
	//创建消息体
	message := datamodels.NewMessage(userId, productId, timeUnix)
	byteMessage, err := json.Marshal(message)
	if err != nil {
		panic(err)
	}

	rabbitmq.PublishSimple(string(byteMessage))

	w.Write([]byte("true"))
}

func main() {
	http.HandleFunc("/getOne", GetProduct)
	http.HandleFunc("/getOrder", GetOrder)
	err := http.ListenAndServe(":8084", nil)
	if err != nil {
		log.Fatal("Err:", err)
	}
}

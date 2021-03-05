# rabbitmq-miaosha
 基于rabbitmq秒杀系统基本架构, 通过异步下单, 完美解决高并发场景下的商品超卖,暴库


### install rabbitMQ 
[Centos7安装RabbitMQ最新版3.8.5](https://blog.csdn.net/weixin_40584261/article/details/106826044)


### Wrk压测工具安装

```shell 
git clone https://github.com/wg/wrk 
```


### 使用说明
```shell 
// 生产端
go run getOrder.go

// 消费端
go run consumer.go
```

### 其他秒杀优化思路

+ 使用CDN，权限验证尽可能地拦截流量
+ 启用消息队列流量削峰机制
+ 增加IP限流策略 如漏桶法限流,设置异常流量黑名单
+ 秒杀开始后 后台再返回秒杀接口地址
+ 更高级点儿,像小米一样 前台不请求后台接口或后台随机丢弃请求 直接提示抢购失败, 这两种从数学期望角度看 并不影响抢购的公平性。
+ 多个秒杀商品可以分批次秒杀
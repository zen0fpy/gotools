### 消息队列系统的设计


#### 生产者代理
1. 内部保证消息不会丢失、不重复投递
2. 不阻塞上游业务，本地保存
3. 发送失败，要failback


#### meta center
1. etcd
2. zookeeper
3. redis-cluster


#### 队列代理
1. kafka
2. redis
3. rocket-mq
4. nsq



#### 消费者代理
1. 不重复消费
2. 需要保存消费记录
3. 保存消费记录的节点，高可用，持久化
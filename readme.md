# Raft-Simple
* 一个raft简单实现，只实现kv的存储，kv均为字符串类型
* A simple raft implementation。Only support key(string) value(string) 

## Server

### 效率优化
* Leader的日志同步采用批量方式，100条或10ms内无新请求进入，就触发同步消息
* 心跳机制，每次批处理最多等到100ms，防止出发超时导致领导者选举的发生
* 消息接收与消息同步采用异步方式处理
* Leader给从节点发送消息采用并发
* 防止消息错乱，打造流水线式的处理

### 节点通信
* 采用grpc
*  心跳or日志同步

### 持久化
* 简化存储，每个节点分为日志索引和日志内容
* 文件1：
    * 文件头 RAFT01  RAFT标识+两位版本号
    * 日志个数 uint64 存储
    * 日志内容索引 uint64 uint32, uint64 uint32，分别代表kv的索引地址和内容长度
* 文件2：
    *  存储kv日志内容，根据文件1中日志内容索引进行查询和同步

### 日志复制
* Leader通过心跳将更新消息发送给Follower，Follower收到消息后将日志存到自己的日志
* Leader收到大多数复制成功请求后，在自己节点更新状态机。并给客户端返回结果
* Follower发现领导者已提交了某条日志项，自己还没应用，立即将这条日志项提交更新本地状态机


### 超时机制
* 心跳超时：150-300ms内，Follower没收到Leader心跳的话，转变为Candidate，发起选举投票
* 选举超时：150-300ms内，Candidate未收到过半(包括自己)的选票的话，选举失败，增加任期重新发起选举

### 节点扩容
* 单节点扩容(待实现)，测试暂时使用3节点
* 日志全量复制及重放(待实现)

## Client
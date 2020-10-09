![中文文档](./readme_CN.md)
# Raft-Simple
* A simple raft implementation。Only support key(string) value(string) 

## Server

### Efficiency Optimization
* Send leader's Log on batch, act on 100 message or 10ms no more new message
* heartbeat, at most wait for 100ms, avoid to elect cause by timeout
* async on message receiving and message replicating
* Leader send messages in parallel
* Avoid to message disorder， using pipeline processing

### Communication between Nodes
* Communication by GRPC between nodes
* Heartbeat or Log replication
* Election

### Persistence
* Only storage Log Index and Log Content 
* file1：
    * Header RAFT01  "RAFT"+2 byte version No
    * Log count storage by uint64
    * Log Index, uint64 uint32, uint64 uint32，means kv index address and content length
* file2：
    *  storage kv Log Content, connected with file1 

### Log Replication
* Leader send message to Follower by heartbeat, Follower save Log 
* Leader receives responses from most Followers, update state machine on local, then response client a success
* Follower commit the Log to state machine when it realized this Log was committed by Leader


### Timeout
* Heartbeat Tiimeout： 150-300ms，Follower will translate to Candidate, if the heartbeat was not received in time
* Election Timeout: 150-300ms, Candidate will add the term and launch an election again, if Candidate didn't receive more than half of the votes 

### Node Expansion
* single node expansion(support later), test on 3 nodes now
* Log full copy and replay(support later)

## Client
package node

import (
	"github.com/wangyanga9/raft-simple/enum"
	"math/rand"
	"sync"
	"time"
)

var allNodePort = map[int]string{
	0: "8001",
	1: "8002",
	2: "8003",
}

var NodeIns *Node

type KVPair struct {
	term uint64
	index uint64
	key string
	value string
}

type Node struct {
	Role enum.NodeRole
	s sync.RWMutex
	Port string
	LeaderPort string
	HeartbeatTimeout time.Duration //心跳超时
	ElectionTimeout time.Duration //选举超时
	VotedTerm int64

	StateMachine map[string]string  //状态机
	SubmittedIndex int64  //已提交的
	Logs []KVPair  //日志，后续需要落盘
}

func NewNode(port string) *Node {
	rand.Seed(time.Now().Unix())
	return &Node{
		s:              sync.RWMutex{},
		Role:           enum.Follower,
		Port:           port,
		HeartbeatTimeout: time.Duration(rand.Intn(150) + 150) * time.Millisecond,
		ElectionTimeout: time.Duration(rand.Intn(150) + 150) * time.Millisecond,
		VotedTerm: 0,
		StateMachine:   make(map[string]string),
		SubmittedIndex: 0,
		Logs:           make([]KVPair, 0),
	}
}


//func(n *Node)
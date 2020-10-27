package node

import (
	"errors"
	"github.com/wangyanga9/raft-simple/enum"
	"math/rand"
	"sync"
	"time"
)

//var allNodePort = map[int]string{
//	0: "8001",
//	1: "8002",
//	2: "8003",
//}

var NodeIns *Node

type Peer struct {
	IP string
	Port string
}

func InitNode(port string, peerPorts ...string) (err error) {
	NodeIns, err = NewNode(port, peerPorts)
	return err
}

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

	Peers []string
	HeartbeatTimeout time.Duration //心跳超时
	ElectionTimeout time.Duration //选举超时
	VotedTerm int64

	StateMachine map[string]string  //状态机
	SubmittedIndex int64  //已提交的
	Logs []KVPair  //日志，后续需要落盘
}

func NewNode(port string, peerPorts []string) (*Node, error) {
	var err error
	if len(peerPorts) < 2 {
		err = errors.New("raft need at least 3 peer")
	}

	rand.Seed(time.Now().Unix())
	heartbeatTimeout := time.Duration(rand.Intn(150) + 150)
	electionTimeout := time.Duration(rand.Intn(150) + 150)
	for heartbeatTimeout > electionTimeout {
		heartbeatTimeout = time.Duration(rand.Intn(150) + 150)
		electionTimeout = time.Duration(rand.Intn(150) + 150)
	}
	return &Node{
		s:              sync.RWMutex{},
		Role:           enum.Follower,
		Port:           port,
		Peers:			peerPorts,
		HeartbeatTimeout: heartbeatTimeout * time.Millisecond,
		ElectionTimeout: electionTimeout * time.Millisecond,
		VotedTerm: 0,
		StateMachine:   make(map[string]string),
		SubmittedIndex: 0,
		Logs:           make([]KVPair, 0),
	}, err
}


//func(n *Node)
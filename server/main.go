package main

import (
	"flag"
	"fmt"
	"github.com/wangyanga9/raft-simple/common"
	"github.com/wangyanga9/raft-simple/protos"
	"github.com/wangyanga9/raft-simple/server/node"
	"github.com/wangyanga9/raft-simple/server/rpc"
	"os"
)

func main() {
	var port string
	flag.StringVar(&port, "p", "", "port")
	flag.Parse()
	if port == "" {
		fmt.Printf("config file path error:%s, start node with -p port", port)
		os.Exit(-1)
	}
	done := make(chan bool, 1)

	node.NodeIns = node.NewNode(port)
	err := initRpcServer(port)
	if err != nil {
		os.Exit(-1)
	}
	<-done
}

func initRpcServer(port string) error {
	grpcServer, err := common.NewGRPCServer("0.0.0.0:" + port, common.ServerConfig{})
	if err != nil {
		return fmt.Errorf("api GRPCServer failed: %s", err.Error())
	}
	protos.RegisterRaftSupportServer(grpcServer.Server(), &rpc.RpcServer{})
	go func() {
		grpcServer.Start()
	}()

	return nil
}

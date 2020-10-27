package main

import (
	"flag"
	"fmt"
	"github.com/wangyanga9/raft-simple/server"
	"github.com/wangyanga9/raft-simple/server/node"
	"os"
)

func main() {
	// 启动sever
	var port string
	flag.StringVar(&port, "p", "", "port")
	flag.Parse()
	if port == "" {
		fmt.Printf("config file path error:%s, start node with -p port", port)
		os.Exit(-1)
	}
	done := make(chan bool, 1)


	err := server.InitRpcServer(port)
	if err != nil {
		os.Exit(-2)
	}
	// 启动 node 
	err = node.InitNode(port)
	if err != nil {
		os.Exit(-3)
	}

	<-done
}

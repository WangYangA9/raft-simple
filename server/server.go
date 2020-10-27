package server

import (
	"fmt"
	"github.com/wangyanga9/raft-simple/common"
	"github.com/wangyanga9/raft-simple/protos"
	"github.com/wangyanga9/raft-simple/server/rpc"
)



func InitRpcServer(port string) error {
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

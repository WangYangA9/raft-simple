package rpc

import (
	"context"
	"github.com/wangyanga9/raft-simple/protos"
)

type RpcServer struct {
	handler *RpcHandler
}

func (server *RpcServer) Agent(ctx context.Context, req *protos.RaftMsg) (*protos.RaftMsg, error) {
	return server.handler.HandleMsg(req)
}

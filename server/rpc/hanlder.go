package rpc

import (
	"github.com/wangyanga9/raft-simple/enum"
	"github.com/wangyanga9/raft-simple/protos"
	"github.com/wangyanga9/raft-simple/server/node"
	"time"
)

const GrpcTimeout = 150 * time.Millisecond

type RpcHandler struct {
}

func (h *RpcHandler) HandleMsg(req *protos.RaftMsg) (*protos.RaftMsg, error) {
	switch req.Type {
	case protos.RaftMsg_HEARTBEAT:
		return h.handleHeartBeat(req)
	case protos.RaftMsg_VOTE:
		return h.handleVote(req)
	case protos.RaftMsg_GET:
		return h.handleVote(req)
	case protos.RaftMsg_SET:
		return h.handleVote(req)
	default:
		return &protos.RaftMsg{Type: protos.RaftMsg_ERROR, Payload: []byte("Unknown msg type!")}, nil
	}
	//return nil, errors.New("unknown error in handle rpc message")
}

func (h *RpcHandler) handleHeartBeat(req *protos.RaftMsg) (*protos.RaftMsg, error) {
	switch node.NodeIns.Role {
	case enum.Follower:
	case enum.Candidate:
		return &protos.RaftMsg{Type: protos.RaftMsg_ERROR, Payload: []byte("ELECTING")}, nil
	case enum.Leader:
		return &protos.RaftMsg{Type: protos.RaftMsg_ERROR, Payload: []byte("ELECTING")}, nil
	}
	return &protos.RaftMsg{Type: protos.RaftMsg_ERROR, Payload: []byte("UNKNOWN NODE ROLE")}, nil
}


func (h *RpcHandler) handleVote(req *protos.RaftMsg) (*protos.RaftMsg, error) {
	switch node.NodeIns.Role {
	case enum.Follower:

	case enum.Candidate:
		return &protos.RaftMsg{Type: protos.RaftMsg_ERROR, Payload: []byte("ELECTING")}, nil
	case enum.Leader:
		return &protos.RaftMsg{Type: protos.RaftMsg_ERROR, Payload: []byte("ELECTING")}, nil
	}

	return &protos.RaftMsg{Type: protos.RaftMsg_ERROR, Payload: []byte("UNKNOWN NODE ROLE")}, nil
}

func (h *RpcHandler) handleGet(req *protos.RaftMsg) (*protos.RaftMsg, error) {
	switch node.NodeIns.Role {
	case enum.Follower:
		return &protos.RaftMsg{Type: protos.RaftMsg_RESPONSE, Payload: []byte(node.NodeIns.LeaderPort)}, nil
	case enum.Candidate:
		return &protos.RaftMsg{Type: protos.RaftMsg_ERROR, Payload: []byte("ELECTING")}, nil
	case enum.Leader:
		//TODO  状态机查询
	}

	return &protos.RaftMsg{Type: protos.RaftMsg_ERROR, Payload: []byte("UNKNOWN NODE ROLE")}, nil
}

func (h *RpcHandler) handleSet(req *protos.RaftMsg) (*protos.RaftMsg, error) {
	switch node.NodeIns.Role {
	case enum.Follower:
		return &protos.RaftMsg{Type: protos.RaftMsg_RESPONSE, Payload: []byte(node.NodeIns.LeaderPort)}, nil
	case enum.Candidate:
		return &protos.RaftMsg{Type: protos.RaftMsg_ERROR, Payload: []byte("ELECTING")}, nil
	case enum.Leader:
		//TODO
		// 日志记录
		// 发送同步
		// 日志提交到状态机
	}

	return &protos.RaftMsg{Type: protos.RaftMsg_ERROR, Payload: []byte("UNKNOWN NODE ROLE")}, nil
}
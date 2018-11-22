package main

import (
	"context"
	"encoding/json"
	"github.com/godcong/qtrago/proto"
	"github.com/godcong/qtrago/util"
	"log"
)

type qtra struct {
}

func (*qtra) MessageInfo(ctx context.Context, r *proto.MessageRequest) (*proto.MessageReply, error) {
	select {
	case <-ctx.Done():
		log.Println("exit", ctx.Err())
	}
	maps := util.Map{}
	err := json.Unmarshal([]byte(r.Json), &maps)
	i := util.Map{
		"status": "ok",
	}
	if err != nil {
		i.Set("status", "failed")
		reply := &proto.MessageReply{
			Json: string(i.ToJSON()),
		}
		return reply, err
	}

	reply := &proto.MessageReply{
		Json: string(i.ToJSON()),
	}
	return reply, nil
}

func (*qtra) Trade(ctx context.Context, r *proto.TradeRequest) (*proto.TradeReply, error) {
	select {
	case <-ctx.Done():
		log.Println("exit", ctx.Err())
	}
	maps := util.Map{}
	err := json.Unmarshal([]byte(r.Json), &maps)
	i := util.Map{
		"status": "ok",
	}
	if err != nil {
		i.Set("status", "failed")
		reply := &proto.TradeReply{
			Json: string(i.ToJSON()),
		}
		return reply, err
	}

	reply := &proto.TradeReply{
		Json: string(i.ToJSON()),
	}
	return reply, nil
}

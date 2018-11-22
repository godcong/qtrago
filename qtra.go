package main

import (
	"context"
	"encoding/json"
	"github.com/godcong/qtrago/proto"
	"github.com/godcong/qtrago/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type qtra struct {
	*grpc.Server
}

func (q *qtra) Start() {
	log.Println("start grpc server")
	q.Server = grpc.NewServer()
	proto.RegisterQuantitativeTradingServer(q.Server, q)
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	//cli need
	reflection.Register(q.Server)

	if err := q.Server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (*qtra) MessageInfo(ctx context.Context, r *proto.MessageRequest) (*proto.MessageReply, error) {
	log.Println("message")

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
	log.Println("trade")

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

package main

import (
	"log"
	"net"
	"fmt"

	pbEmpty "github.com/golang/protobuf/ptypes/empty"
	pbMsg "github.com/toshim45/grpclab/messaging"
	pbMsgV1 "github.com/toshim45/grpclab/messagingv1"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":8081"
)

type MessageService struct {
	pbMsg.MessageServiceServer
}

func (ms *MessageService) UpdateMessage(ctx context.Context, req *pbMsg.MessageRequest) (res *pbMsg.MessageResponse, err error) {
	log.Printf("got text %s\n", req.Label)
	return &pbMsg.MessageResponse{Label: req.Label, Created: 1000}, nil
}

func (ms *MessageService) CheckMessage(ctx context.Context, req *pbEmpty.Empty) (res *pbMsg.CheckResponse, err error) {
	log.Printf("check\n")
	return &pbMsg.CheckResponse{Status: pbMsg.ServingStatus_SERVING}, nil
}

func (ms *MessageService) ListServingStatus(ctx context.Context, req *pbEmpty.Empty) (res *pbMsg.MapServingStatusResponse, err error) {
	return &pbMsg.MapServingStatusResponse{
		Status: pbMsg.ServingStatus_name,
	}, nil
}

// v1

type MessageServiceV1 struct {
	pbMsgV1.MessageServiceServer
}

func (ms *MessageServiceV1) UpdateMessage(ctx context.Context, req *pbMsgV1.MessageRequest) (res *pbMsgV1.MessageResponse, err error) {
	log.Printf("got text %s\n", req.Label)
	return &pbMsgV1.MessageResponse{Label: fmt.Sprintf("%d",req.Label), Created: 1000}, nil
}

func (ms *MessageServiceV1) CheckMessage(ctx context.Context, req *pbEmpty.Empty) (res *pbMsgV1.CheckResponse, err error) {
	log.Printf("check\n")
	return &pbMsgV1.CheckResponse{Status: pbMsgV1.ServingStatus_SERVING}, nil
}

func (ms *MessageServiceV1) ListServingStatus(ctx context.Context, req *pbEmpty.Empty) (res *pbMsgV1.MapServingStatusResponse, err error) {
	return &pbMsgV1.MapServingStatusResponse{
		Status: pbMsg.ServingStatus_name,
	}, nil
}


func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pbMsg.RegisterMessageServiceServer(s, &MessageService{})
	pbMsgV1.RegisterMessageServiceServer(s, &MessageServiceV1{})

	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

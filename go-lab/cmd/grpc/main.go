package main

import (
	"log"
	"net"

	googlePb "github.com/golang/protobuf/ptypes/empty"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":8081"
)

type MessageService struct {
	MessageServiceServer
}

func (ms *MessageService) UpdateMessage(ctx context.Context, req *MessageRequest) (res *MessageResponse, err error) {
	log.Printf("got text %s\n", req.Label)
	return &MessageResponse{Label: req.Label, Created: 1000}, nil
}

func (ms *MessageService) CheckMessage(ctx context.Context, req *googlePb.Empty) (res *CheckResponse, err error) {
	log.Printf("check\n")
	return &CheckResponse{Status: ServingStatus_SERVING}, nil
}

func (ms *MessageService) ListServingStatus(ctx context.Context, req *googlePb.Empty) (res *MapServingStatusResponse, err error) {
	return &MapServingStatusResponse{
		Status: ServingStatus_name,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	RegisterMessageServiceServer(s, &MessageService{})

	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

package main

import (
	"fmt"
	"grpc/proto"
	"context"
	"net"

	"google.golang.org/grpc"
)

type Message struct {}

func main() {
	lis, err := net.Listen("tcp", ":4040")
	if err != nil {
		fmt.Println(err)
	}
	grpcServer := grpc.NewServer()
	proto.RegisterMessageServiceServer(grpcServer, &Message{})
	if e := grpcServer.Serve(lis); e != nil {
		fmt.Println(e)
	}
}

func (m *Message) GetMessage(ctx context.Context, message *proto.ClientMessage) (*proto.ClientMessage, error) {
	fmt.Println("message from client", message.Message )
	responseMessage := &proto.ClientMessage{
		Message: "Hello from server",
	}
	return responseMessage, nil
}

func (m *Message) GetMessageStream(r *proto.ClientMessage, stream proto.MessageService_GetMessageStreamServer)  error {
	for i := 0; i< 5; i++ {
		resp := proto.ClientMessage{ Message: "Hi from server" }
		if err := stream.Send(&resp); err != nil {
			return err;
		}
	}
	return nil
}
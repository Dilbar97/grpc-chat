package grpc_chat

import (
	"fmt"
	"net"

	"grpc-chat/chatsvc"

	"google.golang.org/grpc"
)

func main() {
	listen, err := net.Listen("tcp", ":8088")
	if err != nil {
		fmt.Println(fmt.Sprintf("failed to listen: %v", err.Error()))
	}

	chatSvr := chatsvc.ChatServer{}
	grpcSrv := grpc.NewServer()

	chatsvc.RegisterChatServiceServer(grpcSrv, &chatSvr)

	if err = grpcSrv.Serve(listen); err != nil {
		fmt.Println(fmt.Sprintf("failed to serve grpc: %v", err.Error()))
	}
}

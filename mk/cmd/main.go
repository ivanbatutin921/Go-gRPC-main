package main

import (
	"fmt"
	"net"

	"google.golang.org/grpc"
	"root/mk/internal/server"
	pb "root/mk/chat"
	db "root/mk/internal/database"
)



func main() {
	db.ConnectToDB()

	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &server.Server{})
	if err := s.Serve(lis); err != nil {
		fmt.Printf("failed to serve: %v", err)
	}
}

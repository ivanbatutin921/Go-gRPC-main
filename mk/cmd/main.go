package main

import (
	"fmt"
	"net"

	pb "root/mk/chat"
	db "root/mk/internal/database"
	"root/mk/internal/server"

	"google.golang.org/grpc"
)

func main() {
	db.ConnectToDB()
	//db.Migration()

	lis, err := net.Listen("tcp", ":50050")
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
	} else {
		fmt.Println("server started")
	}
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &server.Server{})
	if err := s.Serve(lis); err != nil {
		fmt.Printf("failed to serve: %v", err)
	}
}

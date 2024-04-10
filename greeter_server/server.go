package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "root/chat"

	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
	//"google.golang.org/grpc/credentials/insecure"
)

var (
	port = flag.Int("port", 50052, "The server port")
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}
func runGRPCServer() {
	flag.Parse()
	//Создается слушатель TCP на указанном порту
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	//создаем gRPC сервер
	s := grpc.NewServer()

	//регистрируем сервис
	pb.RegisterGreeterServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func main() {

	go runGRPCServer()

	app := fiber.New()
	app.Get("/hello", func(c *fiber.Ctx) error {
		clientData := new(pb.HelloRequest)
		if err := c.BodyParser(clientData); err != nil {
			return err
		}

		conn, err := grpc.Dial(fmt.Sprintf(":%d", *port), grpc.WithInsecure())
		if err != nil {
			log.Fatalf("Failed to dial server: %v", err)
		}
		defer conn.Close()

		client := pb.NewGreeterClient(conn)
		reply, err := client.SayHello(context.Background(), clientData)
		if err != nil {
			return err
		}

		return c.SendString(reply.Message)
	})
	log.Fatal(app.Listen(":8000"))
}

package main

import (
	"log"

	routes "root/greeter_server/routes"
	pb "root/mk/chat"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Server struct {
	app  *fiber.App
	mk   pb.UserServiceClient
	port string
}

func (s *Server) runGrpcServer() {
	conn, err := grpc.Dial("localhost:50050", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("не получилось соединиться: %v", err)
	}
	s.mk = pb.NewUserServiceClient(conn)
}

func (s *Server) allRoutes() {
	userHandler := routes.ServiceHandler(s.mk)

	s.app.Post("/createuser", userHandler.CreateUser)
	s.app.Post("/createmanyusers", userHandler.CreateManyUsers)
	s.app.Get("/getuser/:id", userHandler.ReadUser)
	s.app.Get("/getallusers", userHandler.ReadAllUsers)
	s.app.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
}
func NewServer(port string) *Server {
	s := &Server{
		app:  fiber.New(),
		port: port,
	}
	s.app.Use(logger.New())
	return s
}

func (s *Server) Run() {
	s.runGrpcServer()
	s.allRoutes()
	log.Fatal(s.app.Listen(":" + s.port))
}

func main() {
	s := NewServer("3000")
	s.Run()
}



package main

import (
	"log"

	routes "root/greeter_server/routes"
	pb "root/mk/chat"

	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	mk pb.UserServiceClient
)

func AllRoutes(app *fiber.App, mk pb.UserServiceClient) {
	app.Post("/createuser", routes.CreateUser(mk))
	app.Get("/getuser/:id", routes.ReadUser(mk))
	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
}

func RunGrpcServer() {
	conn, err := grpc.NewClient("localhost:50050", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	mk = pb.NewUserServiceClient(conn)
}

func main() {
	app := fiber.New()

	go RunGrpcServer()

	AllRoutes(app, mk)

	log.Fatal(app.Listen(":3000"))
}

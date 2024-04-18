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

var (
	mk pb.UserServiceClient
)

func RunGrpcServer() {
	conn, err := grpc.NewClient("localhost:50050", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("не получилось соединиться: %v", err)
	}
	mk = pb.NewUserServiceClient(conn)
}

func AllRoutes(app *fiber.App, mk pb.UserServiceClient) {

	app.Post("/createuser", routes.CreateUser(mk))
	app.Post("/createmanyusers", routes.CreateManyUsers(mk))
	app.Get("/getuser/:id", routes.ReadUser(mk))
	app.Get("/getallusers", routes.ReadAllUsers(mk))
	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
}

func main() {
	app := fiber.New()
	app.Use(logger.New())
	RunGrpcServer()

	AllRoutes(app, mk)

	log.Fatal(app.Listen(":3000"))
}

package main

import (
	//"context"
	"flag"
	"fmt"
	"log"
	//"net"

	pb "root/mk/chat"
	routes "root/greeter_server/routes"

	"github.com/gofiber/fiber/v2"
	// "google.golang.org/grpc"
	// "google.golang.org/grpc/credentials/insecure"
	//"google.golang.org/grpc/credentials/insecure"
)

var (
	mk   pb.UnimplementedUserServiceServer
	port = flag.Int("port", 50052, "The server port")
)

type server struct {
	pb.UnimplementedUserServiceServer
}

func AllRoutes(app *fiber.App, s *server) {
	app.Get("/",routes.ReadUser(mk))
	//app.Post("/", s.CreateUser())
}

func main() {
	app := fiber.New()
	s := new(server)

	log.Fatal(app.Listen(fmt.Sprintf(":%d", *port)))
}

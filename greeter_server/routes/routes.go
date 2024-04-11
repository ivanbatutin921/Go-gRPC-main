package routes

import (
	//	"context"
	"context"
	"log"
	pb "root/mk/chat"

	"github.com/gofiber/fiber/v2"
)

func CreateUser(pb pb.UserServiceClient, req *pb.User) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		body := &pb.User{}
		if err := c.BodyParser(body); err != nil {
			log.Fatal("данные из тела не прочитаны", err)
		}
		_, err := pb.CreateUser(context.Background(), req)
		if err != nil {
			return c.Status(500).JSON(err)
		}
		return nil
	}
}

func ReadUser(pb pb.UserServiceClient) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		body := &pb.UserId{}
		if err := c.BodyParser(body); err != nil {
			log.Fatal("данные из тела не прочитаны", err)
		}
		_, err := pb.ReadUser(context.Background(), body)
		if err != nil {
			return c.Status(500).JSON(err)
		}
		return nil
	}
}

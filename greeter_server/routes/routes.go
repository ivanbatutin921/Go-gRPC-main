package routes

import (
	//	"context"

	//"fmt"

	"context"
	"log"
	pb "root/mk/chat"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func CreateUser(mpx pb.UserServiceClient) fiber.Handler {
	return func(c *fiber.Ctx) error {
		body := &pb.User{}
		if err := c.BodyParser(body); err != nil {
			log.Fatal("данные из тела не прочитаны", err)
			return err
		}
		if mpx == nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		user, err := mpx.CreateUser(context.Background(), body)
		if err != nil {
			log.Println(err)
		}		

		return c.JSON(user)
	}
}

func ReadUser(mpx pb.UserServiceClient) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		i, err := strconv.Atoi(id)
		if err != nil {
			return c.SendString("Error")
		}

		if mpx == nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()

		user, err := mpx.ReadUser(ctx, &pb.UserId{Id: int32(i)})
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		return c.JSON(user)
	}
}

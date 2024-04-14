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

func CreateUser(mk pb.UserServiceClient) fiber.Handler {
	return func(c *fiber.Ctx) error {
		data := pb.User{}
		if err := c.BodyParser(&data); err != nil {
			return c.JSON(fiber.Map{"status": "error", "message": "не удалось прочитать тело запроса", "data": err})
		}
		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()

		ch := make(chan pb.User, 1)

		go func() {
			res, err := mk.CreateUser(ctx, &data)
			if err != nil {
				log.Fatal(err)
			}
			ch <- *res

		}()
		data = <-ch
		return c.JSON(data)
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

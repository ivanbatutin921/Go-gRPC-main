package routes

import (
	"context"
	"log"
	"strconv"
	"time"

	pb "root/mk/chat"

	"github.com/gofiber/fiber/v2"
)

type UserServiceHandler struct {
	pb.UnimplementedUserServiceServer
	mk pb.UserServiceClient
}

func ServiceHandler(mk pb.UserServiceClient) *UserServiceHandler {
	return &UserServiceHandler{mk: mk}
}

func (h *UserServiceHandler) CreateUser(c *fiber.Ctx) error {
	data := pb.User{}
	if err := c.BodyParser(&data); err != nil {
		return c.JSON(fiber.Map{"status": "error", "message": "не удалось прочитать тело запроса", "data": err})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	ch := make(chan pb.User, 1)

	go func() {
		res, err := h.mk.CreateUser(ctx, &data)
		if err != nil {
			log.Fatal(err)
		}
		ch <- *res

	}()
	data = <-ch

	return c.JSON(data)
}

func (h *UserServiceHandler) CreateManyUsers(c *fiber.Ctx) error {
	data := pb.UserList{Users: []*pb.User{}}

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Не удалось распарсить данные из тела запроса", "error": err.Error()})
	}

	ch := make(chan pb.UserList, 1)
	go func() {
		res, err := h.mk.CreateManyUsers(context.Background(), &data)
		if err != nil {
			log.Fatal(err)
		}
		ch <- *res

	}()
	data = <-ch

	return c.JSON(data)
}

func (h *UserServiceHandler) ReadUser(c *fiber.Ctx) error {
	id := c.Params("id")
	i, err := strconv.Atoi(id)
	if err != nil {
		return c.SendString("Error")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	user, err := h.mk.ReadUser(ctx, &pb.UserId{Id: int32(i)})
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.JSON(user)
}

func (h *UserServiceHandler) ReadAllUsers(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	data, err := h.mk.ReadAllUsers(ctx, &pb.Empty{})
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.JSON(data)
}

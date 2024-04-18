package server

import (
	"context"
	"log"

	pb "root/mk/chat"
	db "root/mk/internal/database"
	"root/mk/internal/model"
)

type Server struct {
	pb.UnimplementedUserServiceServer
}

func (s *Server) CreateUser(ctx context.Context, req *pb.User) (*pb.User, error) {
	err := db.DB.DB.Create(&model.User{Name: req.Name, Email: req.Email})
	if err.Error != nil {
		log.Fatal(err.Error)
	}
	return &pb.User{
			Name:  req.Name,
			Email: req.Email},
		nil
}

func (s *Server) CreateManyUsers(ctx context.Context, req *pb.UserList) (*pb.UserList, error) {
	// users := make([]*pb.User, len(req.Users))
	// for i, user := range req.Users {
	// 	users[i] = &pb.User{
	// 		Name:  user.Name,
	// 		Email: user.Email,
	// 	}
	// }

	users := &pb.UserList{
		Users: []*pb.User{
			{Name: "user1", Email: "user1@example.com"},
			{Name: "user2", Email: "user2@example.com"},
			{Name: "user3", Email: "user3@example.com"},
		},
	}

	for _, user := range users.Users {
		err := db.DB.DB.Create(&model.User{Name: user.Name, Email: user.Email})
		if err != nil {
			println("error")
		}
	}
	return &pb.UserList{Users: users.Users}, nil
}

func (s *Server) ReadUser(ctx context.Context, req *pb.UserId) (*pb.User, error) {
	user := model.User{}
	err := db.DB.DB.First(&user, req.Id).Error
	if err != nil {
		return nil, err
	}
	return &pb.User{Name: user.Name, Email: user.Email}, nil
}

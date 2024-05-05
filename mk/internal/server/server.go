package server

import (
	"context"
	"log"

	pb "root/mk/chat"
	"root/mk/internal/model"

	"gorm.io/gorm"
)

type Server struct {
	pb.UnimplementedUserServiceServer
	db *gorm.DB
}

func (s *Server) CreateUser(ctx context.Context, req *pb.User) (*pb.User, error) {
	err := s.db.Create(&model.User{Name: req.Name, Email: req.Email}).Error
	if err != nil {
		log.Fatal(err.Error)
	}
	return &pb.User{
			Name:  req.Name,
			Email: req.Email},
		nil
}

func (s *Server) CreateManyUsers(ctx context.Context, req *pb.UserList) (*pb.UserList, error) {
	for _, user := range req.Users {
		err := s.db.Create(&model.User{Name: user.Name, Email: user.Email})
		if err != nil {
			println("error")
		}
	}
	return &pb.UserList{Users: req.Users}, nil
}

func (s *Server) ReadUser(ctx context.Context, req *pb.UserId) (*pb.User, error) {
	user := model.User{}
	err := s.db.First(&user, req.Id).Error
	if err != nil {
		return nil, err
	}
	return &pb.User{Name: user.Name, Email: user.Email}, nil
}

func (s *Server) ReadAllUsers(ctx context.Context, _ *pb.Empty) (*pb.UserList, error) {
	data := pb.UserList{Users: []*pb.User{}}

	var users []model.User
	s.db.Find(&users)
	for _, user := range users {
		data.Users = append(data.Users, &pb.User{Name: user.Name, Email: user.Email})
	}

	return &data, nil
}

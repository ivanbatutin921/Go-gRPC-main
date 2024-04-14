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
		Name: req.Name, 
		Email: req.Email}, 
	nil
}

func (s *Server) ReadUser(ctx context.Context, req *pb.UserId) (*pb.User, error) {
	user := model.User{}
	err := db.DB.DB.First(&user, req.Id).Error
	if err != nil {
		return nil, err
	}
	return &pb.User{Name: user.Name, Email: user.Email}, nil
}

func(s *Server) UpdateUser(ctx context.Context, req *pb.UpdateUserMessage) (*pb.User, error) {
	
}
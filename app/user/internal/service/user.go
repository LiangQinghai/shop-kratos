package service

import (
	"context"
	"encoding/json"

	pb "gitee.com/qinghailiang/shop-kratos/api/user/v1"
	"gitee.com/qinghailiang/shop-kratos/app/user/internal/biz"
)

type UserService struct {
	pb.UnimplementedUserServer
	uc *biz.UserUseCase
}

func NewUserService(uc *biz.UserUseCase) *UserService {
	return &UserService{uc: uc}
}

func (s *UserService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserReply, error) {
	return &pb.CreateUserReply{}, nil
}
func (s *UserService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserReply, error) {
	return &pb.UpdateUserReply{}, nil
}
func (s *UserService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserReply, error) {
	return &pb.DeleteUserReply{}, nil
}
func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserReply, error) {
	return &pb.GetUserReply{}, nil
}
func (s *UserService) ListUser(ctx context.Context, req *pb.ListUserRequest) (*pb.ListUserReply, error) {
	return &pb.ListUserReply{}, nil
}
func (s *UserService) RegisterUser(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterReply, error) {
	json_b, _ := json.Marshal(req)
	println(string(json_b))
	s.uc.Register(ctx, &biz.User{Name: req.Username, Password: req.Password})
	return &pb.RegisterReply{}, nil
}

package grpc

import (
	"context"
	"fmt"
	"go-session16-crud-user-compose/entity"
	"go-session16-crud-user-compose/service"
	"log"

	pb "go-session16-crud-user-compose/proto"

	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type UserHandler struct {
	pb.UnimplementedUserServiceServer
	userService service.IUserService
}

func NewUserHandler(userService service.IUserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) GetUsers(ctx context.Context, _ *emptypb.Empty) (*pb.GetUserResponse, error) {
	users, err := h.userService.GetAllUsers(ctx)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var userProto []*pb.User
	for _, user := range users {
		userProto = append(userProto, &pb.User{
			Id:        int32(user.ID),
			Name:      user.Name,
			Email:     user.Email,
			Password:  user.Password,
			CreatedAt: timestamppb.New(user.CreatedAt),
			UpdatedAt: timestamppb.New(user.UpdatedAt),
		})

	}
	return &pb.GetUserResponse{Users: userProto}, nil

}

func (h *UserHandler) GetUserByID(ctx context.Context, req *pb.GetUserByIDRequest) (*pb.GetUserByIDResponse, error) {
	user, err := h.userService.GetUserByID(ctx, int(req.Id))
	if err != nil {
		log.Println(err)
		return nil, err
	}
	res := &pb.GetUserByIDResponse{
		User: &pb.User{
			Id:        int32(user.ID),
			Name:      user.Name,
			Email:     user.Email,
			Password:  user.Password,
			CreatedAt: timestamppb.New(user.CreatedAt),
			UpdatedAt: timestamppb.New(user.UpdatedAt),
		},
	}
	return res, nil
}

func (h *UserHandler) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.MutationResponse, error) {
	updatedUser, err := h.userService.UpdateUserByID(ctx, int(req.Id), entity.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	})

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &pb.MutationResponse{
		Message: fmt.Sprintf("Success updated user with id %d", updatedUser.ID),
	}, nil
}

// DeleteUser(context.Context, *DeleteUserRequest) (*MutationResponse, error)
func (h *UserHandler) DeleteUserByID(ctx context.Context, req *pb.DeleteUserRequest) (*pb.MutationResponse, error) {
	err := h.userService.DeleteUserByID(ctx, int(req.Id))
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &pb.MutationResponse{Message: fmt.Sprintf("User Deleted with user id , %d", req.Id)}, nil
}

func (h *UserHandler) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.MutationResponse, error) {
	createduser, err := h.userService.CreateUser(ctx, &entity.User{Name: req.Name, Email: req.Email})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &pb.MutationResponse{Message: fmt.Sprintf("User Created with user id , %d", createduser.ID)}, nil

}

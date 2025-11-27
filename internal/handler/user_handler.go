package handler

import (
	"context"

	"saitama/internal/service"

	"github.com/yudhiana/one-for-all/api/gen/user"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type UserHandler struct {
	svc service.UserServiceRepository
	user.UnimplementedUserServiceServer
}

func NewUserHandler(svc service.UserServiceRepository) *UserHandler {
	return &UserHandler{
		svc: svc,
	}
}

func (h *UserHandler) CreateUser(ctx context.Context, req *user.CreateUserRequest) (*user.CreateUserResponse, error) {
	id, err := h.svc.CreateUser(ctx, req.Name, req.Email)
	if err != nil {
		return nil, err
	}

	return &user.CreateUserResponse{
		Id: id,
	}, nil
}

func (h *UserHandler) GetUser(ctx context.Context, req *user.GetUserRequest) (*user.GetUserResponse, error) {
	userData, err := h.svc.GetUser(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &user.GetUserResponse{
		Id:        userData.ID,
		Name:      userData.Name,
		Email:     userData.Email,
		CreatedAt: timestamppb.New(userData.CreatedAt),
		UpdatedAt: timestamppb.New(userData.UpdatedAt),
	}, nil
}

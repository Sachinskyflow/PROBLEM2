package service

import (
	"context"
	proto "problem2/proto"
)

type Service interface {
	Insert(ctx context.Context, req *proto.InsertUserRequest) (*proto.InsertUserResponse, error)
	Get(ctx context.Context, req *proto.GetUserRequest) (*proto.GetUserResponse, error)
}

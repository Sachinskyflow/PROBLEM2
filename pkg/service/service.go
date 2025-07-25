package service

import (
	"context"
	"problem2/pkg/db"
	proto "problem2/proto"

	aerospike "github.com/aerospike/aerospike-client-go/v8"
)

type RecordService struct {
	aero *aerospike.Client
	proto.UnimplementedAppServiceServer
}

func NewRecordService(ctx context.Context, aero *aerospike.Client) *RecordService {
	return &RecordService{
		aero: aero,
	}
}

func (s *RecordService) Insert(ctx context.Context, req *proto.InsertUserRequest) (*proto.InsertUserResponse, error) {
	datastore := db.NewDatastore(ctx, s.aero)
	rec, err := datastore.Insert(ctx, req)
	if err != nil {
		return nil, err
	}
	return rec, nil
}

func (s *RecordService) Get(ctx context.Context, req *proto.GetUserRequest) (*proto.GetUserResponse, error) {
	datastore := db.NewDatastore(ctx, s.aero)
	rec, err := datastore.Get(ctx, req)
	if err != nil {
		return nil, err
	}
	return rec, nil
}

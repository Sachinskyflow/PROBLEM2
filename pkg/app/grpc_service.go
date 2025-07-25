package app

import (
	"context"
	"net"
	proto "problem2/proto"

	service "problem2/pkg/service"

	aerospike "github.com/aerospike/aerospike-client-go/v8"
	grpc "google.golang.org/grpc"
)

type grpcService struct {
	ctx        context.Context
	aero       *aerospike.Client
	grpcServer *grpc.Server
}

func NewGrpcService(ctx context.Context, aero *aerospike.Client) *grpcService {
	recordService := service.NewRecordService(ctx, aero)
	grpcServer := grpc.NewServer()
	proto.RegisterAppServiceServer(grpcServer, recordService)
	return &grpcService{
		ctx:        ctx,
		aero:       aero,
		grpcServer: grpcServer,
	}
}

func (s *grpcService) Start(ctx context.Context) error {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		return err
	}
	err = s.grpcServer.Serve(lis)
	if err != nil {
		return err
	}
	return nil
}
func (s *grpcService) Stop(ctx context.Context) error {
	s.grpcServer.GracefulStop()
	return nil
}

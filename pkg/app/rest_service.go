package app

import (
	"context"
	"fmt"
	"net/http"
	proto "problem2/proto"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type restService struct {
	ctx        context.Context
	httpServer *http.Server
}

func NewRestService(ctx context.Context) (*restService, error) {
	mux := runtime.NewServeMux()

	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := proto.RegisterAppServiceHandlerFromEndpoint(ctx, mux, "localhost:8080", opts)
	if err != nil {
		return nil, fmt.Errorf("Failed to register rest service: %w", err)
	}
	return &restService{
		ctx: ctx,
		httpServer: &http.Server{
			Addr:    ":8081",
			Handler: mux,
		},
	}, nil
}

func (s *restService) Start(ctx context.Context) error {
	return s.httpServer.ListenAndServe()
}

func (s *restService) Stop(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

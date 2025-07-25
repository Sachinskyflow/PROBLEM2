package app

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	aerospike "github.com/aerospike/aerospike-client-go/v8"
)

var grpcServiceEntity *grpcService
var restServiceEntity *restService

func Main(ctx context.Context, aero *aerospike.Client) error {
	err := start(ctx, aero)
	if err != nil {
		fmt.Printf("Failed to start app: %v\n", err)
		return err
	}
	timectx, cancel := context.WithTimeout(ctx, 30*time.Minute)
	ctx, stopChan := signal.NotifyContext(timectx, os.Interrupt, syscall.SIGTERM)
	defer cancel()
	defer stopChan()
	<-ctx.Done()
	err = stop(ctx)
	if err != nil {
		fmt.Printf("Failed to stop app: %v\n", err)
		return err
	}

	return nil
}

func start(ctx context.Context, aero *aerospike.Client) error {
	grpcServiceEntity = NewGrpcService(ctx, aero)
	go func() {
		err := grpcServiceEntity.Start(ctx)
		if err != nil {
			fmt.Printf("Failed to start grpc service: %v\n", err)
		}
	}()
	var err error
	restServiceEntity, err = NewRestService(ctx)
	if err != nil {
		return err
	}
	go func() {
		err := restServiceEntity.Start(ctx)
		if err != nil {
			fmt.Printf("Failed to start rest service: %v\n", err)
		}
	}()
	return nil
}
func stop(ctx context.Context) error {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		if grpcServiceEntity != nil {
			err := grpcServiceEntity.Stop(ctx)
			if err != nil {
				fmt.Printf("Failed to stop grpc service: %v", err)
			}
		}
	}()
	go func() {
		defer wg.Done()
		if restServiceEntity != nil {
			err := restServiceEntity.Stop(ctx)
			if err != nil {
				fmt.Printf("Failed to stop rest service: %v", err)
			}
		}
	}()
	wg.Wait()
	return nil
}

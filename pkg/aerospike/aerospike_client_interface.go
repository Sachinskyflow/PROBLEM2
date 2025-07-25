package aerospike

import (
	"context"
	proto "problem2/proto"

	"github.com/aerospike/aerospike-client-go/v8"
)

type AerospikeClient interface {
	AeroGet(ctx context.Context, key *aerospike.Key, roll_no int64) (*aerospike.Record, error)
	AeroKey(ctx context.Context, roll_no int64) (*aerospike.Key, error)
	AeroBin(ctx context.Context, req *proto.InsertUserRequest) ([]*aerospike.Bin, error)
	AeroPut(ctx context.Context, key *aerospike.Key, bins []*aerospike.Bin) error
}

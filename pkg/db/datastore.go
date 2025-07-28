package db

import (
	"context"
	proto "problem2/proto"

	aerospike "github.com/aerospike/aerospike-client-go/v8"
)

type Datastore struct {
	aero *aerospike.Client
}

func NewDatastore(ctx context.Context, aero *aerospike.Client) *Datastore {
	return &Datastore{
		aero: aero,
	}
}

const (
	namespace = "test"
	set       = "record"
	name      = "name"
	physics   = "physics"
	chemistry = "chemistry"
	biology   = "biology"
	maths     = "maths"
	english   = "english"
)

func (d *Datastore) Insert(ctx context.Context, req *proto.InsertUserRequest) (*proto.InsertUserResponse, error) {
	key, err := d.AeroKey(ctx, req.RollNo)
	if err != nil {
		return nil, err
	}
	bins, err := d.AeroBin(ctx, req)
	if err != nil {
		return nil, err
	}
	err = d.AeroPut(ctx, key, bins)
	if err != nil {
		return nil, err
	}
	return &proto.InsertUserResponse{
		Message: "Record inserted successfully",
	}, nil
}

func (d *Datastore) Get(ctx context.Context, req *proto.GetUserRequest) (*proto.GetUserResponse, error) {
	key, err := d.AeroKey(ctx, req.RollNo)
	if err != nil {
		return nil, err
	}
	rec, err := d.AeroGet(ctx, key, req.RollNo)
	if err != nil {
		return nil, err
	}
	return rec, nil
}

func (d *Datastore) AeroKey(ctx context.Context, roll_no int64) (*aerospike.Key, error) {
	key, err := aerospike.NewKey(namespace, set, roll_no)
	if err != nil {
		return nil, err
	}
	return key, nil
}

func (d *Datastore) AeroBin(ctx context.Context, req *proto.InsertUserRequest) ([]*aerospike.Bin, error) {
	bins := make([]*aerospike.Bin, 0)
	bins = append(bins, aerospike.NewBin(name, req.Name))
	bins = append(bins, aerospike.NewBin(physics, req.Physics))
	bins = append(bins, aerospike.NewBin(chemistry, req.Chemistry))
	bins = append(bins, aerospike.NewBin(biology, req.Biology))
	bins = append(bins, aerospike.NewBin(maths, req.Maths))
	bins = append(bins, aerospike.NewBin(english, req.English))

	return bins, nil
}

func (d *Datastore) AeroPut(ctx context.Context, key *aerospike.Key, bins []*aerospike.Bin) error {
	err := d.aero.PutBins(nil, key, bins...)
	if err != nil {
		return err
	}
	return nil
}

func (d *Datastore) AeroGet(ctx context.Context, key *aerospike.Key, roll_no int64) (*proto.GetUserResponse, error) {
	rec, err := d.aero.Get(nil, key, name, physics, chemistry, biology, maths, english)
	if err != nil {
		return nil, err
	}

	return &proto.GetUserResponse{
		RollNo:    roll_no,
		Name:      rec.Bins[name].(string),
		Physics:   int64(rec.Bins[physics].(int)),
		Chemistry: int64(rec.Bins[chemistry].(int)),
		Biology:   int64(rec.Bins[biology].(int)),
		Maths:     int64(rec.Bins[maths].(int)),
		English:   int64(rec.Bins[english].(int)),
	}, nil
}

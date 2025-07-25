package test

import (
	"context"
	"testing"

	proto "problem2/proto"

	aerospike "github.com/aerospike/aerospike-client-go/v8"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "TestSuite")
}

type StudentTestSuite struct {
	ctx        context.Context
	flowClient proto.AppServiceClient
	aeroClient *aerospike.Client
}

var studentSuite *StudentTestSuite

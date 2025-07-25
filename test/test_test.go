package test

import (
	"context"
	"time"

	proto "problem2/proto"

	aerospike "github.com/aerospike/aerospike-client-go/v8"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"google.golang.org/grpc"
)

var _ = BeforeSuite(func() {
	studentSuite = &StudentTestSuite{
		ctx:        context.Background(),
		flowClient: nil,
		aeroClient: nil,
	}
	appConnection, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	Expect(err).To(BeNil())
	studentSuite.flowClient = proto.NewAppServiceClient(appConnection)
	studentSuite.aeroClient, err = aerospike.NewClient("localhost", 3000)
	Expect(err).To(BeNil())
	time.Sleep(1 * time.Second)
})

var _ = Describe("Student", func() {
	It("should insert and get a student", func() {
		By("inserting a student")
		insertReq := &proto.InsertUserRequest{
			RollNo:    1,
			Name:      "Sachin",
			Physics:   90,
			Chemistry: 85,
			Biology:   95,
			Maths:     88,
			English:   92,
		}
		insertResp, insertErr := studentSuite.flowClient.Insert(studentSuite.ctx, insertReq)
		Expect(insertErr).To(BeNil())
		Expect(insertResp).NotTo(BeNil())
		Expect(insertResp.Message).To(ContainSubstring("Record inserted successfully"))

		By("getting a student")
		getReq := &proto.GetUserRequest{
			RollNo: 1,
		}
		getResp, getErr := studentSuite.flowClient.Get(studentSuite.ctx, getReq)
		Expect(getErr).To(BeNil())
		Expect(getResp).NotTo(BeNil())
		Expect(getResp.RollNo).To(Equal(getReq.RollNo))
		Expect(getResp.Name).To(Equal(insertReq.Name))
		Expect(getResp.Physics).To(Equal(insertReq.Physics))
		Expect(getResp.Chemistry).To(Equal(insertReq.Chemistry))
		Expect(getResp.Biology).To(Equal(insertReq.Biology))
		Expect(getResp.Maths).To(Equal(insertReq.Maths))
		Expect(getResp.English).To(Equal(insertReq.English))

		By("getting a student which does not exist")
		getReq1 := &proto.GetUserRequest{
			RollNo: 2,
		}
		getResp1, getErr1 := studentSuite.flowClient.Get(studentSuite.ctx, getReq1)
		Expect(getErr1).NotTo(BeNil())
		Expect(getResp1).To(BeNil())
	})
})

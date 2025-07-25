package service_test

import (
	"context"
	"problem2/pkg/service"
	proto "problem2/proto"

	aerospike "github.com/aerospike/aerospike-client-go/v8"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Service", func() {
	It("should insert and get a user", func() {
		By("inserting a user")
		aeroClient, err := aerospike.NewClient("localhost", 3000)
		Expect(err).To(BeNil())
		service := service.NewRecordService(context.Background(), aeroClient)
		insertReq := &proto.InsertUserRequest{
			RollNo:    1,
			Name:      "Sachin",
			Physics:   90,
			Chemistry: 85,
			Biology:   95,
			Maths:     88,
			English:   92,
		}
		insertResp, insertErr := service.Insert(context.Background(), insertReq)
		Expect(insertErr).To(BeNil())
		Expect(insertResp).NotTo(BeNil())
		Expect(insertResp.Message).To(ContainSubstring("Record inserted successfully"))

		By("getting the user")
		getReq := &proto.GetUserRequest{
			RollNo: 1,
		}
		getResp, getErr := service.Get(context.Background(), getReq)
		Expect(getErr).To(BeNil())
		Expect(getResp).NotTo(BeNil())
		Expect(getResp.Name).To(Equal("Sachin"))
		Expect(getResp.Physics).To(Equal(int64(90)))
		Expect(getResp.Chemistry).To(Equal(int64(85)))
		Expect(getResp.Biology).To(Equal(int64(95)))
		Expect(getResp.Maths).To(Equal(int64(88)))
		Expect(getResp.English).To(Equal(int64(92)))

		By("getting the user which does not exist")
		getReq1 := &proto.GetUserRequest{
			RollNo: 3,
		}
		getResp1, getErr1 := service.Get(context.Background(), getReq1)
		Expect(getErr1).NotTo(BeNil())
		Expect(getResp1).To(BeNil())
	})
})

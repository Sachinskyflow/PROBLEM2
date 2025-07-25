package db_test

import (
	"context"
	"problem2/pkg/db"
	proto "problem2/proto"

	"github.com/aerospike/aerospike-client-go/v8"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Insert and Get Requests", func() {
	It("test insert and get request", func() {
		By("inserting a user")
		var insertReq = &proto.InsertUserRequest{
			RollNo:    1,
			Name:      "Sachin",
			Physics:   90,
			Chemistry: 85,
			Biology:   95,
			Maths:     88,
			English:   92,
		}
		aeroClient, err := aerospike.NewClient("localhost", 3000)
		Expect(err).To(BeNil())
		datastore := db.NewDatastore(context.Background(), aeroClient)
		insertResp, insertErr := datastore.Insert(context.Background(), insertReq)
		key, keyErr := datastore.AeroKey(context.Background(), insertReq.RollNo)
		Expect(keyErr).To(BeNil())
		bins, binsErr := datastore.AeroBin(context.Background(), insertReq)
		Expect(binsErr).To(BeNil())
		putErr := datastore.AeroPut(context.Background(), key, bins)
		Expect(putErr).To(BeNil())
		Expect(insertErr).To(BeNil())
		Expect(insertResp).NotTo(BeNil())
		Expect(insertResp.Message).To(ContainSubstring("Record inserted successfully"))

		By("inserting a user with no name")
		var insertReq1 = &proto.InsertUserRequest{
			RollNo:    2,
			Physics:   90,
			Chemistry: 85,
			Biology:   95,
			Maths:     88,
			English:   92,
		}
		insertResp1, insertErr1 := datastore.Insert(context.Background(), insertReq1)
		Expect(insertErr1).To(BeNil())
		key1, keyErr1 := datastore.AeroKey(context.Background(), insertReq1.RollNo)
		Expect(keyErr1).To(BeNil())
		bins1, binsErr1 := datastore.AeroBin(context.Background(), insertReq1)
		Expect(binsErr1).To(BeNil())
		putErr1 := datastore.AeroPut(context.Background(), key1, bins1)
		Expect(putErr1).To(BeNil())
		Expect(insertResp1).NotTo(BeNil())
		Expect(insertResp1.Message).To(ContainSubstring("Record inserted successfully"))

		By("getting the user")
		var getReq = &proto.GetUserRequest{
			RollNo: 1,
		}
		getResp, getErr := datastore.Get(context.Background(), getReq)
		Expect(getErr).To(BeNil())
		Expect(getResp).NotTo(BeNil())
		Expect(getResp.RollNo).To(Equal(insertReq.RollNo))
		Expect(getResp.Name).To(Equal(insertReq.Name))
		Expect(getResp.Physics).To(Equal(insertReq.Physics))
		Expect(getResp.Chemistry).To(Equal(insertReq.Chemistry))
		Expect(getResp.Biology).To(Equal(insertReq.Biology))
		Expect(getResp.Maths).To(Equal(insertReq.Maths))
		Expect(getResp.English).To(Equal(insertReq.English))

		By("getting the user which does not exist")
		var getReq1 = &proto.GetUserRequest{
			RollNo: 2,
		}
		getResp1, get1Err := datastore.Get(context.Background(), getReq1)
		Expect(get1Err).To(BeNil())
		Expect(getResp1).NotTo(BeNil())
		Expect(getResp1.RollNo).To(Equal(insertReq1.RollNo))
		Expect(getResp1.Name).To(Equal(""))
		Expect(getResp1.Physics).To(Equal(insertReq1.Physics))
		Expect(getResp1.Chemistry).To(Equal(insertReq1.Chemistry))
		Expect(getResp1.Biology).To(Equal(insertReq1.Biology))
		Expect(getResp1.Maths).To(Equal(insertReq1.Maths))
		Expect(getResp1.English).To(Equal(insertReq1.English))

		By("getting the user which does not exist")
		var getReq2 = &proto.GetUserRequest{
			RollNo: 3,
		}
		getResp2, get2Err := datastore.Get(context.Background(), getReq2)
		Expect(get2Err).NotTo(BeNil())
		Expect(getResp2).To(BeNil())
	})
})

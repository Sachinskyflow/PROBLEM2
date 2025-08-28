package main_test

import (
	"context"
	"testing"

	aero "github.com/aerospike/aerospike-client-go/v8"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/skyflowapi/flowdb/pkg/config"
	"github.com/skyflowapi/flowdb/pkg/db/util"
)

func TestUtil(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "DB Util Suite")
}

type UtilSuite struct {
	ctx context.Context
}

var (
	testSuite *UtilSuite
	ctx       context.Context
)

var _ = BeforeSuite(func() {
	ctx = context.Background()
	testSuite = &UtilSuite{
		ctx: ctx,
	}

	// Set up test configuration
	config.GetAppConfig().AeroDatabaseConfig.RackAware = true
	config.GetAppConfig().AeroDatabaseConfig.TraceEnabled = true
})

var _ = Describe("ScanPolicyForSC", func() {
	Context("when creating scan policy for strong consistency", func() {
		When("rack aware is enabled", func() {
			BeforeEach(func() {
				config.GetAppConfig().AeroDatabaseConfig.RackAware = true
			})

			It("should set ReadModeSC to ReadModeSCAllowReplica", func() {
				policy := aero.NewScanPolicy()
				util.ScanPolicyForSC(policy)

				Expect(policy.ReadModeSC).To(Equal(aero.ReadModeSCAllowReplica))
			})

			It("should set ReplicaPolicy to PREFER_RACK", func() {
				policy := aero.NewScanPolicy()
				util.ScanPolicyForSC(policy)

				Expect(policy.ReplicaPolicy).To(Equal(aero.PREFER_RACK))
			})
		})

		When("rack aware is disabled", func() {
			BeforeEach(func() {
				config.GetAppConfig().AeroDatabaseConfig.RackAware = false
			})

			It("should not modify ReadModeSC", func() {
				policy := aero.NewScanPolicy()
				originalReadModeSC := policy.ReadModeSC
				util.ScanPolicyForSC(policy)

				Expect(policy.ReadModeSC).To(Equal(originalReadModeSC))
			})

			It("should not modify ReplicaPolicy", func() {
				policy := aero.NewScanPolicy()
				originalReplicaPolicy := policy.ReplicaPolicy
				util.ScanPolicyForSC(policy)

				Expect(policy.ReplicaPolicy).To(Equal(originalReplicaPolicy))
			})
		})

		It("should not modify other policy settings", func() {
			policy := aero.NewScanPolicy()
			originalMaxRetries := policy.MaxRetries
			originalSocketTimeout := policy.SocketTimeout

			util.ScanPolicyForSC(policy)

			Expect(policy.MaxRetries).To(Equal(originalMaxRetries))
			Expect(policy.SocketTimeout).To(Equal(originalSocketTimeout))
		})
	})
})

var _ = Describe("QueryPolicyForSC", func() {
	Context("when creating query policy for strong consistency", func() {
		When("rack aware is enabled", func() {
			BeforeEach(func() {
				config.GetAppConfig().AeroDatabaseConfig.RackAware = true
			})

			It("should set ReadModeSC to ReadModeSCAllowReplica", func() {
				policy := aero.NewQueryPolicy()
				util.QueryPolicyForSC(policy)

				Expect(policy.ReadModeSC).To(Equal(aero.ReadModeSCAllowReplica))
			})

			It("should set ReplicaPolicy to PREFER_RACK", func() {
				policy := aero.NewQueryPolicy()
				util.QueryPolicyForSC(policy)

				Expect(policy.ReplicaPolicy).To(Equal(aero.PREFER_RACK))
			})
		})

		When("rack aware is disabled", func() {
			BeforeEach(func() {
				config.GetAppConfig().AeroDatabaseConfig.RackAware = false
			})

			It("should not modify ReadModeSC", func() {
				policy := aero.NewQueryPolicy()
				originalReadModeSC := policy.ReadModeSC
				util.QueryPolicyForSC(policy)

				Expect(policy.ReadModeSC).To(Equal(originalReadModeSC))
			})

			It("should not modify ReplicaPolicy", func() {
				policy := aero.NewQueryPolicy()
				originalReplicaPolicy := policy.ReplicaPolicy
				util.QueryPolicyForSC(policy)

				Expect(policy.ReplicaPolicy).To(Equal(originalReplicaPolicy))
			})
		})

		It("should not modify other policy settings", func() {
			policy := aero.NewQueryPolicy()
			originalMaxRetries := policy.MaxRetries
			originalSocketTimeout := policy.SocketTimeout

			util.QueryPolicyForSC(policy)

			Expect(policy.MaxRetries).To(Equal(originalMaxRetries))
			Expect(policy.SocketTimeout).To(Equal(originalSocketTimeout))
		})
	})
})

var _ = Describe("BatchPolicyForSC", func() {
	Context("when creating batch policy for strong consistency", func() {
		When("rack aware is enabled", func() {
			BeforeEach(func() {
				config.GetAppConfig().AeroDatabaseConfig.RackAware = true
			})

			It("should set ReadModeSC to ReadModeSCAllowReplica", func() {
				policy := aero.NewBatchPolicy()
				util.BatchPolicyForSC(policy)

				Expect(policy.ReadModeSC).To(Equal(aero.ReadModeSCAllowReplica))
			})

			It("should set ReplicaPolicy to PREFER_RACK", func() {
				policy := aero.NewBatchPolicy()
				util.BatchPolicyForSC(policy)

				Expect(policy.ReplicaPolicy).To(Equal(aero.PREFER_RACK))
			})
		})

		When("rack aware is disabled", func() {
			BeforeEach(func() {
				config.GetAppConfig().AeroDatabaseConfig.RackAware = false
			})

			It("should not modify ReadModeSC", func() {
				policy := aero.NewBatchPolicy()
				originalReadModeSC := policy.ReadModeSC
				util.BatchPolicyForSC(policy)

				Expect(policy.ReadModeSC).To(Equal(originalReadModeSC))
			})

			It("should not modify ReplicaPolicy", func() {
				policy := aero.NewBatchPolicy()
				originalReplicaPolicy := policy.ReplicaPolicy
				util.BatchPolicyForSC(policy)

				Expect(policy.ReplicaPolicy).To(Equal(originalReplicaPolicy))
			})
		})

		It("should not modify other policy settings", func() {
			policy := aero.NewBatchPolicy()
			originalMaxRetries := policy.MaxRetries
			originalSocketTimeout := policy.SocketTimeout

			util.BatchPolicyForSC(policy)

			Expect(policy.MaxRetries).To(Equal(originalMaxRetries))
			Expect(policy.SocketTimeout).To(Equal(originalSocketTimeout))
		})
	})
})

var _ = Describe("WritePolicyForSC", func() {
	Context("when creating write policy for strong consistency", func() {
		When("rack aware is enabled", func() {
			BeforeEach(func() {
				config.GetAppConfig().AeroDatabaseConfig.RackAware = true
			})

			It("should set ReadModeSC to ReadModeSCAllowReplica", func() {
				policy := aero.NewWritePolicy(0, 0)
				util.WritePolicyForSC(policy)

				Expect(policy.ReadModeSC).To(Equal(aero.ReadModeSCAllowReplica))
			})

			It("should set ReplicaPolicy to PREFER_RACK", func() {
				policy := aero.NewWritePolicy(0, 0)
				util.WritePolicyForSC(policy)

				Expect(policy.ReplicaPolicy).To(Equal(aero.PREFER_RACK))
			})
		})

		When("rack aware is disabled", func() {
			BeforeEach(func() {
				config.GetAppConfig().AeroDatabaseConfig.RackAware = false
			})

			It("should not modify ReadModeSC", func() {
				policy := aero.NewWritePolicy(0, 0)
				originalReadModeSC := policy.ReadModeSC
				util.WritePolicyForSC(policy)

				Expect(policy.ReadModeSC).To(Equal(originalReadModeSC))
			})

			It("should not modify ReplicaPolicy", func() {
				policy := aero.NewWritePolicy(0, 0)
				originalReplicaPolicy := policy.ReplicaPolicy
				util.WritePolicyForSC(policy)

				Expect(policy.ReplicaPolicy).To(Equal(originalReplicaPolicy))
			})
		})

		It("should not modify other policy settings", func() {
			policy := aero.NewWritePolicy(0, 0)
			originalMaxRetries := policy.MaxRetries
			originalSocketTimeout := policy.SocketTimeout

			util.WritePolicyForSC(policy)

			Expect(policy.MaxRetries).To(Equal(originalMaxRetries))
			Expect(policy.SocketTimeout).To(Equal(originalSocketTimeout))
		})
	})
})

var _ = Describe("AddAerospikeTrace", func() {
	Context("when adding aerospike trace to message", func() {
		When("trace is enabled", func() {
			BeforeEach(func() {
				config.GetAppConfig().AeroDatabaseConfig.TraceEnabled = true
			})

			It("should append trace information to message", func() {
				message := "test message"
				trace := "test trace\nline 2"

				result := util.AddAerospikeTrace(message, trace)

				Expect(result).To(ContainSubstring(message))
				Expect(result).To(ContainSubstring("aerospike sdk trace"))
				Expect(result).To(ContainSubstring("test trace"))
			})
		})

		When("trace is disabled", func() {
			BeforeEach(func() {
				config.GetAppConfig().AeroDatabaseConfig.TraceEnabled = false
			})

			It("should return original message", func() {
				message := "test message"
				trace := "test trace"

				result := util.AddAerospikeTrace(message, trace)

				Expect(result).To(Equal(message))
			})
		})

		When("trace contains multiple lines", func() {
			BeforeEach(func() {
				config.GetAppConfig().AeroDatabaseConfig.TraceEnabled = true
			})

			It("should handle multiline trace", func() {
				message := "test message"
				trace := "line1\nline2\nline3"

				result := util.AddAerospikeTrace(message, trace)

				Expect(result).To(ContainSubstring(message))
				Expect(result).To(ContainSubstring("aerospike sdk trace"))
			})
		})
	})
})

var _ = Describe("BasePolicyForSC", func() {
	Context("when creating base policy for strong consistency", func() {
		When("rack aware is enabled", func() {
			BeforeEach(func() {
				config.GetAppConfig().AeroDatabaseConfig.RackAware = true
			})

			It("should set ReadModeSC to ReadModeSCAllowReplica", func() {
				policy := &aero.BasePolicy{}
				util.BasePolicyForSC(policy)

				Expect(policy.ReadModeSC).To(Equal(aero.ReadModeSCAllowReplica))
			})

			It("should set ReplicaPolicy to PREFER_RACK", func() {
				policy := &aero.BasePolicy{}
				util.BasePolicyForSC(policy)

				Expect(policy.ReplicaPolicy).To(Equal(aero.PREFER_RACK))
			})
		})

		When("rack aware is disabled", func() {
			BeforeEach(func() {
				config.GetAppConfig().AeroDatabaseConfig.RackAware = false
			})

			It("should not modify ReadModeSC", func() {
				policy := &aero.BasePolicy{}
				originalReadModeSC := policy.ReadModeSC
				util.BasePolicyForSC(policy)

				Expect(policy.ReadModeSC).To(Equal(originalReadModeSC))
			})

			It("should not modify ReplicaPolicy", func() {
				policy := &aero.BasePolicy{}
				originalReplicaPolicy := policy.ReplicaPolicy
				util.BasePolicyForSC(policy)

				Expect(policy.ReplicaPolicy).To(Equal(originalReplicaPolicy))
			})
		})
	})
})

var _ = Describe("BasePolicyForSCWithReplicaPreference", func() {
	Context("when creating base policy with replica preference", func() {
		It("should always set ReadModeSC to ReadModeSCAllowReplica", func() {
			policy := &aero.BasePolicy{}
			util.BasePolicyForSCWithReplicaPreference(policy)

			Expect(policy.ReadModeSC).To(Equal(aero.ReadModeSCAllowReplica))
		})

		When("rack aware is enabled", func() {
			BeforeEach(func() {
				config.GetAppConfig().AeroDatabaseConfig.RackAware = true
			})

			It("should set ReplicaPolicy to PREFER_RACK", func() {
				policy := &aero.BasePolicy{}
				util.BasePolicyForSCWithReplicaPreference(policy)

				Expect(policy.ReplicaPolicy).To(Equal(aero.PREFER_RACK))
			})
		})

		When("rack aware is disabled", func() {
			BeforeEach(func() {
				config.GetAppConfig().AeroDatabaseConfig.RackAware = false
			})

			It("should set ReplicaPolicy to MASTER_PROLES", func() {
				policy := &aero.BasePolicy{}
				util.BasePolicyForSCWithReplicaPreference(policy)

				Expect(policy.ReplicaPolicy).To(Equal(aero.MASTER_PROLES))
			})
		})
	})
})

var _ = Describe("BatchReadPolicyForSCWithReplicaPreference", func() {
	Context("when creating batch read policy with replica preference", func() {
		It("should always set ReadModeSC to ReadModeSCAllowReplica", func() {
			policy := &aero.BatchReadPolicy{}
			util.BatchReadPolicyForSCWithReplicaPreference(policy)

			Expect(policy.ReadModeSC).To(Equal(aero.ReadModeSCAllowReplica))
		})

		It("should only modify ReadModeSC", func() {
			policy := &aero.BatchReadPolicy{}
			originalReadModeSC := policy.ReadModeSC

			util.BatchReadPolicyForSCWithReplicaPreference(policy)

			Expect(policy.ReadModeSC).To(Equal(aero.ReadModeSCAllowReplica))
			Expect(policy.ReadModeSC).NotTo(Equal(originalReadModeSC))
		})
	})
})

var _ = Describe("DeletePolicyForSC", func() {
	Context("when creating delete policy for strong consistency", func() {
		When("enterprise edition is enabled", func() {
			BeforeEach(func() {
				config.GetAppConfig().AeroService.IsEnterpriseEdition = true
			})

			It("should return policy with DurableDelete enabled", func() {
				policy := util.DeletePolicyForSC()

				Expect(policy).NotTo(BeNil())
				Expect(policy.DurableDelete).To(BeTrue())
			})
		})

		When("enterprise edition is disabled", func() {
			BeforeEach(func() {
				config.GetAppConfig().AeroService.IsEnterpriseEdition = false
			})

			It("should return policy with DurableDelete disabled", func() {
				policy := util.DeletePolicyForSC()

				Expect(policy).NotTo(BeNil())
				Expect(policy.DurableDelete).To(BeFalse())
			})
		})
	})
})

var _ = Describe("BatchDeletePolicyForSC", func() {
	Context("when creating batch delete policy for strong consistency", func() {
		When("enterprise edition is enabled", func() {
			BeforeEach(func() {
				config.GetAppConfig().AeroService.IsEnterpriseEdition = true
			})

			It("should return policy with DurableDelete enabled", func() {
				policy := util.BatchDeletePolicyForSC()

				Expect(policy).NotTo(BeNil())
				Expect(policy.DurableDelete).To(BeTrue())
			})
		})

		When("enterprise edition is disabled", func() {
			BeforeEach(func() {
				config.GetAppConfig().AeroService.IsEnterpriseEdition = false
			})

			It("should return policy with DurableDelete disabled", func() {
				policy := util.BatchDeletePolicyForSC()

				Expect(policy).NotTo(BeNil())
				Expect(policy.DurableDelete).To(BeFalse())
			})
		})
	})
})

var _ = Describe("Policy Integration Tests", func() {
	Context("when using policies with different rack aware settings", func() {
		It("should maintain consistency across all policy types when rack aware is true", func() {
			// Set rack aware to true
			config.GetAppConfig().AeroDatabaseConfig.RackAware = true

			// Test that all policy types behave consistently
			scanPolicy := aero.NewScanPolicy()
			queryPolicy := aero.NewQueryPolicy()
			batchPolicy := aero.NewBatchPolicy()
			writePolicy := aero.NewWritePolicy(0, 0)

			// Apply policies
			util.ScanPolicyForSC(scanPolicy)
			util.QueryPolicyForSC(queryPolicy)
			util.BatchPolicyForSC(batchPolicy)
			util.WritePolicyForSC(writePolicy)

			// Verify all policies have the same ReadModeSC and ReplicaPolicy
			expectedReadModeSC := aero.ReadModeSCAllowReplica
			expectedReplicaPolicy := aero.PREFER_RACK

			Expect(scanPolicy.ReadModeSC).To(Equal(expectedReadModeSC))
			Expect(queryPolicy.ReadModeSC).To(Equal(expectedReadModeSC))
			Expect(batchPolicy.ReadModeSC).To(Equal(expectedReadModeSC))
			Expect(writePolicy.ReadModeSC).To(Equal(expectedReadModeSC))

			Expect(scanPolicy.ReplicaPolicy).To(Equal(expectedReplicaPolicy))
			Expect(queryPolicy.ReplicaPolicy).To(Equal(expectedReplicaPolicy))
			Expect(batchPolicy.ReplicaPolicy).To(Equal(expectedReplicaPolicy))
			Expect(writePolicy.ReplicaPolicy).To(Equal(expectedReplicaPolicy))
		})

		It("should handle rack aware false consistently", func() {
			// Set rack aware to false
			config.GetAppConfig().AeroDatabaseConfig.RackAware = false

			// Test that all policy types behave consistently when rack aware is false
			scanPolicy := aero.NewScanPolicy()
			queryPolicy := aero.NewQueryPolicy()
			batchPolicy := aero.NewBatchPolicy()
			writePolicy := aero.NewWritePolicy(0, 0)

			// Apply policies
			util.ScanPolicyForSC(scanPolicy)
			util.QueryPolicyForSC(queryPolicy)
			util.BatchPolicyForSC(batchPolicy)
			util.WritePolicyForSC(writePolicy)

			// Verify all policies maintain their default values
			Expect(scanPolicy.ReadModeSC).To(Equal(aero.ReadModeSC(0)))  // Default value
			Expect(queryPolicy.ReadModeSC).To(Equal(aero.ReadModeSC(0))) // Default value
			Expect(batchPolicy.ReadModeSC).To(Equal(aero.ReadModeSC(0))) // Default value
			Expect(writePolicy.ReadModeSC).To(Equal(aero.ReadModeSC(0))) // Default value

			Expect(scanPolicy.ReplicaPolicy).To(Equal(aero.ReplicaPolicy(3)))  // Default value (SEQUENCE)
			Expect(queryPolicy.ReplicaPolicy).To(Equal(aero.ReplicaPolicy(3))) // Default value (SEQUENCE)
			Expect(batchPolicy.ReplicaPolicy).To(Equal(aero.ReplicaPolicy(3))) // Default value (SEQUENCE)
			Expect(writePolicy.ReplicaPolicy).To(Equal(aero.ReplicaPolicy(3))) // Default value (SEQUENCE)
		})
	})
})

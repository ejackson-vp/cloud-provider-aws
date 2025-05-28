package voltagepark

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws/credentials"
	cloudprovider "k8s.io/cloud-provider"

	"k8s.io/cloud-provider-aws/pkg/providers/v1/config"
	"k8s.io/cloud-provider-aws/pkg/providers/v1/iface"
)

func TestVoltageParkVariant_IsSupportedNode(t *testing.T) {
	variant := &voltageParkVariant{}

	testCases := []struct {
		name     string
		nodeName string
		expected bool
	}{
		{
			name:     "node starting with g",
			nodeName: "g442",
			expected: true,
		},
		{
			name:     "node starting with g lowercase",
			nodeName: "g443",
			expected: true,
		},
		{
			name:     "node starting with g uppercase",
			nodeName: "GPU-server",
			expected: false, // only lowercase g
		},
		{
			name:     "node containing g but not starting",
			nodeName: "worker-gpu-1",
			expected: false,
		},
		{
			name:     "regular AWS node",
			nodeName: "ip-10-0-1-100.us-west-2.compute.internal",
			expected: false,
		},
		{
			name:     "fargate node",
			nodeName: "fargate-ip-10-0-1-100.us-west-2.compute.internal",
			expected: false,
		},
		{
			name:     "node starting with other letter",
			nodeName: "worker-node-1",
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := variant.IsSupportedNode(tc.nodeName)
			if result != tc.expected {
				t.Errorf("IsSupportedNode(%q) = %v, expected %v", tc.nodeName, result, tc.expected)
			}
		})
	}
}

func TestVoltageParkVariant_InstanceExists(t *testing.T) {
	variant := &voltageParkVariant{}

	// InstanceExists should always return true to prevent deletion
	exists, err := variant.InstanceExists("any-instance-id", "any-vpc-id")
	if err != nil {
		t.Errorf("InstanceExists returned unexpected error: %v", err)
	}
	if !exists {
		t.Errorf("InstanceExists should always return true, got false")
	}
}

func TestVoltageParkVariant_InstanceShutdown(t *testing.T) {
	variant := &voltageParkVariant{}

	// InstanceShutdown should always return false to indicate not shutdown
	shutdown, err := variant.InstanceShutdown("any-instance-id", "any-vpc-id")
	if err != nil {
		t.Errorf("InstanceShutdown returned unexpected error: %v", err)
	}
	if shutdown {
		t.Errorf("InstanceShutdown should always return false, got true")
	}
}

func TestVoltageParkVariant_NotImplementedMethods(t *testing.T) {
	variant := &voltageParkVariant{}

	// Test NodeAddresses returns NotImplemented
	_, err := variant.NodeAddresses("instance-id", "vpc-id")
	if err != cloudprovider.NotImplemented {
		t.Errorf("NodeAddresses should return NotImplemented, got: %v", err)
	}

	// Test GetZone returns NotImplemented
	_, err = variant.GetZone("instance-id", "vpc-id", "us-west-2")
	if err != cloudprovider.NotImplemented {
		t.Errorf("GetZone should return NotImplemented, got: %v", err)
	}

	// Test InstanceTypeByProviderID returns NotImplemented
	_, err = variant.InstanceTypeByProviderID("provider-id")
	if err != cloudprovider.NotImplemented {
		t.Errorf("InstanceTypeByProviderID should return NotImplemented, got: %v", err)
	}
}

func TestVoltageParkVariant_Initialize(t *testing.T) {
	variant := &voltageParkVariant{}

	// Mock objects for testing
	cloudConfig := &config.CloudConfig{}
	credentials := &credentials.Credentials{}
	var provider config.SDKProvider
	var ec2API iface.EC2
	region := "us-west-2"

	err := variant.Initialize(cloudConfig, credentials, provider, ec2API, region)
	if err != nil {
		t.Errorf("Initialize should not return error, got: %v", err)
	}

	// No fields to verify since the struct is now empty and doesn't store anything
}

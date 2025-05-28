package voltagepark

import (
	"strings"

	"github.com/aws/aws-sdk-go/aws/credentials"
	v1 "k8s.io/api/core/v1"
	cloudprovider "k8s.io/cloud-provider"

	"k8s.io/cloud-provider-aws/pkg/providers/v1/config"
	"k8s.io/cloud-provider-aws/pkg/providers/v1/iface"
	"k8s.io/cloud-provider-aws/pkg/providers/v1/variant"
)

type voltageParkVariant struct{}

func (v *voltageParkVariant) Initialize(cloudConfig *config.CloudConfig, credentials *credentials.Credentials, provider config.SDKProvider, ec2API iface.EC2, region string) error {
	// No initialization needed for this simple variant
	return nil
}

// IsSupportedNode determines if this variant should handle the given node
// This variant protects nodes whose names start with "g"
func (v *voltageParkVariant) IsSupportedNode(nodeName string) bool {
	// Protect nodes that start with "g"
	return strings.HasPrefix(nodeName, "g")
}

// InstanceExists always returns true to prevent CCM from deleting these nodes
// This is the key method that prevents node deletion
func (v *voltageParkVariant) InstanceExists(instanceID, vpcID string) (bool, error) {
	// Always return true to prevent deletion of nodes starting with "g"
	return true, nil
}

// InstanceShutdown returns false since we don't want these nodes to be considered shutdown
func (v *voltageParkVariant) InstanceShutdown(instanceID, vpcID string) (bool, error) {
	// Return false to indicate the instance is not shutdown
	return false, nil
}

// NodeAddresses returns not implemented - let Kubernetes handle node addresses
func (v *voltageParkVariant) NodeAddresses(instanceID, vpcID string) ([]v1.NodeAddress, error) {
	return nil, cloudprovider.NotImplemented
}

// GetZone returns not implemented - let Kubernetes handle zone information
func (v *voltageParkVariant) GetZone(instanceID, vpcID, region string) (cloudprovider.Zone, error) {
	return cloudprovider.Zone{}, cloudprovider.NotImplemented
}

// InstanceTypeByProviderID returns not implemented - let Kubernetes handle instance type
func (v *voltageParkVariant) InstanceTypeByProviderID(id string) (string, error) {
	return "", cloudprovider.NotImplemented
}

func init() {
	v := &voltageParkVariant{}
	variant.RegisterVariant(
		"voltagepark",
		v,
	)
}

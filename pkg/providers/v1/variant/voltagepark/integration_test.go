package voltagepark

import (
	"testing"

	"k8s.io/cloud-provider-aws/pkg/providers/v1/variant"
)

func TestVoltageParkVariantRegistration(t *testing.T) {
	// Test that our variant is properly registered
	testNodeName := "g444"

	// Check if any variant supports this node
	if !variant.IsVariantNode(testNodeName) {
		t.Errorf("Expected variant to support node %s, but none found", testNodeName)
	}

	// Check if we can get the specific variant
	v := variant.GetVariant(testNodeName)
	if v == nil {
		t.Errorf("Expected to get variant for node %s, but got nil", testNodeName)
	}

	// Check if the node type is correct
	nodeType := variant.NodeType(testNodeName)
	if nodeType != "voltagepark" {
		t.Errorf("Expected node type 'voltagepark', got '%s'", nodeType)
	}

	// Verify it's our specific variant by testing the interface
	if v != nil {
		exists, err := v.InstanceExists("test-instance", "test-vpc")
		if err != nil {
			t.Errorf("InstanceExists returned error: %v", err)
		}
		if !exists {
			t.Errorf("Expected InstanceExists to return true for voltagepark variant")
		}
	}
}

func TestVoltageParkVariantDoesNotSupportOtherNodes(t *testing.T) {
	// Test that nodes not starting with "g" are not supported by our variant
	otherNodeNames := []string{
		"ip-10-0-1-100.us-west-2.compute.internal",
		"worker-node-1",
		"regular-node",
		"fargate-node",
		"cpu-worker",
	}

	for _, nodeName := range otherNodeNames {
		v := variant.GetVariant(nodeName)
		if v != nil {
			nodeType := variant.NodeType(nodeName)
			if nodeType == "voltagepark" {
				t.Errorf("voltagepark variant should not support node %s", nodeName)
			}
		}
	}
}

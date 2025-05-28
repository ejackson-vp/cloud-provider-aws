# Voltagepark Variant

The `voltagepark` variant prevents the AWS Cloud Controller Manager (CCM) from deleting nodes whose names start with "g". This is useful for protecting specific GPU or other specialized nodes that should not be managed by the AWS CCM.

## How It Works

The variant identifies nodes that start with "g" and prevents the CCM from deleting them by:

1. **Node Identification**: The `IsSupportedNode` method identifies nodes that should be protected from deletion based on the naming pattern.
2. **Prevent Deletion**: The `InstanceExists` method always returns `true` for supported nodes, preventing the CCM from deleting them.
3. **Minimal Implementation**: Other methods return `cloudprovider.NotImplemented` to let Kubernetes handle those operations normally.

## Node Naming Pattern

The variant protects nodes with the following naming pattern:

- Nodes starting with `g` (lowercase)

## Customization

You can customize the node identification logic by modifying the `IsSupportedNode` method in `voltagepark.go`. For example:

```go
func (v *voltageParkVariant) IsSupportedNode(nodeName string) bool {
    // Add your custom logic here
    return strings.HasPrefix(nodeName, "gpu-") ||
           strings.HasPrefix(nodeName, "graphics-")
}
```

## Usage

The variant is automatically registered when the AWS cloud provider is loaded. No additional configuration is required.

### Example Node Names

The following node names would be protected from deletion:

- `g442`
- `g0443`

### Example Node Names NOT Protected

The following node names would NOT be protected (handled by normal AWS logic):

- `ip-10-0-1-100.us-west-2.compute.internal` (regular EC2 instance)
- `fargate-ip-10-0-1-100.us-west-2.compute.internal` (Fargate node)
- `worker-node-1` (doesn't start with "g")
- `cpu-worker-1` (doesn't start with "g")
- `GPU-server` (uppercase G, not lowercase)

## Testing

Run the tests to verify the variant works correctly:

```bash
go test ./pkg/providers/v1/variant/voltagepark/... -v
```

## Integration

The variant is automatically imported in `pkg/providers/v1/aws.go` and registers itself during package initialization. When the AWS CCM starts, it will automatically use this variant for nodes that start with "g".

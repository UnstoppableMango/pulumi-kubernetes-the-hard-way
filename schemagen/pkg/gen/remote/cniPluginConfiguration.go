package remote

import (
	"maps"
	"slices"

	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/schemagen/pkg/gen/props"
	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/schemagen/pkg/gen/types"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

func generateCniPluginConfiguration(commandSpec schema.PackageSpec) schema.ResourceSpec {
	inputs := map[string]schema.PropertySpec{
		"connection": {
			Description: "The parameters with which to connect to the remote host.",
			TypeSpec:    types.ExtType(commandSpec, "Connection", "remote"),
		},
		"subnet": props.String("The subnet to use for the CNI bridge plugin configuration."),
	}

	requiredInputs := []string{"connection", "subnet"}

	outputs := map[string]schema.PropertySpec{
		"bridge": {
			Description: "The bridge plugin configuration.",
			TypeSpec:    types.LocalResource("CniBridgePluginConfiguration", "remote"),
		},
		"etcCniMkdir": {
			Description: "The /etc/cni/net.d mkdir operation.",
			TypeSpec:    types.LocalResource("Mkdir", "tools"),
		},
		"file": {
			Description: "The file on the remote system.",
			TypeSpec:    types.LocalResource("File", "remote"),
		},
		"loopback": {
			Description: "The loopback plugin configuration.",
			TypeSpec:    types.LocalResource("CniLoopbackPluginConfiguration", "remote"),
		},
	}
	maps.Copy(outputs, inputs)

	requiredOutputs := slices.Concat(requiredInputs, []string{
		"bridge",
		"etcCniMkdir",
		"file",
		"loopback",
	})

	return schema.ResourceSpec{
		IsComponent: true,
		ObjectTypeSpec: schema.ObjectTypeSpec{
			Description: "The CNI plugin configuration.",
			Properties:  outputs,
			Required:    requiredOutputs,
		},
		InputProperties: inputs,
		RequiredInputs:  requiredInputs,
	}
}

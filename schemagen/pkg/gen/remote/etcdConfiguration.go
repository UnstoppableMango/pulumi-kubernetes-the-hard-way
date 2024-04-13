package remote

import (
	"maps"
	"slices"

	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/schemagen/pkg/gen/types"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

func generateEtcdConfiguration(commandSpec schema.PackageSpec) schema.ResourceSpec {
	inputs := map[string]schema.PropertySpec{
		"caPem": {
			Description: "The PEM encoded certificate authority data.",
			TypeSpec:    types.String,
		},
		"certPem": {
			Description: "The PEM encoded certificate data.",
			TypeSpec:    types.String,
		},
		"configurationDirectory": {
			Description: "The directory to store etcd configuration.",
			TypeSpec:    types.String,
			Default:     "/etc/etcd",
		},
		"connection": {
			Description: "The parameters with which to connect to the remote host.",
			TypeSpec:    types.ExtType(commandSpec, "Connection", "remote"),
		},
		"dataDirectory": {
			Description: "The directory etcd will store its data.",
			TypeSpec:    types.String,
			Default:     "/var/lib/etcd",
		},
		"etcdPath": {
			Description: "The path to the `etcd` binary.",
			TypeSpec:    types.String,
		},
		"internalIp": {
			Description: "The IP used to serve client requests and communicate with etcd peers.",
			TypeSpec:    types.String,
		},
		"keyPem": {
			Description: "The PEM encoded key data.",
			TypeSpec:    types.String,
		},
	}

	requiredInputs := []string{
		"caPem",
		"certPem",
		"connection",
		"etcdPath",
		"internalIp",
		"keyPem",
	}

	outputs := map[string]schema.PropertySpec{
		"caFile": {
			Description: "The remote certificate authority file.",
			TypeSpec:    types.LocalResource("File", "remote"),
		},
		"certFile": {
			Description: "The remote certificate file.",
			TypeSpec:    types.LocalResource("File", "remote"),
		},
		"configurationMkdir": {
			Description: "The configuration mkdir operation.",
			TypeSpec:    types.LocalResource("Mkdir", "tools"),
		},
		"dataMkdir": {
			Description: "The data mkdir operation.",
			TypeSpec:    types.LocalResource("Mkdir", "tools"),
		},
		"keyFile": {
			Description: "The remote key file.",
			TypeSpec:    types.LocalResource("File", "remote"),
		},
		"value": {
			Description: "A bag of properties to be consumed by other resources.",
			TypeSpec:    types.LocalType("EtcdConfigurationProps", "remote"),
		},
	}
	maps.Copy(outputs, inputs)

	requiredOutputs := slices.Concat(
		requiredInputs,
		[]string{
			"caFile",
			"certFile",
			"configurationDirectory",
			"configurationMkdir",
			"dataDirectory",
			"dataMkdir",
			"keyFile",
			"value",
		},
	)

	return schema.ResourceSpec{
		IsComponent: true,
		ObjectTypeSpec: schema.ObjectTypeSpec{
			Description: "Configures etcd on a remote system.",
			Properties:  outputs,
			Required:    requiredOutputs,
		},
		InputProperties: inputs,
		RequiredInputs:  requiredInputs,
	}
}

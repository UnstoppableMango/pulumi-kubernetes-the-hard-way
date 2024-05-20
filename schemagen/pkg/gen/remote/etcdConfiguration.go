package remote

import (
	"maps"
	"slices"

	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/schemagen/pkg/gen/props"
	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/schemagen/pkg/gen/types"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

func generateEtcdConfiguration(commandSpec schema.PackageSpec) schema.PackageSpec {
	inputs := map[string]schema.PropertySpec{
		"caPem":   props.String("The PEM encoded certificate authority data."),
		"certPem": props.String("The PEM encoded certificate data."),
		"configurationDirectory": {
			Description: "The directory to store etcd configuration.",
			TypeSpec:    types.String,
			Default:     "/etc/etcd",
		},
		"connection": props.Connection(commandSpec),
		"dataDirectory": {
			Description: "The directory etcd will store its data.",
			TypeSpec:    types.String,
			Default:     "/var/lib/etcd",
		},
		"etcdPath":   props.String("The path to the `etcd` binary."),
		"internalIp": props.String("The IP used to serve client requests and communicate with etcd peers."),
		"keyPem":     props.String("The PEM encoded key data."), // Secret I'm pretty sure
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
		"configurationChmod": {
			Description: "The configuration chmod operation.",
			TypeSpec:    types.LocalResource("Chmod", "tools"),
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

	requiredOutputs := slices.Concat(requiredInputs, []string{
		"caFile",
		"certFile",
		"configurationDirectory",
		"configurationMkdir",
		"dataDirectory",
		"dataMkdir",
		"keyFile",
		"value",
	})

	return schema.PackageSpec{
		Types: map[string]schema.ComplexTypeSpec{
			name("EtcdConfigurationProps"): { // TODO: This name kinda sucks
				ObjectTypeSpec: schema.ObjectTypeSpec{
					Description: "Props for resources that consume etcd configuration.",
					Type:        "object",
					Properties: map[string]schema.PropertySpec{
						"caFilePath":    props.String("Path to the certificate authority file on the remote system."),
						"certFilePath":  props.String("Path to the certificate file on the remote system."),
						"dataDirectory": props.String("Etcd's data directory."),
						"etcdPath":      props.String("Path to the etcd binary."),
						"internalIp":    props.String("Internal IP of the etcd node."),
						"keyFilePath":   props.String("Path to the private key file on the remote system."),
						"name":          props.String("Name of the etcd node."),
					},
					Required: []string{
						"caFilePath",
						"certFilePath",
						"dataDirectory",
						"etcdPath",
						"internalIp",
						"keyFilePath",
						"name",
					},
				},
			},
		},
		Functions: map[string]schema.FunctionSpec{},
		Resources: map[string]schema.ResourceSpec{
			name("EtcdConfiguration"): {
				IsComponent: true,
				ObjectTypeSpec: schema.ObjectTypeSpec{
					Description: "Configures etcd on a remote system.",
					Properties:  outputs,
					Required:    requiredOutputs,
				},
				InputProperties: inputs,
				RequiredInputs:  requiredInputs,
			},
		},
	}
}

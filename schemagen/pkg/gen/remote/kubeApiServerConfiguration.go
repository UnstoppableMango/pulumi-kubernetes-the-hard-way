package remote

import (
	"maps"
	"slices"

	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/schemagen/pkg/gen/props"
	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/schemagen/pkg/gen/types"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

func generateKubeApiServerConfiguration(commandSpec schema.PackageSpec) schema.ResourceSpec {
	inputs := map[string]schema.PropertySpec{
		"caPem":   props.String("The PEM encoded certificate authority data."),
		"caKey":   props.String("The PEM encoded certificate authority key."),
		"certPem": props.String("The PEM encoded Kube API Server certificate data."),
		"configurationDirectory": {
			Description: "The directory to store Kubernetes Control Plane configuration.",
			TypeSpec:    types.String,
			Default:     "/etc/kubernetes/config",
		},
		"connection": {
			Description: "The parameters with which to connect to the remote host.",
			TypeSpec:    types.ExtType(commandSpec, "Connection", "remote"),
		},
		"encryptionConfig":   props.String("The YAML encryption configuration manifest."),
		"keyPem":             props.String("The PEM encoded Kube API Server certificate key."),
		"kubectlPath":        props.String("The path to the 'kubectl' binary."),
		"path":               props.String("The path to the 'kube-apiserver' binary."),
		"serviceAccountsPem": props.String("The PEM encoded Service Accounts certificate data."),
		"serviceAccountsKey": props.String("The PEM encoded Service Accounts certificate key."),
	}

	requiredInputs := []string{
		"caPem",
		"caKey",
		"connection",
		"encryptionConfig",
		"keyPem",
		"certPem",
		"serviceAccountsKey",
		"serviceAccountsPem",
	}

	outputs := map[string]schema.PropertySpec{
		"configurationMkdir": {
			Description: "Configuration mkdir operation",
			TypeSpec:    types.LocalType("Mkdir", "tools"),
		},
	}
	maps.Copy(outputs, inputs)

	requiredOutputs := slices.Concat(
		requiredInputs,
		[]string{},
	)

	return schema.ResourceSpec{
		IsComponent: true,
		ObjectTypeSpec: schema.ObjectTypeSpec{
			Description: "Configures Kubernetes API Server on a remote system.",
			Properties:  outputs,
			Required:    requiredOutputs,
		},
		InputProperties: inputs,
		RequiredInputs:  requiredInputs,
	}
}

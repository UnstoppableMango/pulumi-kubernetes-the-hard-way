package remote

import (
	"maps"
	"slices"

	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/schemagen/pkg/gen/props"
	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/schemagen/pkg/gen/types"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

func generateKubeControllerManagerConfiguration(commandSpec schema.PackageSpec) schema.ResourceSpec {
	inputs := map[string]schema.PropertySpec{
		"caPem": props.String("The PEM encoded certificate authority data."),
		"caKey": props.String("The PEM encoded certificate authority key."),
		"configurationDirectory": {
			Description: "The directory to store Kubernetes Control Plane configuration.",
			TypeSpec:    types.String,
			Default:     "/etc/kubernetes/config",
		},
		"connection": {
			Description: "The parameters with which to connect to the remote host.",
			TypeSpec:    types.ExtType(commandSpec, "Connection", "remote"),
		},
		"encryptionConfig": props.String("The YAML encryption configuration manifest."),
		"kubeControllerManagerKubeconfig": {
			Description: "The kube-controller-manager kubeconfig configuration",
			TypeSpec:    types.LocalType("Kubeconfig", "config"),
		},
		"kubeControllerManagerPath": props.String("The path to the 'kube-controller-manager' binary."),
		"kubectlPath":               props.String("The path to the 'kubectl' binary."),
		"serviceAccountsPem":        props.String("The PEM encoded Service Accounts certificate data."),
		"serviceAccountsKey":        props.String("The PEM encoded Service Accounts certificate key."),
	}

	requiredInputs := []string{
		"caPem",
		"caKey",
		"connection",
		"encryptionConfig",
		"kubeControllerManagerKubeconfig",
		"serviceAccountsKey",
		"serviceAccountsPem",
	}

	outputs := map[string]schema.PropertySpec{}
	maps.Copy(outputs, inputs)

	requiredOutputs := slices.Concat(
		requiredInputs,
		[]string{},
	)

	return schema.ResourceSpec{
		IsComponent: true,
		ObjectTypeSpec: schema.ObjectTypeSpec{
			Description: "Configures Kubernetes Control Plane on a remote system.",
			Properties:  outputs,
			Required:    requiredOutputs,
		},
		InputProperties: inputs,
		RequiredInputs:  requiredInputs,
	}
}
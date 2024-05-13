package remote

import (
	"maps"
	"slices"

	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/schemagen/pkg/gen/props"
	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/schemagen/pkg/gen/types"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

func generateKubernetesControlPlaneConfiguration() schema.ResourceSpec {
	inputs := map[string]schema.PropertySpec{
		"caPem":               props.String("The PEM encoded certificate authority data."),
		"caKey":               props.String("The PEM encoded certificate authority key."),
		"kubeApiServerPem":    props.String("The PEM encoded Kube API Server certificate data."),
		"kubeApiServerKey":    props.String("The PEM encoded Kube API Server certificate key."),
		"serviceAccountsPem":  props.String("The PEM encoded Service Accounts certificate data."),
		"serviceAccountsKey":  props.String("The PEM encoded Service Accounts certificate key."),
		"encryptionConfig":    props.String("The YAML encryption configuration manifest."),
		"kubeSchedulerConfig": props.String("The kube-scheduler configuration manifest."),
		"kubeControllerManagerKubeconfig": {
			Description: "The kube-controller-manager kubeconfig configuration",
			TypeSpec:    types.LocalType("Kubeconfig", "config"),
		},
		"kubeSchedulerKubeconfig": {
			Description: "The kube-scheduler kubeconfig configuration",
			TypeSpec:    types.LocalType("Kubeconfig", "config"),
		},
		"configurationDirectory": {
			Description: "The directory to store Kubernetes Control Plane configuration.",
			TypeSpec:    types.String,
			Default:     "/etc/kubernetes/config",
		},
		"kubeApiServerPath":         props.String("The path to the 'kube-apiserver' binary."),
		"kubeControllerManagerPath": props.String("The path to the 'kube-controller-manager' binary."),
		"kubeSchedulerPath":         props.String("The path to the 'kube-scheduler' binary."),
		"kubectlPath":               props.String("The path to the 'kubectl' binary."),
	}

	requiredInputs := []string{
		"caPem",
		"caKey",
		"kubeApiServerPem",
		"kubeApiServerKey",
		"serviceAccountsPem",
		"serviceAccountsKey",
		"encryptionConfig",
		"kubeSchedulerConfig",
		"kubeControllerManagerKubeconfig",
		"kubeSchedulerKubeconfig",
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

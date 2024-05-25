package remote

import (
	"maps"
	"slices"

	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/schemagen/pkg/gen/props"
	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/schemagen/pkg/gen/types"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

func generateControlPlaneNode(commandSpec schema.PackageSpec) schema.PackageSpec {
	inputs := map[string]schema.PropertySpec{
		"caCertificatePath":              props.String("The path to the root certificate authority certificate."),
		"caPrivateKeyPath":               props.String("The path to the root certificate authority private key."),
		"connection":                     props.Connection(commandSpec),
		"encryptionConfig":               props.String("The v1/EncryptionConfig yaml."),
		"kubeApiServerCertificatePath":   props.String("The path to the kube-apiserver certificate."),
		"kubeApiServerPrivateKeyPath":    props.String("The path to the kube-apiserver private key."),
		"serviceAccountsCertificatePath": props.String("The path to the service accounts certificate."),
		"serviceAccountsPrivateKeyPath":  props.String("The path to the service accounts private key."),
	}

	requiredInputs := []string{"connection"}

	outputs := map[string]schema.PropertySpec{
		"kubeApiServerInstall": {
			Description: "The kube-apiserver install.",
			TypeSpec:    types.LocalResource("KubeApiServerInstall", "remote"),
		},
		"kubeApiServerService": {
			Description: "The kube-apiserver systemd service.",
			TypeSpec:    types.LocalResource("KubeApiServerService", "remote"),
		},
		"kubeControllerManagerInstall": {
			Description: "The kube-controller-manager install.",
			TypeSpec:    types.LocalResource("KubeControllerManagerInstall", "remote"),
		},
		"kubeControllerManagerService": {
			Description: "The kube-controller-manager systemd service.",
			TypeSpec:    types.LocalResource("KubeControllerManagerService", "remote"),
		},
		"kubectlInstall": {
			Description: "The kubectl install.",
			TypeSpec:    types.LocalResource("KubectlInstall", "remote"),
		},
		"kubernetesConfigurationMkdir": {
			Description: "The kubernetes configuration mkdir operation.",
			TypeSpec:    types.LocalResource("Mkdir", "tools"),
		},
		"kubeSchedulerInstall": {
			Description: "The kube-scheduler isntall.",
			TypeSpec:    types.LocalResource("KubeSchedulerInstall", "remote"),
		},
		"kubeSchedulerService": {
			Description: "The kube-scheduler systemd service.",
		},
		"varLibKubernetesMkdir": {
			Description: "The /var/lib/kubernetes mkdir operation.",
			TypeSpec:    types.LocalResource("Mkdir", "tools"),
		},
	}
	maps.Copy(outputs, inputs)

	requiredOutputs := slices.Concat(requiredInputs, []string{
		"kubeApiServerInstall",
		"kubeControllerManagerInstall",
		"kubectlInstall",
		"kubernetesConfigurationMkdir",
		"kubeSchedulerInstall",
		"varLibKubernetesMkdir",
	})

	return schema.PackageSpec{
		Types:     map[string]schema.ComplexTypeSpec{},
		Functions: map[string]schema.FunctionSpec{},
		Resources: map[string]schema.ResourceSpec{
			name("ControlPlaneNode"): {
				IsComponent: true,
				ObjectTypeSpec: schema.ObjectTypeSpec{
					Description: "A kubernetes control plane node.",
					Properties:  outputs,
					Required:    requiredOutputs,
				},
				InputProperties: inputs,
				RequiredInputs:  requiredInputs,
			},
		},
	}
}

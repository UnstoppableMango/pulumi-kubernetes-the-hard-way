package remote

import (
	"maps"
	"slices"

	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/schemagen/pkg/gen/props"
	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/schemagen/pkg/gen/types"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

func generateControlPlaneNode(commandSpec, commandxSpec schema.PackageSpec) schema.PackageSpec {
	inputs := map[string]schema.PropertySpec{
		"apiServerCount": props.Integer("The number of kube-apiserver instance."),
		"architecture": {
			Description: "The node's CPU architecture.",
			TypeSpec:    types.LocalType("Architecture", "remote"),
		},
		"audiLogPath":                           props.String("The path to store the audit log file."),
		"caCertificatePath":                     props.String("The path to the root certificate authority certificate."),
		"caPrivateKeyPath":                      props.String("The path to the root certificate authority private key."),
		"clusterCIDR":                           props.String("The cluster CIDR."),
		"clusterName":                           props.String("The cluster name."),
		"connection":                            props.Connection(commandSpec),
		"encryptionConfigYaml":                  props.String("The v1/EncryptionConfig yaml."),
		"kubeApiServerCertificatePath":          props.String("The path to the kube-apiserver certificate."),
		"kubeApiServerInstallDirectory":         props.String("The directory to store the kube-apiserver binary."),
		"kubeApiServerPrivateKeyPath":           props.String("The path to the kube-apiserver private key."),
		"kubeControllerManagerInstallDirectory": props.String("The directory to store the kube-controller-manager binary."),
		"kubeControllerManagerKubeconfigPath":   props.String("The path to the kube-controller-manager kubeconfig file."),
		"kubectlInstallDirectory":               props.String("The path to store the kubectl binary."),
		"kubernetesVersion":                     props.String("The version of kubernetes to use."),
		"kubeSchedulerConfigYaml":               props.String("The kube-scheduler config yaml."),
		"kubeSchedulerInstallDirectory":         props.String("The directory to store the kube-scheduler binary."),
		"kubeSchedulerKubeconfigPath":           props.String("The path to the kube-scheduler kubeconfig file."),
		"nodeName":                              props.String("The name of the node."),
		"serviceAccountsCertificatePath":        props.String("The path to the service accounts certificate."),
		"serviceAccountsPrivateKeyPath":         props.String("The path to the service accounts private key."),
		"serviceClusterIpRange":                 props.String("The IP range to use for cluster services."),
	}

	requiredInputs := []string{
		"apiServerCount",
		"architecture",
		"connection",
		"caCertificatePath",
		"caPrivateKeyPath",
		"encryptionConfigYaml",
		"kubeApiServerCertificatePath",
		"kubeApiServerPrivateKeyPath",
		"kubeControllerManagerKubeconfigPath",
		"kubeSchedulerConfigYaml",
		"kubeSchedulerKubeconfigPath",
		"serviceAccountsCertificatePath",
		"serviceAccountsPrivateKeyPath",
	}

	outputs := map[string]schema.PropertySpec{
		"encryptionConfigFile": {
			Description: "The remote encryption config file.",
			TypeSpec:    types.LocalResource("File", "remote"),
		},
		"kubeApiServerInstall": {
			Description: "The kube-apiserver install.",
			TypeSpec:    types.LocalResource("KubeApiServerInstall", "remote"),
		},
		"kubeApiServerService": {
			Description: "The kube-apiserver systemd service.",
			TypeSpec:    types.LocalResource("SystemdService", "remote"),
		},
		"kubeControllerManagerInstall": {
			Description: "The kube-controller-manager install.",
			TypeSpec:    types.LocalResource("KubeControllerManagerInstall", "remote"),
		},
		"kubeControllerManagerService": {
			Description: "The kube-controller-manager systemd service.",
			TypeSpec:    types.LocalResource("SystemdService", "remote"),
		},
		"kubectlInstall": {
			Description: "The kubectl install.",
			TypeSpec:    types.LocalResource("KubectlInstall", "remote"),
		},
		"kubernetesConfigurationMkdir": {
			Description: "The kubernetes configuration mkdir operation.",
			TypeSpec:    types.ExtResource(commandxSpec, "Mkdir", "remote"),
		},
		"kubeSchedulerInstall": {
			Description: "The kube-scheduler isntall.",
			TypeSpec:    types.LocalResource("KubeSchedulerInstall", "remote"),
		},
		"kubeSchedulerService": {
			Description: "The kube-scheduler systemd service.",
			TypeSpec:    types.LocalResource("SystemdService", "remote"),
		},
		"varLibKubernetesMkdir": {
			Description: "The /var/lib/kubernetes mkdir operation.",
			TypeSpec:    types.ExtResource(commandxSpec, "Mkdir", "remote"),
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

package remote

import (
	"maps"

	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/schemagen/pkg/gen/props"
	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/schemagen/pkg/gen/types"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

func generateKubeletService(commandSpec schema.PackageSpec) schema.PackageSpec {
	inputs := map[string]schema.PropertySpec{
		"after": props.ArrayOf("string", "Waits for any units defined here."),
		"configuration": {
			Description: "Kubelet configuration.",
			TypeSpec:    types.LocalType("KubeletConfigurationProps", "remote"),
		},
		"connection":    props.Connection(commandSpec),
		"description":   props.String("Optional systemd unit description."),
		"directory":     props.String("The location to create the service file."),
		"documentation": props.String("Optional systemd unit documentation"),
		"requires":      props.ArrayOf("string", "Requires any units defined here."),
		"restart": {
			Description: "Optionally override the systemd service restart behaviour. Defaults to `on-failure`.",
			TypeSpec:    types.LocalType("SystemdServiceRestart", "remote"),
		},
		"restartSec": props.String("Optionally override the systemd service RestartSec. Defaults to `5`."),
		"wantedBy":   props.String("Optionally override the systemd service wanted-by. Defaults to `multi-user.target`."),
	}

	requiredInputs := []string{"after", "configuration", "connection", "requires"}

	outputs := map[string]schema.PropertySpec{
		"service": {
			Description: "The remote systemd service.",
			TypeSpec:    types.LocalResource("SystemdService", "remote"),
		},
	}
	maps.Copy(outputs, inputs)

	return schema.PackageSpec{
		Types: map[string]schema.ComplexTypeSpec{
			name("KubeletConfigurationProps"): {
				ObjectTypeSpec: schema.ObjectTypeSpec{
					Description: "Props for resources that consume kubelet configuration.",
					Type:        "object",
					Properties: map[string]schema.PropertySpec{
						"configurationFilePath": props.String("Path to the kubelet configuration."),
						"kubeconfigPath":        props.String("Path to the kubeconfig the kubelet will use"),
						"kubeletPath":           props.String("Path to the kubelet binary."),
						"registerNode":          props.Boolean("Whether to register the node. Defaults to `true`."),
						"v":                     props.Integer("Verbosity. Defaults to `2`."),
					},
					Required: []string{
						"configurationFilePath",
						"kubeconfigPath",
						"kubeletPath",
						"registerNode",
						"v",
					},
				},
			},
		},
		Functions: map[string]schema.FunctionSpec{},
		Resources: map[string]schema.ResourceSpec{
			name("KubeletService"): {
				IsComponent: true,
				ObjectTypeSpec: schema.ObjectTypeSpec{
					Description: "Kubelet systemd service file. Will likely get replaced with a static function when https://github.com/pulumi/pulumi/issues/7583 gets resolved.",
					Properties:  outputs,
					Required:    append(requiredInputs, "service"),
				},
				InputProperties: inputs,
				RequiredInputs:  requiredInputs,
			},
		},
	}
}

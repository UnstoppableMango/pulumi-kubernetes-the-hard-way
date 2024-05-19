package config

import (
	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/schemagen/pkg/gen/props"
	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/schemagen/pkg/gen/types"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

func generateFunctions() map[string]schema.FunctionSpec {
	functions := map[string]schema.FunctionSpec{
		name("getKubeconfig"): { // TODO: Need more pems
			Inputs: &schema.ObjectTypeSpec{
				Properties: map[string]schema.PropertySpec{
					"caPem": props.String("Certificate authority data."),
					"options": {
						Description: "Options for creating the kubeconfig.",
						TypeSpec: schema.TypeSpec{
							Plain: true,
							OneOf: []schema.TypeSpec{
								types.LocalType("KubeconfigAdminOptions", "config"),
								types.LocalType("KubeconfigKubeControllerManagerOptions", "config"),
								types.LocalType("KubeconfigKubeProxyOptions", "config"),
								types.LocalType("KubeconfigKubeSchedulerOptions", "config"),
								types.LocalType("KubeconfigWorkerOptions", "config"),
							},
							Discriminator: &schema.DiscriminatorSpec{
								PropertyName: "type",
								Mapping: map[string]string{
									"admin":                   types.LocalTypeRef("KubeconfigAdminOptions", "config"),
									"kube-controller-manager": types.LocalTypeRef("KubeconfigKubeControllerManagerOptions", "config"),
									"kube-proxy":              types.LocalTypeRef("KubeconfigKubeProxyOptions", "config"),
									"kube-scheduler":          types.LocalTypeRef("KubeconfigKubeSchedulerOptions", "config"),
									"worker":                  types.LocalTypeRef("KubeconfigWorkerOptions", "config"),
								},
							},
						},
					},
				},
				Required: []string{"caPem", "options"},
			},
			Outputs: &schema.ObjectTypeSpec{
				Properties: map[string]schema.PropertySpec{
					"result": {
						TypeSpec: types.LocalType("Kubeconfig", "config"),
					},
				},
				Required: []string{"result"},
			},
		},
		name("getKubeletConfiguration"): generateGetKubeletConfiguration(),
		name("getKubeVipManifest"):      generateGetKubeVipManifest(),
	}

	return functions
}

func generateGetKubeVipManifest() schema.FunctionSpec {
	return schema.FunctionSpec{
		Description: "Gets the static pod manifests for KubeVip.",
		Inputs: &schema.ObjectTypeSpec{
			Properties: map[string]schema.PropertySpec{
				"address":        props.String("TODO"),
				"bgpAs":          props.Integer("TODO"),
				"bgpEnable":      props.Boolean("TODO"),
				"bgpPeerAddress": props.String("TODO"),
				"bgpPeerAs":      props.Integer("TODO"),
				"bgpPeerPass":    props.String("TODO"),
				"bgpPeers":       props.String("TODO"), // This could be structured so that consumers don't need to know the format
				"bgpRouterId":    props.String("TODO"),
				"cpEnable":       props.Boolean("TODO"),
				"cpNamespace":    props.String("TODO"),
				"image":          props.String("Override the kube-vip image."),
				"kubeconfigPath": props.String("Path to the kubeconfig on the remote host."),
				"name": {
					Description: "Name of the static pod. Defaults to kube-vip.",
					TypeSpec:    types.String,
					Default:     "kube-vip",
				},
				"namespace": {
					Description: "Namespace for the static pod. Defaults to kube-system.",
					TypeSpec:    types.String,
					Default:     "kube-system",
				},
				"port": {
					Description: "TODO",
					TypeSpec:    types.Integer,
					Default:     6443,
				},
				"svcEnable":         props.Boolean("TODO"),
				"version":           props.String("Version of kube-vip to use."),
				"vipArp":            props.Boolean("TODO"),
				"vipCidr":           props.Integer("TODO"),
				"vipDdns":           props.Boolean("TODO"),
				"vipInterface":      props.String("TODO"),
				"vipLeaderElection": props.Boolean("TODO"),
				"vipLeaseDuration":  props.Integer("TODO"),
				"vipRenewDeadline":  props.Integer("TODO"),
				"vipRetryPeriod":    props.Integer("TODO"),
			},
			Required: []string{"address", "kubeconfigPath", "vipCidr"},
		},
		Outputs: &schema.ObjectTypeSpec{
			Properties: map[string]schema.PropertySpec{
				"result": {
					TypeSpec: types.LocalType("PodManifest", "config"),
				},
			},
			Required: []string{"result"},
		},
	}
}

func generateGetKubeletConfiguration() schema.FunctionSpec {
	return schema.FunctionSpec{
		Description: "Get the kubelet configuration.",
		Inputs: &schema.ObjectTypeSpec{
			Properties: map[string]schema.PropertySpec{
				"anonymous":                props.Boolean(""),
				"webhook":                  props.Boolean(""),
				"clientCAFile":             props.String(""),
				"authorizationMode":        props.String(""),
				"clusterDomain":            props.String(""),
				"clusterDNS":               props.ArrayOf("string", ""),
				"cgroupDriver":             props.String(""),
				"containerRuntimeEndpoint": props.String(""),
				"podCIDR":                  props.String(""),
				"resolvConf":               props.String(""),
				"runtimeRequestTimeout":    props.String(""),
				"tlsCertFile":              props.String(""),
				"tlsPrivateKeyFile":        props.String(""),
			},
			Required: []string{"podCIDR"},
		},
		Outputs: &schema.ObjectTypeSpec{
			Properties: map[string]schema.PropertySpec{
				"result": {
					TypeSpec: types.LocalType("KubeletConfiguration", "config"),
				},
			},
			Required: []string{"result"},
		},
	}
}

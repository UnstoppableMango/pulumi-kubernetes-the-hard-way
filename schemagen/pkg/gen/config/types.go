package config

import (
	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/schemagen/pkg/gen/internal"
	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/schemagen/pkg/gen/props"
	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/schemagen/pkg/gen/types"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

func generateTypes(kubernetesSpec schema.PackageSpec) map[string]schema.ComplexTypeSpec {
	types := map[string]schema.ComplexTypeSpec{
		name("KubeletConfigurationAuthenticationAnonymous"): {
			ObjectTypeSpec: schema.ObjectTypeSpec{
				Type: "object",
				Properties: map[string]schema.PropertySpec{
					"enabled": props.Boolean("TODO"),
				},
				Required: []string{"enabled"},
			},
		},
		name("KubeletConfigurationAuthenticationWebhook"): {
			ObjectTypeSpec: schema.ObjectTypeSpec{
				Type: "object",
				Properties: map[string]schema.PropertySpec{
					"enabled": props.Boolean("TODO"),
				},
				Required: []string{"enabled"},
			},
		},
		name("KubeletConfigurationAuthenticationx509"): {
			ObjectTypeSpec: schema.ObjectTypeSpec{
				Type: "object",
				Properties: map[string]schema.PropertySpec{
					"clientCAFile": props.String("TODO"),
				},
				Required: []string{"clientCAFile"},
			},
		},
		name("KubeletConfigurationAuthentication"): {
			ObjectTypeSpec: schema.ObjectTypeSpec{
				Type: "object",
				Properties: map[string]schema.PropertySpec{
					"anonymous": {
						TypeSpec: types.LocalType("KubeletConfigurationAuthenticationAnonymous", "config"),
					},
					"webhook": {
						TypeSpec: types.LocalType("KubeletConfigurationAuthenticationWebhook", "config"),
					},
					"x509": {
						TypeSpec: types.LocalType("KubeletConfigurationAuthenticationx509", "config"),
					},
				},
				Required: []string{"anonymous", "webhook", "x509"},
			},
		},
		name("KubeletConfigurationAuthorization"): {
			ObjectTypeSpec: schema.ObjectTypeSpec{
				Type: "object",
				Properties: map[string]schema.PropertySpec{
					"mode": props.String(""),
				},
				Required: []string{"mode"},
			},
		},
		name("KubeletConfiguration"): {
			ObjectTypeSpec: schema.ObjectTypeSpec{
				Type: "object",
				Properties: map[string]schema.PropertySpec{
					"kind": {
						TypeSpec: types.String,
						Const:    "KubeletConfiguration",
					},
					"apiVersion": {
						TypeSpec: types.String,
						Const:    "kubelet.config.k8s.io/v1beta1",
					},
					"authentication": {
						TypeSpec: types.LocalType("KubeletConfigurationAuthentication", "config"),
					},
					"authorization": {
						TypeSpec: types.LocalType("KubeletConfigurationAuthorization", "config"),
					},
					"clusterDomain":            props.String("TODO"),
					"clusterDNS":               props.ArrayOf("string", "TODO"),
					"cgroupDriver":             props.String("TODO"),
					"containerRuntimeEndpoint": props.String("TODO"),
					"podCIDR":                  props.String("TODO"),
					"resolvConf":               props.String("TODO"),
					"runtimeRequestTimeout":    props.String("TODO"),
					"tlsCertFile":              props.String("TODO"),
					"tlsPrivateKeyFile":        props.String("TODO"),
				},
				Required: []string{
					"kind",
					"apiVersion",
					"authentication",
					"authorization",
					"clusterDomain",
					"clusterDNS",
					"cgroupDriver",
					"containerRuntimeEndpoint",
					"podCIDR",
					"resolvConf",
					"runtimeRequestTimeout",
					"tlsCertFile",
					"tlsPrivateKeyFile",
				},
			},
		},
		name("KubeProxyConfigurationClientConnection"): {
			ObjectTypeSpec: schema.ObjectTypeSpec{
				Type: "object",
				Properties: map[string]schema.PropertySpec{
					"kubeconfig": props.String("Path to the kubeconfig."),
				},
				Required: []string{"kubeconfig"},
			},
		},
		name("KubeProxyConfiguration"): {
			ObjectTypeSpec: schema.ObjectTypeSpec{
				Type: "object",
				Properties: map[string]schema.PropertySpec{
					"kind": {
						TypeSpec: types.String,
						Const:    "KubeProxyConfiguration",
					},
					"apiVersion": {
						TypeSpec: types.String,
						Const:    "kubeproxy.config.k8s.io/v1alpha1",
					},
					"clientConnection": {
						TypeSpec: types.LocalType("KubeProxyConfigurationClientConnection", "config"),
					},
					"mode":        props.String("TODO"),
					"clusterCIDR": props.String("TODO"),
				},
				Required: []string{"clusterCIDR"},
			},
		},
		name("PodManifest"): generatePodManifest(kubernetesSpec),
		name("User"): {
			ObjectTypeSpec: schema.ObjectTypeSpec{
				Type: "object",
				Properties: map[string]schema.PropertySpec{
					"clientCertificateData": props.String("TODO"),
					"clientKeyData":         props.String("TODO"),
				},
				Required: []string{"clientCertificateData", "clientKeyData"},
			},
		},
	}

	return types
}

func generatePodManifest(kubernetesSpec schema.PackageSpec) schema.ComplexTypeSpec {
	pod := kubernetesSpec.Resources["kubernetes:core/v1:Pod"]

	return schema.ComplexTypeSpec{
		ObjectTypeSpec: schema.ObjectTypeSpec{
			Description: pod.Description,
			Type:        "object",
			Properties:  internal.MakeExternal(pod.Properties, kubernetesSpec),
		},
	}
}

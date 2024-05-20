package config

import (
	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/schemagen/pkg/gen/props"
	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/schemagen/pkg/gen/types"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

func generateGetKubeletConfiguration() schema.PackageSpec {
	function, resource := generatePseudoFunction(
		"Get the kubelet configuration.",
		schema.ObjectTypeSpec{
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
		schema.PropertySpec{
			TypeSpec: types.LocalType("KubeletConfiguration", "config"),
		},
		"yaml", props.String("The yaml representation of the manifest."),
	)

	return schema.PackageSpec{
		Types: map[string]schema.ComplexTypeSpec{
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
		},
		Functions: map[string]schema.FunctionSpec{
			name("getKubeletConfiguration"): function,
		},
		Resources: map[string]schema.ResourceSpec{
			name("KubeletConfiguration"): resource,
		},
	}
}

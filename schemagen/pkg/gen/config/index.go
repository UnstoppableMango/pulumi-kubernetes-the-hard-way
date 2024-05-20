package config

import (
	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/schemagen/pkg/gen/internal"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

func GenerateConfig(kubernetesSpec schema.PackageSpec) schema.PackageSpec {
	base := schema.PackageSpec{
		Types:     generateTypes(kubernetesSpec),
		Resources: map[string]schema.ResourceSpec{},
		Functions: map[string]schema.FunctionSpec{
			name("getKubeconfig"): generateGetKubeconfig(),
		},
	}

	return internal.ExtendSchemas(base,
		generateGetKubeletConfiguration(),
		generateGetKubeProxyConfiguration(),
		generateGetKubeVipManifest(),
	)
}

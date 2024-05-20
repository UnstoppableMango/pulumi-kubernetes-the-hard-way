package config

import (
	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/schemagen/pkg/gen/internal"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

func GenerateConfig(kubernetesSpec schema.PackageSpec) schema.PackageSpec {
	base := schema.PackageSpec{
		Types: map[string]schema.ComplexTypeSpec{
			name("PodManifest"): generatePodManifest(kubernetesSpec),
		},
		Resources: map[string]schema.ResourceSpec{},
		Functions: map[string]schema.FunctionSpec{},
	}

	return internal.ExtendSchemas(base,
		generateGetCniBridgePluginConfiguration(),
		generateGetCniLoopbackPluginConfiguration(),
		generateContainerdConfiguration(),
		generateGetKubeconfig(),
		generateGetKubeletConfiguration(),
		generateGetKubeProxyConfiguration(),
		generateGetKubeVipManifest(),
	)
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

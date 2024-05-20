package config

import (
	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/schemagen/pkg/gen/internal"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

func generateTypes(kubernetesSpec schema.PackageSpec) map[string]schema.ComplexTypeSpec {
	types := map[string]schema.ComplexTypeSpec{
		name("PodManifest"): generatePodManifest(kubernetesSpec),
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

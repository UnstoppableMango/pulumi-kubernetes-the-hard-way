package config

import (
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

func GenerateConfig(kubernetesSpec schema.PackageSpec) schema.PackageSpec {
	functions := generateFunctions()

	resources := map[string]schema.ResourceSpec{
		name("KubeVipManifest"): generateKubeVipManifest(functions[name("getKubeVipManifest")]),
	}

	return schema.PackageSpec{
		Functions: functions,
		Resources: resources,
		Types:     generateTypes(kubernetesSpec),
	}
}

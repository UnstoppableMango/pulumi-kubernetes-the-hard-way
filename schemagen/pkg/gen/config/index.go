package config

import (
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

const configMod = "kubernetes-the-hard-way:config:"

func GenerateConfig(kubernetesSpec schema.PackageSpec) schema.PackageSpec {
	functions := generateFunctions()

	resources := map[string]schema.ResourceSpec{
		configMod + "KubeVipManifest": generateKubeVipManifest(functions[configMod+"getKubeVipManifest"]),
	}

	return schema.PackageSpec{
		Functions: functions,
		Resources: resources,
		Types:     generateTypes(kubernetesSpec),
	}
}

package config

import (
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

func GenerateConfig(kubernetesSpec schema.PackageSpec) schema.PackageSpec {
	functions := generateFunctions()
	getKubeletConfiguration := functions[name("getKubeletConfiguration")]
	getKubeVipManifest := functions[name("getKubeVipManifest")]

	resources := map[string]schema.ResourceSpec{
		name("KubeletConfiguration"): generateKubeletConfiguration(getKubeletConfiguration),
		name("KubeVipManifest"):      generateKubeVipManifest(getKubeVipManifest),
	}

	return schema.PackageSpec{
		Functions: functions,
		Resources: resources,
		Types:     generateTypes(kubernetesSpec),
	}
}

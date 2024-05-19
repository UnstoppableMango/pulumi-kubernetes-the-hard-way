package config

import (
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

func GenerateConfig(kubernetesSpec schema.PackageSpec) schema.PackageSpec {
	getKubeletConfiguration := generateGetKubeletConfiguration()
	getKubeVipManifest := generateGetKubeVipManifest()

	return schema.PackageSpec{
		Functions: map[string]schema.FunctionSpec{
			name("getKubeconfig"):           generateGetKubeconfig(),
			name("getKubeletConfiguration"): getKubeletConfiguration.Function,
			name("getKubeVipManifest"):      getKubeVipManifest.Function,
		},
		Resources: map[string]schema.ResourceSpec{
			name("KubeletConfiguration"): getKubeletConfiguration.Resource,
			name("KubeVipManifest"):      getKubeVipManifest.Resource,
		},
		Types: generateTypes(kubernetesSpec),
	}
}

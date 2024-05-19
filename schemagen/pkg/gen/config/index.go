package config

import (
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

func GenerateConfig(kubernetesSpec schema.PackageSpec) schema.PackageSpec {
	getKubeconfig := generateGetKubeconfig()
	getKubeletConfiguration := generateGetKubeletConfiguration()
	getKubeVipManifest := generateGetKubeVipManifest()

	return schema.PackageSpec{
		Functions: map[string]schema.FunctionSpec{
			name("getKubeconfig"):           getKubeconfig.Function,
			name("getKubeletConfiguration"): getKubeletConfiguration.Function,
			name("getKubeVipManifest"):      getKubeVipManifest.Function,
		},
		Resources: map[string]schema.ResourceSpec{
			name("Kubeconfig"):           getKubeconfig.Resource,
			name("KubeletConfiguration"): getKubeletConfiguration.Resource,
			name("KubeVipManifest"):      getKubeVipManifest.Resource,
		},
		Types: generateTypes(kubernetesSpec),
	}
}

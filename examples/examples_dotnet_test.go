//go:build dotnet || all

package examples

import (
	"path"
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
)

func TestSimpleDotnet(t *testing.T) {
	test := getCsharpBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir:           path.Join(getCwd(t), "simple-dotnet"),
			Quick:         true,
			SkipRefresh:   true,
			RunUpdateTest: false,
		})

	integration.ProgramTest(t, &test)
}

func getCsharpBaseOptions(t *testing.T) integration.ProgramTestOptions {
	return getBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dependencies: []string{
				"UnMango.KubernetesTheHardWay",
			},
		})
}

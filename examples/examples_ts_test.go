//go:build nodejs || all

package examples

import (
	"path"
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
)

func TestSimpleTs(t *testing.T) {
	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir:           path.Join(getCwd(t), "simple"),
			Quick:         true,
			SkipRefresh:   true,
			RunUpdateTest: false,
		})

	integration.ProgramTest(t, &test)
}

func getJSBaseOptions(t *testing.T) integration.ProgramTestOptions {
	base := getBaseOptions(t)
	baseJS := base.With(integration.ProgramTestOptions{
		Dependencies: []string{
			"@unmango/pulumi-kubernetes-the-hard-way",
		},
	})

	return baseJS
}

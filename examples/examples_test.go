package examples

import (
	"os"
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
)

func getBaseOptions() integration.ProgramTestOptions {
	return integration.ProgramTestOptions{
		//ExpectRefreshChanges: true,
		//RetryFailedSteps:     true,
	}
}

func getCwd(t *testing.T) string {
	cwd, err := os.Getwd()
	if err != nil {
		t.FailNow()
	}

	return cwd
}

func getJSBaseOptions() integration.ProgramTestOptions {
	base := getBaseOptions()
	baseJS := base.With(integration.ProgramTestOptions{
		Dependencies: []string{
			"@unmango/pulumi-kubernetes-the-hard-way",
		},
		Verbose: true,
	})

	return baseJS
}

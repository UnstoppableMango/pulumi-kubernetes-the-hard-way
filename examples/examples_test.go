package examples

import (
	"os"
	"path"
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
)

func getBaseOptions(t *testing.T) integration.ProgramTestOptions {
	return integration.ProgramTestOptions{
		Bin: path.Join(getCwd(t), "..", ".pulumi", "bin", "pulumi"),
		LocalProviders: []integration.LocalDependency{{
			Package: "kubernetes-the-hard-way",
			Path:    path.Join(getCwd(t), "..", "bin"),
		}},
	}
}

func getCwd(t *testing.T) string {
	cwd, err := os.Getwd()
	if err != nil {
		t.FailNow()
	}

	return cwd
}

func skipIfShort(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping long-running test in short mode")
	}
}

func skipIfCi(t *testing.T) {
	if x, ok := os.LookupEnv("CI"); ok && x == "true" {
		t.Skip("skipping resource-intensive test in CI")
	}
}

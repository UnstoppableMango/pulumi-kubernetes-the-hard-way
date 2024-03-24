//go:build python || all

package examples

import (
	"os"
	"path"
	"path/filepath"
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
)

func TestSimplePython(t *testing.T) {
	test := getPythonBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir:           path.Join(getCwd(t), "simple-py"),
			Quick:         true,
			SkipRefresh:   true,
			RunUpdateTest: false,
		})

	integration.ProgramTest(t, &test)
}

func getPythonBaseOptions(t *testing.T) integration.ProgramTestOptions {
	base := getBaseOptions(t)
	sdkPath := filepath.Join("..", "sdk", "python", "bin")
	if _, err := os.Stat(sdkPath); os.IsNotExist(err) {
		t.Fatalf("python SDK not found at %s, run `make build_python` first", sdkPath)
	}
	basePy := base.With(integration.ProgramTestOptions{
		Dependencies: []string{sdkPath},
	})

	return basePy
}

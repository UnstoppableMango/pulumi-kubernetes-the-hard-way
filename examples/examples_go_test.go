//go:build go || all

package examples

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
	"github.com/stretchr/testify/require"
)

const sdkPath = "../sdk"

func TestSimpleGo(t *testing.T) {
	test := getGoBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir:           path.Join(getCwd(t), "simple-go"),
			Quick:         true,
			SkipRefresh:   true,
			RunUpdateTest: false,
		})

	integration.ProgramTest(t, &test)
}

func getGoBaseOptions(t *testing.T) integration.ProgramTestOptions {
	base := getBaseOptions(t)

	rootSdkPath, err := filepath.Abs(sdkPath)
	require.NoError(t, err)

	goDepRoot := os.Getenv("PULUMI_GO_DEP_ROOT")
	if goDepRoot == "" {
		goDepRoot, err = filepath.Abs("../..")
		if err != nil {
			t.Fatal(err)
		}
	}

	baseGo := base.With(integration.ProgramTestOptions{
		Dependencies: []string{
			fmt.Sprintf("github.com/unstoppablemango/pulumi-kubernetes-the-hard-way/sdk=%s", rootSdkPath),
		},
		Env: []string{
			fmt.Sprintf("PULUMI_GO_DEP_ROOT=%s", goDepRoot),
		},
	})

	return baseGo
}

package examples

import (
	"path"
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
)

func TestSimple(t *testing.T) {
	test := getJSBaseOptions().
		With(integration.ProgramTestOptions{
			Dir:           path.Join(getCwd(t), "./simple"),
			Quick:         true,
			SkipRefresh:   true,
			RunUpdateTest: false,
		})

	integration.ProgramTest(t, &test)
}

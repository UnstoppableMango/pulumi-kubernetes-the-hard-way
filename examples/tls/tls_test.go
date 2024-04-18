//go:build nodejs || all

package tls

import (
	"path"
	"testing"

	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/examples"
	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
	"github.com/pulumi/pulumi/sdk/v3/go/common/apitype"
	"github.com/stretchr/testify/assert"
)

func TestTlsRootCaTs(t *testing.T) {
	validateSimple := func(t *testing.T, res apitype.ResourceV3) {
		assert.NotEmpty(t, res.Outputs)
		allowedUses, ok := res.Outputs["allowedUses"]
		assert.True(t, ok, "Outputs did not contain `allowedUses`")
		assert.ElementsMatch(t,
			[]string{"cert_signing", "key_encipherment", "server_auth", "client_auth"},
			allowedUses,
		)

		algorithm, ok := res.Outputs["algorithm"]
		assert.True(t, ok, "Outputs did not contain `algorithm`")
		assert.Equal(t, "RSA", algorithm)

		validityPeriodHours, ok := res.Outputs["validityPeriodHours"]
		assert.True(t, ok, "Outputs did not contain `validityPeriodHours`")
		assert.Equal(t, 256., validityPeriodHours)

		assert.Contains(t, res.Outputs, "cert")
		assert.Contains(t, res.Outputs, "certPem")
		assert.Contains(t, res.Outputs, "key")
		assert.Contains(t, res.Outputs, "privateKeyPem")
		assert.Contains(t, res.Outputs, "publicKeyPem")
	}

	test := examples.GetJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir:   path.Join(examples.GetCwd(t), "tls", "root-ca-ts"),
			Quick: true,
			ExtraRuntimeValidation: func(t *testing.T, stack integration.RuntimeValidationStackInfo) {
				validatedResources := []string{}
				for _, res := range stack.Deployment.Resources {
					if res.Type != "kubernetes-the-hard-way:tls:RootCa" {
						continue
					}
					switch res.URN.Name() {
					case "simple":
						validateSimple(t, res)
						validatedResources = append(validatedResources, "simple")
					}
				}

				assert.Equal(t, []string{"simple"}, validatedResources, "Not all resources were validated")
			},
		})

	integration.ProgramTest(t, &test)
}

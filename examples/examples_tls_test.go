//go:build nodejs || all

package examples

import (
	"path"
	"testing"

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

	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir:   path.Join(getCwd(t), "tls", "root-ca-ts"),
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

func TestTlsCertificateTs(t *testing.T) {
	validateSimple := func(t *testing.T, res apitype.ResourceV3) {
		assert.NotEmpty(t, res.Outputs)
		allowedUses, ok := res.Outputs["allowedUses"]
		assert.True(t, ok, "Outputs did not contain `allowedUses`")
		assert.ElementsMatch(t, []string{"digital_signature"}, allowedUses)

		algorithm, ok := res.Outputs["algorithm"]
		assert.True(t, ok, "Outputs did not contain `algorithm`")
		assert.Equal(t, "RSA", algorithm)

		validityPeriodHours, ok := res.Outputs["validityPeriodHours"]
		assert.True(t, ok, "Outputs did not contain `validityPeriodHours`")
		assert.Equal(t, 69., validityPeriodHours)

		assert.Contains(t, res.Outputs, "cert")
		assert.Contains(t, res.Outputs, "certPem")
		assert.Contains(t, res.Outputs, "key")
		assert.Contains(t, res.Outputs, "privateKeyPem")
		assert.Contains(t, res.Outputs, "publicKeyPem")
	}

	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir:   path.Join(getCwd(t), "tls", "certificate-ts"),
			Quick: true,
			ExtraRuntimeValidation: func(t *testing.T, stack integration.RuntimeValidationStackInfo) {
				validatedResources := []string{}
				for _, res := range stack.Deployment.Resources {
					if res.Type != "kubernetes-the-hard-way:tls:Certificate" {
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

func TestTlsClusterPkiTs(t *testing.T) {
	validateSimple := func(t *testing.T, res apitype.ResourceV3) {
		assert.NotEmpty(t, res.Outputs)

		algorithm, ok := res.Outputs["algorithm"]
		assert.True(t, ok, "Outputs did not contain `algorithm`")
		assert.Equal(t, "RSA", algorithm)

		clusterName, ok := res.Outputs["clusterName"]
		assert.True(t, ok, "Outputs did not contain `clusterName`")
		assert.Equal(t, "my-cluster", clusterName)

		ecdsaCurve, ok := res.Outputs["ecdsaCurve"]
		assert.True(t, ok, "Outputs did not contain `ecdsaCurve`")
		assert.Equal(t, "P256", ecdsaCurve)

		assert.Contains(t, res.Outputs, "admin")
		assert.Contains(t, res.Outputs, "ca")
		assert.Contains(t, res.Outputs, "controllerManager")
		assert.Contains(t, res.Outputs, "kubeProxy")
		assert.Contains(t, res.Outputs, "kubeScheduler")
		assert.Contains(t, res.Outputs, "kubernetes")

		kubelet, ok := res.Outputs["kubelet"]
		assert.True(t, ok, "Outputs did not contain `kubelet`")

		kubeletMap, ok := kubelet.(map[string]interface{})
		assert.True(t, ok, "Output `kubelet` was not the correct type")

		assert.Contains(t, kubeletMap, "node0")

		node1, ok := kubeletMap["node1"]
		assert.True(t, ok, "Output `kubelet` did not contain `node1`")

		_, ok = node1.(map[string]interface{})
		assert.True(t, ok, "Certificate `node1` was not the correct type")
		// assert.Contains(t, node1Cert, "subject") // Hmmm what's going on here
	}

	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir:   path.Join(getCwd(t), "tls", "cluster-pki-ts"),
			Quick: true,
			ExtraRuntimeValidation: func(t *testing.T, stack integration.RuntimeValidationStackInfo) {
				node0Pem, ok := stack.Outputs["node0Pem"]
				assert.True(t, ok, "Outputs did not contain `node0Pem`")
				assert.NotEmpty(t, node0Pem)

				validatedResources := []string{}
				for _, res := range stack.Deployment.Resources {
					if res.Type != "kubernetes-the-hard-way:tls:ClusterPki" {
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

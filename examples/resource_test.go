package examples

import (
	"fmt"
	"path"
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
	"github.com/pulumi/pulumi/sdk/v3/go/common/apitype"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/maps"
)

type ResourceValidator func(t *testing.T, res apitype.ResourceV3)

type validatorToken struct {
	Type      tokens.Type
	Name      string
	Validator ResourceValidator
}

type validatorKey struct {
	Type tokens.Type
	Name string
}

func Validate(typ tokens.Type, name string, validator ResourceValidator) validatorToken {
	return validatorToken{Type: typ, Name: name, Validator: validator}
}

func ResourceValidation(validators ...validatorToken) integration.ProgramTestOptions {
	return integration.ProgramTestOptions{
		ExtraRuntimeValidation: func(t *testing.T, stack integration.RuntimeValidationStackInfo) {
			lookup := map[validatorKey]ResourceValidator{}
			for _, token := range validators {
				key := validatorKey{Type: token.Type, Name: token.Name}
				lookup[key] = token.Validator
			}

			validated := []validatorKey{}
			for _, res := range stack.Deployment.Resources {
				key := validatorKey{Type: res.Type, Name: res.URN.Name()}
				testName := fmt.Sprintf("Testing resource %s", res.URN)
				if validator, ok := lookup[key]; ok && t.Run(testName, func(t *testing.T) {
					validator(t, res)
				}) {
					validated = append(validated, key)
				}
			}

			assert.ElementsMatch(t, maps.Keys(lookup), validated)
		},
	}
}

func ResourceTest(t *testing.T) {
	validateBridge := func(t *testing.T, res apitype.ResourceV3) {
		assert.NotEmpty(t, res.Outputs)

		bridge, ok := res.Outputs["bridge"]
		assert.True(t, ok, "Output `bridge` was not set")
		assert.Equal(t, "cni0", bridge)

		name, ok := res.Outputs["name"]
		assert.True(t, ok, "Output `name` was not set")
		assert.Equal(t, "bridge", name)

		typ, ok := res.Outputs["type"]
		assert.True(t, ok, "Output `type` was not set")
		assert.Equal(t, "bridge", typ)

		cniVersion, ok := res.Outputs["cniVersion"]
		assert.True(t, ok, "Output `cniVersion` was not set")
		assert.Equal(t, "1.0.0", cniVersion)

		path, ok := res.Outputs["path"]
		assert.True(t, ok, "Output `path` was not set")
		assert.Equal(t, "/etc/cni/net.d/10-bridge.conf", path)

		subnet, ok := res.Outputs["subnet"]
		assert.True(t, ok, "Output `subnet` was not set")
		assert.Equal(t, "10.0.69.0/24", subnet)

		isGateway, ok := res.Outputs["isGateway"]
		assert.True(t, ok, "Output `isGateway` was not set")
		assert.Equal(t, true, isGateway)

		ipMasq, ok := res.Outputs["ipMasq"]
		assert.True(t, ok, "Output `ipMasq` was not set")
		assert.Equal(t, true, ipMasq)

		assert.Contains(t, res.Outputs, "file")
		assert.Contains(t, res.Outputs, "ipam")
	}

	validateLoopback := func(t *testing.T, res apitype.ResourceV3) {
		assert.NotEmpty(t, res.Outputs)

		cniVersion, ok := res.Outputs["cniVersion"]
		assert.True(t, ok, "Output `cniVersion` was not set")
		assert.Equal(t, "1.1.0", cniVersion)

		name, ok := res.Outputs["name"]
		assert.True(t, ok, "Output `name` was not set")
		assert.Equal(t, "lo", name)

		path, ok := res.Outputs["path"]
		assert.True(t, ok, "Output `path` was not set")
		assert.Equal(t, "/etc/cni/net.d/99-loopback.conf", path)

		typ, ok := res.Outputs["type"]
		assert.True(t, ok, "Output `type` was not set")
		assert.Equal(t, "loopback", typ)

		assert.Contains(t, res.Outputs, "file")
	}

	validatePlugins := func(t *testing.T, res apitype.ResourceV3) {
		assert.NotEmpty(t, res.Outputs)

		subnet, ok := res.Outputs["subnet"]
		assert.True(t, ok, "Output `subnet` was not set")
		assert.Equal(t, "10.0.69.0/24", subnet)

		directory, ok := res.Outputs["directory"]
		assert.True(t, ok, "Output `directory` was not set")
		assert.Equal(t, "/etc/cni2/net.d", directory)

		assert.Contains(t, res.Outputs, "bridge")
		assert.Contains(t, res.Outputs, "connection")
		assert.Contains(t, res.Outputs, "loopback")
		assert.Contains(t, res.Outputs, "mkdir")
	}

	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: path.Join(getCwd(t), "remote", "cni-plugins-ts"),
			Config: map[string]string{
				"host":     "localhost",
				"port":     node.Port,
				"user":     username,
				"password": password,
			},
			ExtraRuntimeValidation: func(t *testing.T, stack integration.RuntimeValidationStackInfo) {
				validatedResources := []string{}
				for _, res := range stack.Deployment.Resources {
					switch res.Type {
					case "kubernetes-the-hard-way:remote:CniBridgePluginConfiguration":
						switch res.URN.Name() {
						case "simple":
							validateBridge(t, res)
							validatedResources = append(validatedResources, "simple")
						}
					case "kubernetes-the-hard-way:remote:CniLoopbackPluginConfiguration":
						switch res.URN.Name() {
						case "simple":
							validateLoopback(t, res)
							validatedResources = append(validatedResources, "simple")
						}
					case "kubernetes-the-hard-way:remote:CniPluginConfiguration":
						switch res.URN.Name() {
						case "all":
							validatePlugins(t, res)
							validatedResources = append(validatedResources, "all")
						}
					}
				}

				assert.ElementsMatch(t, []string{
					"simple",
					"simple",
					"all",
				}, validatedResources, "Not all resources were validated")
			},
		})

	integration.ProgramTest(t, &test)
}

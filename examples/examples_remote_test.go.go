//go:build remote || all

package examples

import (
	"testing"

	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/examples/internal/rt"
	"github.com/pulumi/pulumi/sdk/v3/go/common/apitype"
	"github.com/stretchr/testify/assert"
)

func TestCniPluginsTs(t *testing.T) {
	rt.ResourceTest(t, "remote/cni-plugins-ts", getJSBaseOptions(t), func(ctx *rt.ResourceContext) {
		rt.Validate(ctx, "kubernetes-the-hard-way:config:CniBridgePluginConfiguration", "simple", func(t *testing.T, res apitype.ResourceV3) {
			assert.NotEmpty(t, res.Outputs)

			rt.ExpectOutput(t, res, "result", map[string]interface{}{
				"name":       "bridge",
				"type":       "bridge",
				"bridge":     "cni0",
				"cniVersion": "1.0.0",
				"isGateway":  true,
				"ipMasq":     true,
				"subnet":     "10.0.69.0/24",
				"ipam": map[string]interface{}{
					"type": "host-local",
					"ranges": []interface{}{
						map[string]interface{}{"subnet": "10.0.69.0/24"},
					},
					"routes": []interface{}{
						map[string]interface{}{"dst": "0.0.0.0/0"},
					},
				},
			})

			assert.Contains(t, res.Outputs, "yaml")
		})
		rt.Validate(ctx, "kubernetes-the-hard-way:config:CniLoopbackPluginConfiguration", "simple", func(t *testing.T, res apitype.ResourceV3) {
			assert.NotEmpty(t, res.Outputs)

			rt.ExpectOutput(t, res, "result", map[string]interface{}{
				"name":       "lo",
				"type":       "loopback",
				"cniVersion": "1.1.0",
			})

			assert.Contains(t, res.Outputs, "yaml")
		})
	})
}

package tools

import (
	"fmt"

	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

func implicitOutputs(inputs, outputs map[string]schema.PropertySpec) map[string]schema.PropertySpec {
	for k, v := range inputs {
		outputs[k] = v
	}

	return outputs
}

func name(x string) string {
	return fmt.Sprintf("kubernetes-the-hard-way:tools:%s", x)
}

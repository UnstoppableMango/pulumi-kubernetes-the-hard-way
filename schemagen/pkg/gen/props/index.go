package props

import (
	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/schemagen/pkg/gen/types"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

func Integer(description string) schema.PropertySpec {
	return schema.PropertySpec{
		Description: description,
		TypeSpec:    types.Integer,
	}
}

func String(description string) schema.PropertySpec {
	return schema.PropertySpec{
		Description: description,
		TypeSpec:    types.String,
	}
}

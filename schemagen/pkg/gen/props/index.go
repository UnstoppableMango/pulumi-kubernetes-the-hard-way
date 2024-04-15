package props

import (
	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/schemagen/pkg/gen/types"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

func ArrayOf(typ string, description string) schema.PropertySpec {
	return schema.PropertySpec{
		Description: description,
		TypeSpec:    types.ArrayOf(typ),
	}
}

func Boolean(description string) schema.PropertySpec {
	return schema.PropertySpec{
		Description: description,
		TypeSpec:    types.Boolean,
	}
}

func Connection(commandSpec schema.PackageSpec) schema.PropertySpec {
	return schema.PropertySpec{
		Description: "The parameters with which to connect to the remote host.",
		TypeSpec:    types.ExtType(commandSpec, "Connection", "remote"),
	}
}

func Integer(description string) schema.PropertySpec {
	return schema.PropertySpec{
		Description: description,
		TypeSpec:    types.Integer,
	}
}

func OneOrMoreStrings(desctiption string) schema.PropertySpec {
	return schema.PropertySpec{
		Description: desctiption,
		TypeSpec:    types.OneOrMoreStrings,
	}
}

func String(description string) schema.PropertySpec {
	return schema.PropertySpec{
		Description: description,
		TypeSpec:    types.String,
	}
}

func StringMap(description string) schema.PropertySpec {
	return schema.PropertySpec{
		Description: description,
		TypeSpec:    types.StringMap,
	}
}

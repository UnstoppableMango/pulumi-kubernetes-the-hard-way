package rt

import (
	"testing"

	"github.com/pulumi/pulumi/sdk/v3/go/common/apitype"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
)

type ResourceValidator func(t *testing.T, res apitype.ResourceV3)

type validatorKey struct {
	Type tokens.Type
	Name string
}

func Validate(ctx *ResourceContext, typ tokens.Type, name string, validator ResourceValidator) {
	ctx.tokens[validatorKey{Type: typ, Name: name}] = validator
}

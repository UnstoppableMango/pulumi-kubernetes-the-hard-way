package tools

import (
	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/schemagen/pkg/gen/props"
	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/schemagen/pkg/gen/types"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

func generateHostnamectl() schema.ResourceSpec {
	inputs := map[string]schema.PropertySpec{
		"arg": props.String("The argument for the specified `command`."),
		"command": {
			Description: "Corresponds to the {COMMAND} argument.",
			TypeSpec:    types.LocalType("HostnamectlCommand", "tools"),
		},
		"help": props.Boolean("Print a short help text and exit."),
		"host": props.String("Execute the operation remotely. Specify a hostname, or a username and hostname separated by '@', to connect to."),
		"json": {
			Description: "Shows output formatted as JSON.",
			TypeSpec:    types.LocalType("HostnamectlJsonMode", "tools"),
		},
		"machine":       props.String("Execute operation on a local container. Specify a container name to connect to, optionally prefixed by a user name to connect as and a separating '@' character."),
		"noAskPassword": props.Boolean("Do not query the user for authentication for privileged operations."),
		"pretty":        props.Boolean("If status is invoked (or no explicit command is given) and one of these switches is specified, hostnamectl will print out just this selected hostname. Same as `static` and `transient`."),
		"static":        props.Boolean("If status is invoked (or no explicit command is given) and one of these switches is specified, hostnamectl will print out just this selected hostname. Same as `transient` and `pretty`."),
		"transient":     props.Boolean("If status is invoked (or no explicit command is given) and one of these switches is specified, hostnamectl will print out just this selected hostname. Same as `static` and `pretty`."),
		"version":       props.Boolean("Print a short version string and exit."),
	}

	required := []string{"command"}

	return schema.ResourceSpec{
		ObjectTypeSpec: schema.ObjectTypeSpec{
			Description: "Abstraction over the `hostnamectl` utility on a remote system.",
			Properties:  implicitOutputs(inputs, map[string]schema.PropertySpec{}),
			Required: append(required,
				"help",
				"json",
				"noAskPassword",
				"pretty",
				"static",
				"transient",
				"version",
			),
		},
		InputProperties: inputs,
		RequiredInputs:  required,
	}
}

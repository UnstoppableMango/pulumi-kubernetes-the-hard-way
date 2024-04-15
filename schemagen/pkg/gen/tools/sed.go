package tools

import (
	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/schemagen/pkg/gen/props"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

func generateSed() schema.ResourceSpec {
	inputs := map[string]schema.PropertySpec{
		"debug":          props.Boolean("annotate program execution."),
		"expressions":    props.OneOrMoreStrings("add the script to the commands to be executed."),
		"files":          props.OneOrMoreStrings("add the contents of script-file to the commands to be executed."),
		"followSymlinks": props.Boolean("follow symlinks when processing in place"),
		"help":           props.Boolean("display this help and exit."),
		"inPlace":        props.String("edit files in place (makes backup if SUFFIX supplied)"),
		"inputFiles":     props.OneOrMoreStrings("corresponds to the [input-file]... argument(s)."),
		"lineLength":     props.Integer("specify the desired line-wrap length for the `l' command"),
		"nullData":       props.Boolean("separate lines by NUL characters"),
		"posix":          props.Boolean("disable all GNU extensions."),
		"quiet":          props.Boolean("suppress automatic printing of pattern space. Same as `silent`."),
		"regexpExtended": props.Boolean("use extended regular expressions in the script (for portability use POSIX -E)."),
		"script":         props.String("script only if no other script."),
		"sandbox":        props.Boolean("operate in sandbox mode (disable e/r/w commands)."),
		"separate":       props.Boolean("consider files as separate rather than as a single, continuous long stream."),
		"silent":         props.Boolean("suppress automatic printing of pattern space. Same as `quiet`."),
		"unbuffered":     props.Boolean("load minimal amounts of data from the input files and flush the output buffers more often."),
		"version":        props.Boolean("output version information and exit."),
	}

	required := []string{}

	return schema.ResourceSpec{
		ObjectTypeSpec: schema.ObjectTypeSpec{
			Description: "Abstraction over the `sed` utility on a remote system.",
			Properties:  implicitOutputs(inputs, map[string]schema.PropertySpec{}),
			Required: append(required,
				"debug",
				"expressions",
				"files",
				"followSymlinks",
				"help",
				"inputFiles",
				"nullData",
				"posix",
				"quiet",
				"regexpExtended",
				"sandbox",
				"separate",
				"silent",
				"unbuffered",
				"version",
			),
		},
		InputProperties: inputs,
		RequiredInputs:  required,
	}
}

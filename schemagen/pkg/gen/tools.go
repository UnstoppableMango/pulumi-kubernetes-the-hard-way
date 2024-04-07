package gen

import (
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

const toolsMod = "kubernetes-the-hard-way:tools:"

func generateTools(commandSpec schema.PackageSpec) schema.PackageSpec {
	types := map[string]schema.ComplexTypeSpec{
		toolsMod + "SystemctlCommand": {
			ObjectTypeSpec: schema.ObjectTypeSpec{
				Type: "string",
			},
			Enum: []schema.EnumValueSpec{
				{Value: "bind"},
				{Value: "cat"},
				{Value: "clean"},
				{Value: "daemon-reload"},
				{Value: "disable"},
				{Value: "enable"},
				{Value: "freeze"},
				{Value: "is-active"},
				{Value: "is-enabled"},
				{Value: "is-failed"},
				{Value: "isolate"},
				{Value: "kill"},
				{Value: "list-automounts"},
				{Value: "list-dependencies"},
				{Value: "list-paths"},
				{Value: "list-sockets"},
				{Value: "list-timers"},
				{Value: "list-units"},
				{Value: "mask"},
				{Value: "mount-image"},
				{Value: "reenable"},
				{Value: "reload"},
				{Value: "reload-or-restart"},
				{Value: "restart"},
				{Value: "set-property"},
				{Value: "show"},
				{Value: "start"},
				{Value: "status"},
				{Value: "stop"},
				{Value: "thaw"},
				{Value: "try-reload-or-restart"},
				{Value: "try-restart"},
				{Value: "unmask"},
			},
		},
		toolsMod + "TeeMode": {
			ObjectTypeSpec: schema.ObjectTypeSpec{
				Type: "string",
			},
			Enum: []schema.EnumValueSpec{
				{Value: "warn"},
				{Value: "warn-nopipe"},
				{Value: "exit"},
				{Value: "exit-nopipe"},
			},
		},
	}

	commonProps := map[string]schema.PropertySpec{
		"binaryPath": {
			Description: "Path to the binary on the remote system. If omitted, the tool is assumed to be on $PATH",
			TypeSpec: schema.TypeSpec{
				Type: "string",
			},
		},
		"connection": {
			Description: "Connection details for the remote system",
			TypeSpec: schema.TypeSpec{
				Ref: refType(commandSpec, "Connection", "remote"),
			},
		},
		"environment": {
			Description: "Environment variables",
			TypeSpec: schema.TypeSpec{
				Type: "object",
				AdditionalProperties: &schema.TypeSpec{
					Type: "string",
				},
			},
		},
		"lifecycle": {
			Description: "At what stage(s) in the resource lifecycle should the command be run",
			TypeSpec: schema.TypeSpec{
				Ref:   "#/types/kubernetes-the-hard-way:tools:CommandLifecycle",
				Plain: true,
			},
		},
		"stdin": {
			Description: "TODO",
			TypeSpec: schema.TypeSpec{
				Type: "string",
			},
		},
		"triggers": {
			Description: "TODO",
			TypeSpec: schema.TypeSpec{
				Type: "array",
				Items: &schema.TypeSpec{
					Ref: "pulumi.json#/Any",
				},
			},
		},
	}

	commonOutputs := map[string]schema.PropertySpec{
		"command": {
			Description: "The underlying command",
			TypeSpec: schema.TypeSpec{
				Ref: refResource(commandSpec, "Command", "remote"),
			},
		},
		"stderr": {
			Description: "TODO",
			TypeSpec: schema.TypeSpec{
				Type: "string",
			},
		},
		"stdin": {
			Description: "TODO",
			TypeSpec: schema.TypeSpec{
				Type: "string",
			},
		},
		"stdout": {
			Description: "TODO",
			TypeSpec: schema.TypeSpec{
				Type: "string",
			},
		},
	}

	tools := map[string]schema.ResourceSpec{
		"Etcdctl":   generateEtcdctl(),
		"Mkdir":     generateMkdir(),
		"Mktemp":    generateMktemp(),
		"Mv":        generateMv(),
		"Rm":        generateRm(),
		"Systemctl": generateSystemctl(),
		"Tar":       generateTar(),
		"Tee":       generateTee(),
		"Wget":      generateWget(),
	}

	resources := map[string]schema.ResourceSpec{}
	for tool, resource := range tools {
		for propName, prop := range commonProps {
			resource.InputProperties[propName] = prop
			resource.Properties[propName] = prop
		}
		for name, prop := range commonOutputs {
			resource.Properties[name] = prop
		}

		resource.IsComponent = true
		resource.RequiredInputs = append(resource.RequiredInputs,
			"connection",
		)
		resource.Required = append(resource.Required,
			"binaryPath",
			"command",
			"connection",
			"environment",
			"stderr",
			"stdout",
			"triggers",
		)

		resources[toolsMod+tool] = resource
	}

	return schema.PackageSpec{
		Functions: map[string]schema.FunctionSpec{},
		Resources: resources,
		Types:     types,
	}
}

func generateEtcdctl() schema.ResourceSpec {
	inputs := map[string]schema.PropertySpec{
		"caCert": {
			Description: "TODO",
			TypeSpec: schema.TypeSpec{
				Type: "string",
			},
		},
		"cert": {
			Description: "TODO",
			TypeSpec: schema.TypeSpec{
				Type: "string",
			},
		},
		"commands": {
			Description: "TODO",
			TypeSpec: schema.TypeSpec{
				Ref: "#/types/kubernetes-the-hard-way:tools:EtcdctlCommand",
			},
		},
		"endpoints": {
			Description: "TODO",
			TypeSpec: schema.TypeSpec{
				Type: "string",
			},
		},
		"key": {
			Description: "TODO",
			TypeSpec: schema.TypeSpec{
				Type: "string",
			},
		},
	}

	required := []string{
		"caCert",
		"cert",
		"commands",
		"endpoints",
		"key",
	}

	return schema.ResourceSpec{
		ObjectTypeSpec: schema.ObjectTypeSpec{
			Description: "Abstraction over the `etcdctl` utility on a remote system.",
			Properties:  implicitOutputs(inputs, map[string]schema.PropertySpec{}),
			Required:    required,
		},
		InputProperties: inputs,
		RequiredInputs:  required,
	}
}

func generateMkdir() schema.ResourceSpec {
	inputs := map[string]schema.PropertySpec{
		"directory": {
			Description: "The fully qualified path of the directory on the remote system.",
			TypeSpec: schema.TypeSpec{
				Type: "string",
			},
		},
		"parents": {
			Description: "Corresponds to the `--parents` option.",
			TypeSpec: schema.TypeSpec{
				Type: "boolean",
			},
		},
		"removeOnDelete": { // TODO: Reconsider this SOLID violation
			Description: "Remove the created directory when the `Mkdir` resource is deleted or updated.",
			TypeSpec: schema.TypeSpec{
				Type: "boolean",
			},
		},
	}

	required := []string{
		"directory",
	}

	return schema.ResourceSpec{
		ObjectTypeSpec: schema.ObjectTypeSpec{
			Description: "Abstraction over the `mkdir` utility on a remote system.",
			Properties:  implicitOutputs(inputs, map[string]schema.PropertySpec{}),
			Required: append(required,
				"parents",
				"removeOnDelete"),
		},
		InputProperties: inputs,
		RequiredInputs:  required,
	}
}

func generateMktemp() schema.ResourceSpec {
	inputs := map[string]schema.PropertySpec{
		"directory": {
			Description: "Corresponds to the `--directory` option.",
			TypeSpec: schema.TypeSpec{
				Type: "boolean",
			},
		},
		"dryRun": {
			Description: "Corresponds to the `--dry-run` option.",
			TypeSpec: schema.TypeSpec{
				Type: "boolean",
			},
		},
		"quiet": {
			Description: "Corresponds to the `--quiet` option.",
			TypeSpec: schema.TypeSpec{
				Type: "boolean",
			},
		},
		"suffix": {
			Description: "Corresponds to the `--suffix` option.",
			TypeSpec: schema.TypeSpec{
				Type: "string",
			},
		},
		"template": {
			Description: "Corresponds to the [TEMPLATE] argument.",
			TypeSpec: schema.TypeSpec{
				Type: "boolean",
			},
		},
		"tmpdir": {
			Description: "Corresponds to the `--tmpdir` option.",
			TypeSpec: schema.TypeSpec{
				Type: "string",
			},
		},
	}

	return schema.ResourceSpec{
		ObjectTypeSpec: schema.ObjectTypeSpec{
			Description: "Abstraction over the `mkdir` utility on a remote system.",
			Properties:  implicitOutputs(inputs, map[string]schema.PropertySpec{}),
			Required: []string{
				"dryRun",
				"quiet",
			},
		},
		InputProperties: inputs,
		RequiredInputs:  []string{},
	}
}

func generateMv() schema.ResourceSpec {
	inputs := map[string]schema.PropertySpec{
		"backup": {
			Description: "Corresponds to the `-b` and `--backup` options depending on whether [CONTROL] is supplied.",
			TypeSpec: schema.TypeSpec{
				Type:  "boolean",
				Plain: true,
			},
		},
		"context": {
			Description: "Corresponds to the `--context` option.",
			TypeSpec: schema.TypeSpec{
				Type: "boolean",
			},
		},
		"control": {
			Description: "Corresponds to the [CONTROL] argument for the `--backup` option.",
			TypeSpec: schema.TypeSpec{
				Type: "boolean",
			},
		},
		"dest": {
			Description: "Corresponds to the [DEST] argument.",
			TypeSpec: schema.TypeSpec{
				Type: "string",
			},
		},
		"directory": {
			Description: "Corresponds to the [DIRECTORY] argument.",
			TypeSpec: schema.TypeSpec{
				Type: "string",
			},
		},
		"force": {
			Description: "Corresponds to the `--force` option.",
			TypeSpec: schema.TypeSpec{
				Type: "boolean",
			},
		},
		"noClobber": {
			Description: "Corresponds to the `--no-clobber` option.",
			TypeSpec: schema.TypeSpec{
				Type: "boolean",
			},
		},
		"noTargetDirectory": {
			Description: "Corresponds to the `--no-target-directory` option.",
			TypeSpec: schema.TypeSpec{
				Type: "boolean",
			},
		},
		"source": {
			Description: "Corresponds to the [SOURCE] argument.",
			TypeSpec: schema.TypeSpec{
				OneOf: []schema.TypeSpec{
					{Type: "string"},
					arrayOfStrings,
				},
			},
		},
		"stripTrailingSlashes": {
			Description: "Corresponds to the `--strip-trailing-slashes` option.",
			TypeSpec: schema.TypeSpec{
				Type: "boolean",
			},
		},
		"suffix": {
			Description: "Corresponds to the `--suffix` option.",
			TypeSpec: schema.TypeSpec{
				Type: "string",
			},
		},
		"targetDirectory": {
			Description: "Corresponds to the `--target-directory` option.",
			TypeSpec: schema.TypeSpec{
				Type: "boolean",
			},
		},
		"update": {
			Description: "Corresponds to the `--update` option.",
			TypeSpec: schema.TypeSpec{
				Type: "boolean",
			},
		},
		"verbose": {
			Description: "Corresponds to the `--verbose` option.",
			TypeSpec: schema.TypeSpec{
				Type: "boolean",
			},
		},
	}

	required := []string{
		"source",
	}

	return schema.ResourceSpec{
		ObjectTypeSpec: schema.ObjectTypeSpec{
			Description: "Abstraction over the `mv` utility on a remote system.",
			Properties:  implicitOutputs(inputs, map[string]schema.PropertySpec{}),
			Required: append(required,
				"backup",
				"context",
				"force",
				"noClobber",
				"noTargetDirectory",
				"source",
				"stripTrailingSlashes",
				"update",
				"verbose",
			),
		},
		InputProperties: inputs,
		RequiredInputs:  required,
	}
}

func generateRm() schema.ResourceSpec {
	inputs := map[string]schema.PropertySpec{
		"dir": {
			Description: "Corresponds to the `--dir` option.",
			TypeSpec: schema.TypeSpec{
				Type:  "boolean",
				Plain: true,
			},
		},
		"files": {
			Description: "Corresponds to the [FILE] argument.",
			TypeSpec: schema.TypeSpec{
				OneOf: []schema.TypeSpec{
					{Type: "string"},
					arrayOfStrings,
				},
			},
		},
		"force": {
			Description: "Corresponds to the `--force` option.",
			TypeSpec: schema.TypeSpec{
				Type: "boolean",
			},
		},
		"onDelete": { // TODO: Reconsider this SOLID violation
			Description: "Whether rm should be run when the resource is created or deleted.",
		},
		"recursive": {
			Description: "Corresponds to the `--recursive` option.",
			TypeSpec: schema.TypeSpec{
				Type: "boolean",
			},
		},
		"verbose": {
			Description: "Corresponds to the `--verbose` option.",
			TypeSpec: schema.TypeSpec{
				Type: "boolean",
			},
		},
	}

	required := []string{
		"files",
	}

	return schema.ResourceSpec{
		ObjectTypeSpec: schema.ObjectTypeSpec{
			Description: "Abstraction over the `rm` utility on a remote system.",
			Properties:  implicitOutputs(inputs, map[string]schema.PropertySpec{}),
			Required: append(required,
				"dir",
				"force",
				"onDelete",
				"recursive",
				"verbose",
			),
		},
		InputProperties: inputs,
		RequiredInputs:  required,
	}
}

func generateSystemctl() schema.ResourceSpec {
	inputs := map[string]schema.PropertySpec{
		"command": { // The common output will overwrite this, that's fine.
			Description: "Corresponds to the COMMAND argument.",
			TypeSpec: schema.TypeSpec{
				Ref:   localType("SystemctlCommand", "tools"),
				Plain: true,
			},
		},
		"pattern": {
			Description: "Corresponds to the [PATTERN] argument",
			TypeSpec: schema.TypeSpec{
				Type: "string",
			},
		},
		"unit": {
			Description: "Corresponds to the [UNIT...] argument.",
			TypeSpec: schema.TypeSpec{
				OneOf: []schema.TypeSpec{
					{Type: "string"},
					arrayOfStrings,
				},
			},
		},
	}

	required := []string{
		"command",
		"unit",
	}

	return schema.ResourceSpec{
		ObjectTypeSpec: schema.ObjectTypeSpec{
			Description: "Abstraction over the `systemctl` utility on a remote system.",
			Properties: implicitOutputs(inputs, map[string]schema.PropertySpec{
				"systemctlCommand": inputs["command"],
			}),
			Required: []string{
				"systemctlCommand",
				"unit",
			},
		},
		InputProperties: inputs,
		RequiredInputs:  required,
	}
}

func generateTar() schema.ResourceSpec {
	inputs := map[string]schema.PropertySpec{
		"archive": {
			Description: "Corresponds to the [ARCHIVE] argument.",
			TypeSpec: schema.TypeSpec{
				Type: "string",
			},
		},
		"directory": {
			Description: "Corresponds to the `--directory` option.",
			TypeSpec: schema.TypeSpec{
				Type: "string",
			},
		},
		"extract": {
			Description: "Corresponds to the `--extract` option.",
			TypeSpec: schema.TypeSpec{
				Type: "boolean",
			},
		},
		"files": {
			Description: "Corresponds to the [FILE] argument.",
			TypeSpec: schema.TypeSpec{
				OneOf: []schema.TypeSpec{
					{Type: "string"},
					arrayOfStrings,
				},
			},
		},
		"gzip": {
			Description: "Corresponds to the `--gzip` option.",
			TypeSpec: schema.TypeSpec{
				Type: "boolean",
			},
		},
		"onDelete": { // TODO: Reconsider this SOLID violation
			Description: "Whether rm should be run when the resource is created or deleted.",
		},
		"recursive": {
			Description: "Corresponds to the `--strip-components` option.",
			TypeSpec: schema.TypeSpec{
				Type: "integer",
			},
		},
	}

	required := []string{
		"archive",
	}

	return schema.ResourceSpec{
		ObjectTypeSpec: schema.ObjectTypeSpec{
			Description: "Abstraction over the `rm` utility on a remote system.",
			Properties:  implicitOutputs(inputs, map[string]schema.PropertySpec{}),
			Required: append(required,
				"extract",
				"files",
			),
		},
		InputProperties: inputs,
		RequiredInputs:  required,
	}
}

func generateTee() schema.ResourceSpec {
	inputs := map[string]schema.PropertySpec{
		"append": {
			Description: "Append to the given FILEs, do not overwrite",
			TypeSpec: schema.TypeSpec{
				Type: "boolean",
			},
		},
		"files": {
			Description: "Corresponds to the [FILE] argument.",
			TypeSpec: schema.TypeSpec{
				OneOf: []schema.TypeSpec{
					{Type: "string"},
					arrayOfStrings,
				},
			},
		},
		"ignoreInterrupts": {
			Description: "Ignore interrupt signals.",
			TypeSpec: schema.TypeSpec{
				Type: "boolean",
			},
		},
		"outputError": {
			Description: "Set behavior on write error.",
			TypeSpec: schema.TypeSpec{
				Ref: localType("TeeMode", "tools"),
			},
		},
		"pipe": {
			Description: "Operate in a more appropriate MODE with pipes.",
			TypeSpec: schema.TypeSpec{
				Type: "boolean",
			},
		},
		"version": {
			Description: "Output version information and exit.",
			TypeSpec: schema.TypeSpec{
				Type: "boolean",
			},
		},
	}

	required := []string{
		"files",
		"stdin",
	}

	return schema.ResourceSpec{
		ObjectTypeSpec: schema.ObjectTypeSpec{
			Description: "Abstraction over the `rm` utility on a remote system.",
			Properties:  implicitOutputs(inputs, map[string]schema.PropertySpec{}),
			Required: append(required,
				"append",
				"ignoreInterrupts",
				"pipe",
				"version",
			),
		},
		InputProperties: inputs,
		RequiredInputs:  required,
	}
}

func generateWget() schema.ResourceSpec {
	inputs := map[string]schema.PropertySpec{
		"directoryPrefix": {
			Description: "The  directory prefix is the directory where all other files and subdirectories will be saved to, i.e. the top of the retrieval tree.  The default is . (the current directory).",
			TypeSpec: schema.TypeSpec{
				Type: "string",
			},
		},
		"httpsOnly": {
			Description: "When in recursive mode, only HTTPS links are followed.",
			TypeSpec: schema.TypeSpec{
				Type: "boolean",
			},
		},
		"noVerbose": {
			Description: "Turn off verbose without being completely quiet (use -q for that), which means that error messages and basic information still get printed.",
			TypeSpec: schema.TypeSpec{
				Type: "boolean",
			},
		},
		"outputDocument": {
			Description: "The  documents  will  not  be  written  to the appropriate files, but all will be concatenated together and written to file.",
			TypeSpec: schema.TypeSpec{
				Type: "string",
			},
		},
		"quiet": {
			Description: "Turn off Wget's output.",
			TypeSpec: schema.TypeSpec{
				Type: "boolean",
			},
		},
		"timestamping": {
			Description: "Turn on time-stamping.",
			TypeSpec: schema.TypeSpec{
				Type: "boolean",
			},
		},
		"url": {
			Description: "Corresponds to the [URL...] argument.",
			TypeSpec: schema.TypeSpec{
				OneOf: []schema.TypeSpec{
					{Type: "string"},
					arrayOfStrings,
				},
			},
		},
	}

	required := []string{
		"url",
	}

	return schema.ResourceSpec{
		ObjectTypeSpec: schema.ObjectTypeSpec{
			Description: "Abstraction over the `wget` utility on a remote system.",
			Properties:  implicitOutputs(inputs, map[string]schema.PropertySpec{}),
			Required: append(required,
				"httpsOnly",
				"quiet",
				"noVerbose",
				"timestamping",
			),
		},
		InputProperties: inputs,
		RequiredInputs:  required,
	}
}

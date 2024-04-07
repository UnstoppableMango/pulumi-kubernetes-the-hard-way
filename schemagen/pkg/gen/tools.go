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
			TypeSpec:    typeSpecs.String,
		},
		"connection": {
			Description: "Connection details for the remote system",
			TypeSpec: schema.TypeSpec{
				Ref: refType(commandSpec, "Connection", "remote"),
			},
		},
		"environment": {
			Description: "Environment variables",
			TypeSpec:    typeSpecs.StringMap,
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
			TypeSpec:    typeSpecs.String,
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
			TypeSpec:    typeSpecs.String,
		},
		"stdin": {
			Description: "TODO",
			TypeSpec:    typeSpecs.String,
		},
		"stdout": {
			Description: "TODO",
			TypeSpec:    typeSpecs.String,
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
			TypeSpec:    typeSpecs.String,
		},
		"cert": {
			Description: "TODO",
			TypeSpec:    typeSpecs.String,
		},
		"commands": {
			Description: "TODO",
			TypeSpec: schema.TypeSpec{
				Ref: localType("EtcdctlCommand", "tools"),
			},
		},
		"endpoints": {
			Description: "TODO",
			TypeSpec:    typeSpecs.String,
		},
		"key": {
			Description: "TODO",
			TypeSpec:    typeSpecs.String,
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
			TypeSpec:    typeSpecs.String,
		},
		"parents": {
			Description: "Corresponds to the `--parents` option.",
			TypeSpec:    typeSpecs.Boolean,
		},
		"removeOnDelete": { // TODO: Reconsider this SOLID violation
			Description: "Remove the created directory when the `Mkdir` resource is deleted or updated.",
			TypeSpec:    typeSpecs.Boolean,
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
			TypeSpec:    typeSpecs.Boolean,
		},
		"dryRun": {
			Description: "Corresponds to the `--dry-run` option.",
			TypeSpec:    typeSpecs.Boolean,
		},
		"quiet": {
			Description: "Corresponds to the `--quiet` option.",
			TypeSpec:    typeSpecs.Boolean,
		},
		"suffix": {
			Description: "Corresponds to the `--suffix` option.",
			TypeSpec:    typeSpecs.String,
		},
		"template": {
			Description: "Corresponds to the [TEMPLATE] argument.",
			TypeSpec:    typeSpecs.Boolean,
		},
		"tmpdir": {
			Description: "Corresponds to the `--tmpdir` option.",
			TypeSpec:    typeSpecs.String,
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
			TypeSpec:    typeSpecs.Boolean,
		},
		"control": {
			Description: "Corresponds to the [CONTROL] argument for the `--backup` option.",
			TypeSpec:    typeSpecs.Boolean,
		},
		"dest": {
			Description: "Corresponds to the [DEST] argument.",
			TypeSpec:    typeSpecs.String,
		},
		"directory": {
			Description: "Corresponds to the [DIRECTORY] argument.",
			TypeSpec:    typeSpecs.String,
		},
		"force": {
			Description: "Corresponds to the `--force` option.",
			TypeSpec:    typeSpecs.Boolean,
		},
		"noClobber": {
			Description: "Corresponds to the `--no-clobber` option.",
			TypeSpec:    typeSpecs.Boolean,
		},
		"noTargetDirectory": {
			Description: "Corresponds to the `--no-target-directory` option.",
			TypeSpec:    typeSpecs.Boolean,
		},
		"source": {
			Description: "Corresponds to the [SOURCE] argument.",
			TypeSpec:    typeSpecs.OneOrMoreStrings,
		},
		"stripTrailingSlashes": {
			Description: "Corresponds to the `--strip-trailing-slashes` option.",
			TypeSpec:    typeSpecs.Boolean,
		},
		"suffix": {
			Description: "Corresponds to the `--suffix` option.",
			TypeSpec:    typeSpecs.String,
		},
		"targetDirectory": {
			Description: "Corresponds to the `--target-directory` option.",
			TypeSpec:    typeSpecs.Boolean,
		},
		"update": {
			Description: "Corresponds to the `--update` option.",
			TypeSpec:    typeSpecs.Boolean,
		},
		"verbose": {
			Description: "Corresponds to the `--verbose` option.",
			TypeSpec:    typeSpecs.Boolean,
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
			TypeSpec:    typeSpecs.OneOrMoreStrings,
		},
		"force": {
			Description: "Corresponds to the `--force` option.",
			TypeSpec:    typeSpecs.Boolean,
		},
		"onDelete": { // TODO: Reconsider this SOLID violation
			Description: "Whether rm should be run when the resource is created or deleted.",
		},
		"recursive": {
			Description: "Corresponds to the `--recursive` option.",
			TypeSpec:    typeSpecs.Boolean,
		},
		"verbose": {
			Description: "Corresponds to the `--verbose` option.",
			TypeSpec:    typeSpecs.Boolean,
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
			TypeSpec:    typeSpecs.String,
		},
		"unit": {
			Description: "Corresponds to the [UNIT...] argument.",
			TypeSpec:    typeSpecs.OneOrMoreStrings,
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
			TypeSpec:    typeSpecs.String,
		},
		"directory": {
			Description: "Corresponds to the `--directory` option.",
			TypeSpec:    typeSpecs.String,
		},
		"extract": {
			Description: "Corresponds to the `--extract` option.",
			TypeSpec:    typeSpecs.Boolean,
		},
		"files": {
			Description: "Corresponds to the [FILE] argument.",
			TypeSpec:    typeSpecs.OneOrMoreStrings,
		},
		"gzip": {
			Description: "Corresponds to the `--gzip` option.",
			TypeSpec:    typeSpecs.Boolean,
		},
		"onDelete": { // TODO: Reconsider this SOLID violation
			Description: "Whether rm should be run when the resource is created or deleted.",
		},
		"recursive": {
			Description: "Corresponds to the `--strip-components` option.",
			TypeSpec:    typeSpecs.Integer,
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
			TypeSpec:    typeSpecs.Boolean,
		},
		"files": {
			Description: "Corresponds to the [FILE] argument.",
			TypeSpec:    typeSpecs.OneOrMoreStrings,
		},
		"ignoreInterrupts": {
			Description: "Ignore interrupt signals.",
			TypeSpec:    typeSpecs.Boolean,
		},
		"outputError": {
			Description: "Set behavior on write error.",
			TypeSpec: schema.TypeSpec{
				Ref: localType("TeeMode", "tools"),
			},
		},
		"pipe": {
			Description: "Operate in a more appropriate MODE with pipes.",
			TypeSpec:    typeSpecs.Boolean,
		},
		"version": {
			Description: "Output version information and exit.",
			TypeSpec:    typeSpecs.Boolean,
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
			TypeSpec:    typeSpecs.String,
		},
		"httpsOnly": {
			Description: "When in recursive mode, only HTTPS links are followed.",
			TypeSpec:    typeSpecs.Boolean,
		},
		"noVerbose": {
			Description: "Turn off verbose without being completely quiet (use -q for that), which means that error messages and basic information still get printed.",
			TypeSpec:    typeSpecs.Boolean,
		},
		"outputDocument": {
			Description: "The  documents  will  not  be  written  to the appropriate files, but all will be concatenated together and written to file.",
			TypeSpec:    typeSpecs.String,
		},
		"quiet": {
			Description: "Turn off Wget's output.",
			TypeSpec:    typeSpecs.Boolean,
		},
		"timestamping": {
			Description: "Turn on time-stamping.",
			TypeSpec:    typeSpecs.Boolean,
		},
		"url": {
			Description: "Corresponds to the [URL...] argument.",
			TypeSpec:    typeSpecs.OneOrMoreStrings,
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

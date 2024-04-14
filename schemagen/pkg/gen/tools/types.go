package tools

import "github.com/pulumi/pulumi/pkg/v3/codegen/schema"

func generateTypes() map[string]schema.ComplexTypeSpec {
	return map[string]schema.ComplexTypeSpec{
		name("CommandLifecycle"): {
			ObjectTypeSpec: schema.ObjectTypeSpec{
				Type: "string",
			},
			Enum: []schema.EnumValueSpec{
				{Value: "create"},
				{Value: "update"},
				{Value: "delete"},
			},
		},
		name("EtcdctlCommand"): {
			ObjectTypeSpec: schema.ObjectTypeSpec{
				Type: "string",
			},
			Enum: []schema.EnumValueSpec{
				{Value: "member"},
				{Value: "list"},
			},
		},
		name("HostnamectlCommand"): {
			ObjectTypeSpec: schema.ObjectTypeSpec{
				Type: "string",
			},
			Enum: []schema.EnumValueSpec{
				{Value: "status", Description: "Show system hostname and related information. If no command is specified, this is the implied default."},
				{Value: "hostname", Description: "If no argument is given, print the system hostname. If an optional argument NAME is provided then the command changes the system hostname to NAME."},
				{Value: "icon-name", Description: "If no argument is given, print the icon name of the system. If an optional argument NAME is provided then the command changes the icon name to NAME."},
				{Value: "chassis", Description: "If no argument is given, print the chassis type. If an optional argument TYPE is provided then the command changes the chassis type to TYPE."},
				{Value: "deployment", Description: "If no argument is given, print the deployment environment. If an optional argument ENVIRONMENT is provided then the command changes the deployment environment to ENVIRONMENT."},
				{Value: "location", Description: "If no argument is given, print the location string for the system. If an optional argument LOCATION is provided then the command changes the location string for the system to LOCATION."},
			},
		},
		name("HostnamectlJsonMode"): {
			ObjectTypeSpec: schema.ObjectTypeSpec{
				Type: "string",
			},
			Enum: []schema.EnumValueSpec{
				{Value: "short"},
				{Value: "pretty"},
				{Value: "off"},
			},
		},
		name("SystemctlCommand"): {
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
		name("TeeMode"): {
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
}

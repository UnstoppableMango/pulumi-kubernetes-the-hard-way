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

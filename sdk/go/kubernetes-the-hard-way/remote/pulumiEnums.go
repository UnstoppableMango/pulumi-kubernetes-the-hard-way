// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package remote

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// CPU architecture
type Architecture string

const (
	ArchitectureAmd64 = Architecture("amd64")
	ArchitectureArm64 = Architecture("arm64")
)

func (Architecture) ElementType() reflect.Type {
	return reflect.TypeOf((*Architecture)(nil)).Elem()
}

func (e Architecture) ToArchitectureOutput() ArchitectureOutput {
	return pulumi.ToOutput(e).(ArchitectureOutput)
}

func (e Architecture) ToArchitectureOutputWithContext(ctx context.Context) ArchitectureOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ArchitectureOutput)
}

func (e Architecture) ToArchitecturePtrOutput() ArchitecturePtrOutput {
	return e.ToArchitecturePtrOutputWithContext(context.Background())
}

func (e Architecture) ToArchitecturePtrOutputWithContext(ctx context.Context) ArchitecturePtrOutput {
	return Architecture(e).ToArchitectureOutputWithContext(ctx).ToArchitecturePtrOutputWithContext(ctx)
}

func (e Architecture) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e Architecture) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e Architecture) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e Architecture) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ArchitectureOutput struct{ *pulumi.OutputState }

func (ArchitectureOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Architecture)(nil)).Elem()
}

func (o ArchitectureOutput) ToArchitectureOutput() ArchitectureOutput {
	return o
}

func (o ArchitectureOutput) ToArchitectureOutputWithContext(ctx context.Context) ArchitectureOutput {
	return o
}

func (o ArchitectureOutput) ToArchitecturePtrOutput() ArchitecturePtrOutput {
	return o.ToArchitecturePtrOutputWithContext(context.Background())
}

func (o ArchitectureOutput) ToArchitecturePtrOutputWithContext(ctx context.Context) ArchitecturePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Architecture) *Architecture {
		return &v
	}).(ArchitecturePtrOutput)
}

func (o ArchitectureOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ArchitectureOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Architecture) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ArchitectureOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ArchitectureOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Architecture) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ArchitecturePtrOutput struct{ *pulumi.OutputState }

func (ArchitecturePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Architecture)(nil)).Elem()
}

func (o ArchitecturePtrOutput) ToArchitecturePtrOutput() ArchitecturePtrOutput {
	return o
}

func (o ArchitecturePtrOutput) ToArchitecturePtrOutputWithContext(ctx context.Context) ArchitecturePtrOutput {
	return o
}

func (o ArchitecturePtrOutput) Elem() ArchitectureOutput {
	return o.ApplyT(func(v *Architecture) Architecture {
		if v != nil {
			return *v
		}
		var ret Architecture
		return ret
	}).(ArchitectureOutput)
}

func (o ArchitecturePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ArchitecturePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *Architecture) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// ArchitectureInput is an input type that accepts values of the Architecture enum
// A concrete instance of `ArchitectureInput` can be one of the following:
//
//	ArchitectureAmd64
//	ArchitectureArm64
type ArchitectureInput interface {
	pulumi.Input

	ToArchitectureOutput() ArchitectureOutput
	ToArchitectureOutputWithContext(context.Context) ArchitectureOutput
}

var architecturePtrType = reflect.TypeOf((**Architecture)(nil)).Elem()

type ArchitecturePtrInput interface {
	pulumi.Input

	ToArchitecturePtrOutput() ArchitecturePtrOutput
	ToArchitecturePtrOutputWithContext(context.Context) ArchitecturePtrOutput
}

type architecturePtr string

func ArchitecturePtr(v string) ArchitecturePtrInput {
	return (*architecturePtr)(&v)
}

func (*architecturePtr) ElementType() reflect.Type {
	return architecturePtrType
}

func (in *architecturePtr) ToArchitecturePtrOutput() ArchitecturePtrOutput {
	return pulumi.ToOutput(in).(ArchitecturePtrOutput)
}

func (in *architecturePtr) ToArchitecturePtrOutputWithContext(ctx context.Context) ArchitecturePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ArchitecturePtrOutput)
}

// Systemd service exit type.
type SystemdServiceExitType string

const (
	SystemdServiceExitTypeMain   = SystemdServiceExitType("main")
	SystemdServiceExitTypeCgroup = SystemdServiceExitType("cgroup")
)

func (SystemdServiceExitType) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemdServiceExitType)(nil)).Elem()
}

func (e SystemdServiceExitType) ToSystemdServiceExitTypeOutput() SystemdServiceExitTypeOutput {
	return pulumi.ToOutput(e).(SystemdServiceExitTypeOutput)
}

func (e SystemdServiceExitType) ToSystemdServiceExitTypeOutputWithContext(ctx context.Context) SystemdServiceExitTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SystemdServiceExitTypeOutput)
}

func (e SystemdServiceExitType) ToSystemdServiceExitTypePtrOutput() SystemdServiceExitTypePtrOutput {
	return e.ToSystemdServiceExitTypePtrOutputWithContext(context.Background())
}

func (e SystemdServiceExitType) ToSystemdServiceExitTypePtrOutputWithContext(ctx context.Context) SystemdServiceExitTypePtrOutput {
	return SystemdServiceExitType(e).ToSystemdServiceExitTypeOutputWithContext(ctx).ToSystemdServiceExitTypePtrOutputWithContext(ctx)
}

func (e SystemdServiceExitType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SystemdServiceExitType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SystemdServiceExitType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SystemdServiceExitType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SystemdServiceExitTypeOutput struct{ *pulumi.OutputState }

func (SystemdServiceExitTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemdServiceExitType)(nil)).Elem()
}

func (o SystemdServiceExitTypeOutput) ToSystemdServiceExitTypeOutput() SystemdServiceExitTypeOutput {
	return o
}

func (o SystemdServiceExitTypeOutput) ToSystemdServiceExitTypeOutputWithContext(ctx context.Context) SystemdServiceExitTypeOutput {
	return o
}

func (o SystemdServiceExitTypeOutput) ToSystemdServiceExitTypePtrOutput() SystemdServiceExitTypePtrOutput {
	return o.ToSystemdServiceExitTypePtrOutputWithContext(context.Background())
}

func (o SystemdServiceExitTypeOutput) ToSystemdServiceExitTypePtrOutputWithContext(ctx context.Context) SystemdServiceExitTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SystemdServiceExitType) *SystemdServiceExitType {
		return &v
	}).(SystemdServiceExitTypePtrOutput)
}

func (o SystemdServiceExitTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SystemdServiceExitTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SystemdServiceExitType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SystemdServiceExitTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SystemdServiceExitTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SystemdServiceExitType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SystemdServiceExitTypePtrOutput struct{ *pulumi.OutputState }

func (SystemdServiceExitTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemdServiceExitType)(nil)).Elem()
}

func (o SystemdServiceExitTypePtrOutput) ToSystemdServiceExitTypePtrOutput() SystemdServiceExitTypePtrOutput {
	return o
}

func (o SystemdServiceExitTypePtrOutput) ToSystemdServiceExitTypePtrOutputWithContext(ctx context.Context) SystemdServiceExitTypePtrOutput {
	return o
}

func (o SystemdServiceExitTypePtrOutput) Elem() SystemdServiceExitTypeOutput {
	return o.ApplyT(func(v *SystemdServiceExitType) SystemdServiceExitType {
		if v != nil {
			return *v
		}
		var ret SystemdServiceExitType
		return ret
	}).(SystemdServiceExitTypeOutput)
}

func (o SystemdServiceExitTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SystemdServiceExitTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SystemdServiceExitType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// SystemdServiceExitTypeInput is an input type that accepts values of the SystemdServiceExitType enum
// A concrete instance of `SystemdServiceExitTypeInput` can be one of the following:
//
//	SystemdServiceExitTypeMain
//	SystemdServiceExitTypeCgroup
type SystemdServiceExitTypeInput interface {
	pulumi.Input

	ToSystemdServiceExitTypeOutput() SystemdServiceExitTypeOutput
	ToSystemdServiceExitTypeOutputWithContext(context.Context) SystemdServiceExitTypeOutput
}

var systemdServiceExitTypePtrType = reflect.TypeOf((**SystemdServiceExitType)(nil)).Elem()

type SystemdServiceExitTypePtrInput interface {
	pulumi.Input

	ToSystemdServiceExitTypePtrOutput() SystemdServiceExitTypePtrOutput
	ToSystemdServiceExitTypePtrOutputWithContext(context.Context) SystemdServiceExitTypePtrOutput
}

type systemdServiceExitTypePtr string

func SystemdServiceExitTypePtr(v string) SystemdServiceExitTypePtrInput {
	return (*systemdServiceExitTypePtr)(&v)
}

func (*systemdServiceExitTypePtr) ElementType() reflect.Type {
	return systemdServiceExitTypePtrType
}

func (in *systemdServiceExitTypePtr) ToSystemdServiceExitTypePtrOutput() SystemdServiceExitTypePtrOutput {
	return pulumi.ToOutput(in).(SystemdServiceExitTypePtrOutput)
}

func (in *systemdServiceExitTypePtr) ToSystemdServiceExitTypePtrOutputWithContext(ctx context.Context) SystemdServiceExitTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SystemdServiceExitTypePtrOutput)
}

// Systemd service restart behavior.
type SystemdServiceRestart string

const (
	SystemdServiceRestartNo           = SystemdServiceRestart("no")
	SystemdServiceRestart_On_Success  = SystemdServiceRestart("on-success")
	SystemdServiceRestart_On_Failure  = SystemdServiceRestart("on-failure")
	SystemdServiceRestart_On_Abnormal = SystemdServiceRestart("on-abnormal")
	SystemdServiceRestart_On_Watchdog = SystemdServiceRestart("on-watchdog")
	SystemdServiceRestart_On_Abort    = SystemdServiceRestart("on-abort")
	SystemdServiceRestartAlways       = SystemdServiceRestart("always")
)

func (SystemdServiceRestart) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemdServiceRestart)(nil)).Elem()
}

func (e SystemdServiceRestart) ToSystemdServiceRestartOutput() SystemdServiceRestartOutput {
	return pulumi.ToOutput(e).(SystemdServiceRestartOutput)
}

func (e SystemdServiceRestart) ToSystemdServiceRestartOutputWithContext(ctx context.Context) SystemdServiceRestartOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SystemdServiceRestartOutput)
}

func (e SystemdServiceRestart) ToSystemdServiceRestartPtrOutput() SystemdServiceRestartPtrOutput {
	return e.ToSystemdServiceRestartPtrOutputWithContext(context.Background())
}

func (e SystemdServiceRestart) ToSystemdServiceRestartPtrOutputWithContext(ctx context.Context) SystemdServiceRestartPtrOutput {
	return SystemdServiceRestart(e).ToSystemdServiceRestartOutputWithContext(ctx).ToSystemdServiceRestartPtrOutputWithContext(ctx)
}

func (e SystemdServiceRestart) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SystemdServiceRestart) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SystemdServiceRestart) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SystemdServiceRestart) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SystemdServiceRestartOutput struct{ *pulumi.OutputState }

func (SystemdServiceRestartOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemdServiceRestart)(nil)).Elem()
}

func (o SystemdServiceRestartOutput) ToSystemdServiceRestartOutput() SystemdServiceRestartOutput {
	return o
}

func (o SystemdServiceRestartOutput) ToSystemdServiceRestartOutputWithContext(ctx context.Context) SystemdServiceRestartOutput {
	return o
}

func (o SystemdServiceRestartOutput) ToSystemdServiceRestartPtrOutput() SystemdServiceRestartPtrOutput {
	return o.ToSystemdServiceRestartPtrOutputWithContext(context.Background())
}

func (o SystemdServiceRestartOutput) ToSystemdServiceRestartPtrOutputWithContext(ctx context.Context) SystemdServiceRestartPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SystemdServiceRestart) *SystemdServiceRestart {
		return &v
	}).(SystemdServiceRestartPtrOutput)
}

func (o SystemdServiceRestartOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SystemdServiceRestartOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SystemdServiceRestart) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SystemdServiceRestartOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SystemdServiceRestartOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SystemdServiceRestart) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SystemdServiceRestartPtrOutput struct{ *pulumi.OutputState }

func (SystemdServiceRestartPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemdServiceRestart)(nil)).Elem()
}

func (o SystemdServiceRestartPtrOutput) ToSystemdServiceRestartPtrOutput() SystemdServiceRestartPtrOutput {
	return o
}

func (o SystemdServiceRestartPtrOutput) ToSystemdServiceRestartPtrOutputWithContext(ctx context.Context) SystemdServiceRestartPtrOutput {
	return o
}

func (o SystemdServiceRestartPtrOutput) Elem() SystemdServiceRestartOutput {
	return o.ApplyT(func(v *SystemdServiceRestart) SystemdServiceRestart {
		if v != nil {
			return *v
		}
		var ret SystemdServiceRestart
		return ret
	}).(SystemdServiceRestartOutput)
}

func (o SystemdServiceRestartPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SystemdServiceRestartPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SystemdServiceRestart) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// SystemdServiceRestartInput is an input type that accepts values of the SystemdServiceRestart enum
// A concrete instance of `SystemdServiceRestartInput` can be one of the following:
//
//	SystemdServiceRestartNo
//	SystemdServiceRestart_On_Success
//	SystemdServiceRestart_On_Failure
//	SystemdServiceRestart_On_Abnormal
//	SystemdServiceRestart_On_Watchdog
//	SystemdServiceRestart_On_Abort
//	SystemdServiceRestartAlways
type SystemdServiceRestartInput interface {
	pulumi.Input

	ToSystemdServiceRestartOutput() SystemdServiceRestartOutput
	ToSystemdServiceRestartOutputWithContext(context.Context) SystemdServiceRestartOutput
}

var systemdServiceRestartPtrType = reflect.TypeOf((**SystemdServiceRestart)(nil)).Elem()

type SystemdServiceRestartPtrInput interface {
	pulumi.Input

	ToSystemdServiceRestartPtrOutput() SystemdServiceRestartPtrOutput
	ToSystemdServiceRestartPtrOutputWithContext(context.Context) SystemdServiceRestartPtrOutput
}

type systemdServiceRestartPtr string

func SystemdServiceRestartPtr(v string) SystemdServiceRestartPtrInput {
	return (*systemdServiceRestartPtr)(&v)
}

func (*systemdServiceRestartPtr) ElementType() reflect.Type {
	return systemdServiceRestartPtrType
}

func (in *systemdServiceRestartPtr) ToSystemdServiceRestartPtrOutput() SystemdServiceRestartPtrOutput {
	return pulumi.ToOutput(in).(SystemdServiceRestartPtrOutput)
}

func (in *systemdServiceRestartPtr) ToSystemdServiceRestartPtrOutputWithContext(ctx context.Context) SystemdServiceRestartPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SystemdServiceRestartPtrOutput)
}

// Systemd service type.
type SystemdServiceType string

const (
	SystemdServiceTypeSimple         = SystemdServiceType("simple")
	SystemdServiceTypeExec           = SystemdServiceType("exec")
	SystemdServiceTypeForking        = SystemdServiceType("forking")
	SystemdServiceTypeOneshot        = SystemdServiceType("oneshot")
	SystemdServiceTypeDbus           = SystemdServiceType("dbus")
	SystemdServiceTypeNotify         = SystemdServiceType("notify")
	SystemdServiceType_Notify_Reload = SystemdServiceType("notify-reload")
	SystemdServiceTypeIdle           = SystemdServiceType("idle")
)

func (SystemdServiceType) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemdServiceType)(nil)).Elem()
}

func (e SystemdServiceType) ToSystemdServiceTypeOutput() SystemdServiceTypeOutput {
	return pulumi.ToOutput(e).(SystemdServiceTypeOutput)
}

func (e SystemdServiceType) ToSystemdServiceTypeOutputWithContext(ctx context.Context) SystemdServiceTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SystemdServiceTypeOutput)
}

func (e SystemdServiceType) ToSystemdServiceTypePtrOutput() SystemdServiceTypePtrOutput {
	return e.ToSystemdServiceTypePtrOutputWithContext(context.Background())
}

func (e SystemdServiceType) ToSystemdServiceTypePtrOutputWithContext(ctx context.Context) SystemdServiceTypePtrOutput {
	return SystemdServiceType(e).ToSystemdServiceTypeOutputWithContext(ctx).ToSystemdServiceTypePtrOutputWithContext(ctx)
}

func (e SystemdServiceType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SystemdServiceType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SystemdServiceType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SystemdServiceType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SystemdServiceTypeOutput struct{ *pulumi.OutputState }

func (SystemdServiceTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemdServiceType)(nil)).Elem()
}

func (o SystemdServiceTypeOutput) ToSystemdServiceTypeOutput() SystemdServiceTypeOutput {
	return o
}

func (o SystemdServiceTypeOutput) ToSystemdServiceTypeOutputWithContext(ctx context.Context) SystemdServiceTypeOutput {
	return o
}

func (o SystemdServiceTypeOutput) ToSystemdServiceTypePtrOutput() SystemdServiceTypePtrOutput {
	return o.ToSystemdServiceTypePtrOutputWithContext(context.Background())
}

func (o SystemdServiceTypeOutput) ToSystemdServiceTypePtrOutputWithContext(ctx context.Context) SystemdServiceTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SystemdServiceType) *SystemdServiceType {
		return &v
	}).(SystemdServiceTypePtrOutput)
}

func (o SystemdServiceTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SystemdServiceTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SystemdServiceType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SystemdServiceTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SystemdServiceTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SystemdServiceType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SystemdServiceTypePtrOutput struct{ *pulumi.OutputState }

func (SystemdServiceTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemdServiceType)(nil)).Elem()
}

func (o SystemdServiceTypePtrOutput) ToSystemdServiceTypePtrOutput() SystemdServiceTypePtrOutput {
	return o
}

func (o SystemdServiceTypePtrOutput) ToSystemdServiceTypePtrOutputWithContext(ctx context.Context) SystemdServiceTypePtrOutput {
	return o
}

func (o SystemdServiceTypePtrOutput) Elem() SystemdServiceTypeOutput {
	return o.ApplyT(func(v *SystemdServiceType) SystemdServiceType {
		if v != nil {
			return *v
		}
		var ret SystemdServiceType
		return ret
	}).(SystemdServiceTypeOutput)
}

func (o SystemdServiceTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SystemdServiceTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SystemdServiceType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// SystemdServiceTypeInput is an input type that accepts values of the SystemdServiceType enum
// A concrete instance of `SystemdServiceTypeInput` can be one of the following:
//
//	SystemdServiceTypeSimple
//	SystemdServiceTypeExec
//	SystemdServiceTypeForking
//	SystemdServiceTypeOneshot
//	SystemdServiceTypeDbus
//	SystemdServiceTypeNotify
//	SystemdServiceType_Notify_Reload
//	SystemdServiceTypeIdle
type SystemdServiceTypeInput interface {
	pulumi.Input

	ToSystemdServiceTypeOutput() SystemdServiceTypeOutput
	ToSystemdServiceTypeOutputWithContext(context.Context) SystemdServiceTypeOutput
}

var systemdServiceTypePtrType = reflect.TypeOf((**SystemdServiceType)(nil)).Elem()

type SystemdServiceTypePtrInput interface {
	pulumi.Input

	ToSystemdServiceTypePtrOutput() SystemdServiceTypePtrOutput
	ToSystemdServiceTypePtrOutputWithContext(context.Context) SystemdServiceTypePtrOutput
}

type systemdServiceTypePtr string

func SystemdServiceTypePtr(v string) SystemdServiceTypePtrInput {
	return (*systemdServiceTypePtr)(&v)
}

func (*systemdServiceTypePtr) ElementType() reflect.Type {
	return systemdServiceTypePtrType
}

func (in *systemdServiceTypePtr) ToSystemdServiceTypePtrOutput() SystemdServiceTypePtrOutput {
	return pulumi.ToOutput(in).(SystemdServiceTypePtrOutput)
}

func (in *systemdServiceTypePtr) ToSystemdServiceTypePtrOutputWithContext(ctx context.Context) SystemdServiceTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SystemdServiceTypePtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ArchitectureInput)(nil)).Elem(), Architecture("amd64"))
	pulumi.RegisterInputType(reflect.TypeOf((*ArchitecturePtrInput)(nil)).Elem(), Architecture("amd64"))
	pulumi.RegisterInputType(reflect.TypeOf((*SystemdServiceExitTypeInput)(nil)).Elem(), SystemdServiceExitType("main"))
	pulumi.RegisterInputType(reflect.TypeOf((*SystemdServiceExitTypePtrInput)(nil)).Elem(), SystemdServiceExitType("main"))
	pulumi.RegisterInputType(reflect.TypeOf((*SystemdServiceRestartInput)(nil)).Elem(), SystemdServiceRestart("no"))
	pulumi.RegisterInputType(reflect.TypeOf((*SystemdServiceRestartPtrInput)(nil)).Elem(), SystemdServiceRestart("no"))
	pulumi.RegisterInputType(reflect.TypeOf((*SystemdServiceTypeInput)(nil)).Elem(), SystemdServiceType("simple"))
	pulumi.RegisterInputType(reflect.TypeOf((*SystemdServiceTypePtrInput)(nil)).Elem(), SystemdServiceType("simple"))
	pulumi.RegisterOutputType(ArchitectureOutput{})
	pulumi.RegisterOutputType(ArchitecturePtrOutput{})
	pulumi.RegisterOutputType(SystemdServiceExitTypeOutput{})
	pulumi.RegisterOutputType(SystemdServiceExitTypePtrOutput{})
	pulumi.RegisterOutputType(SystemdServiceRestartOutput{})
	pulumi.RegisterOutputType(SystemdServiceRestartPtrOutput{})
	pulumi.RegisterOutputType(SystemdServiceTypeOutput{})
	pulumi.RegisterOutputType(SystemdServiceTypePtrOutput{})
}

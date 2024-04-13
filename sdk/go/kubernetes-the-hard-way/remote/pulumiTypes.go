// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package remote

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/unstoppablemango/pulumi-kubernetes-the-hard-way/sdk/go/kubernetes-the-hard-way/internal"
)

var _ = internal.GetEnvOrDefault

// Props for resources that consume etcd configuration.
type EtcdConfigurationProps struct {
	// Path to the certificate authority file on the remote system.
	CaFilePath string `pulumi:"caFilePath"`
	// Path to the certificate file on the remote system.
	CertFilePath string `pulumi:"certFilePath"`
	// Etcd's data directory.
	DataDirectory string `pulumi:"dataDirectory"`
	// Path to the etcd binary.
	EtcdPath string `pulumi:"etcdPath"`
	// Internal IP of the etcd node.
	InternalIp string `pulumi:"internalIp"`
	// Path to the private key file on the remote system.
	KeyFilePath string `pulumi:"keyFilePath"`
	// Name of the etcd node.
	Name string `pulumi:"name"`
}

// EtcdConfigurationPropsInput is an input type that accepts EtcdConfigurationPropsArgs and EtcdConfigurationPropsOutput values.
// You can construct a concrete instance of `EtcdConfigurationPropsInput` via:
//
//	EtcdConfigurationPropsArgs{...}
type EtcdConfigurationPropsInput interface {
	pulumi.Input

	ToEtcdConfigurationPropsOutput() EtcdConfigurationPropsOutput
	ToEtcdConfigurationPropsOutputWithContext(context.Context) EtcdConfigurationPropsOutput
}

// Props for resources that consume etcd configuration.
type EtcdConfigurationPropsArgs struct {
	// Path to the certificate authority file on the remote system.
	CaFilePath pulumi.StringInput `pulumi:"caFilePath"`
	// Path to the certificate file on the remote system.
	CertFilePath pulumi.StringInput `pulumi:"certFilePath"`
	// Etcd's data directory.
	DataDirectory pulumi.StringInput `pulumi:"dataDirectory"`
	// Path to the etcd binary.
	EtcdPath pulumi.StringInput `pulumi:"etcdPath"`
	// Internal IP of the etcd node.
	InternalIp pulumi.StringInput `pulumi:"internalIp"`
	// Path to the private key file on the remote system.
	KeyFilePath pulumi.StringInput `pulumi:"keyFilePath"`
	// Name of the etcd node.
	Name pulumi.StringInput `pulumi:"name"`
}

func (EtcdConfigurationPropsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EtcdConfigurationProps)(nil)).Elem()
}

func (i EtcdConfigurationPropsArgs) ToEtcdConfigurationPropsOutput() EtcdConfigurationPropsOutput {
	return i.ToEtcdConfigurationPropsOutputWithContext(context.Background())
}

func (i EtcdConfigurationPropsArgs) ToEtcdConfigurationPropsOutputWithContext(ctx context.Context) EtcdConfigurationPropsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EtcdConfigurationPropsOutput)
}

// EtcdConfigurationPropsArrayInput is an input type that accepts EtcdConfigurationPropsArray and EtcdConfigurationPropsArrayOutput values.
// You can construct a concrete instance of `EtcdConfigurationPropsArrayInput` via:
//
//	EtcdConfigurationPropsArray{ EtcdConfigurationPropsArgs{...} }
type EtcdConfigurationPropsArrayInput interface {
	pulumi.Input

	ToEtcdConfigurationPropsArrayOutput() EtcdConfigurationPropsArrayOutput
	ToEtcdConfigurationPropsArrayOutputWithContext(context.Context) EtcdConfigurationPropsArrayOutput
}

type EtcdConfigurationPropsArray []EtcdConfigurationPropsInput

func (EtcdConfigurationPropsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EtcdConfigurationProps)(nil)).Elem()
}

func (i EtcdConfigurationPropsArray) ToEtcdConfigurationPropsArrayOutput() EtcdConfigurationPropsArrayOutput {
	return i.ToEtcdConfigurationPropsArrayOutputWithContext(context.Background())
}

func (i EtcdConfigurationPropsArray) ToEtcdConfigurationPropsArrayOutputWithContext(ctx context.Context) EtcdConfigurationPropsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EtcdConfigurationPropsArrayOutput)
}

// Props for resources that consume etcd configuration.
type EtcdConfigurationPropsOutput struct{ *pulumi.OutputState }

func (EtcdConfigurationPropsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EtcdConfigurationProps)(nil)).Elem()
}

func (o EtcdConfigurationPropsOutput) ToEtcdConfigurationPropsOutput() EtcdConfigurationPropsOutput {
	return o
}

func (o EtcdConfigurationPropsOutput) ToEtcdConfigurationPropsOutputWithContext(ctx context.Context) EtcdConfigurationPropsOutput {
	return o
}

// Path to the certificate authority file on the remote system.
func (o EtcdConfigurationPropsOutput) CaFilePath() pulumi.StringOutput {
	return o.ApplyT(func(v EtcdConfigurationProps) string { return v.CaFilePath }).(pulumi.StringOutput)
}

// Path to the certificate file on the remote system.
func (o EtcdConfigurationPropsOutput) CertFilePath() pulumi.StringOutput {
	return o.ApplyT(func(v EtcdConfigurationProps) string { return v.CertFilePath }).(pulumi.StringOutput)
}

// Etcd's data directory.
func (o EtcdConfigurationPropsOutput) DataDirectory() pulumi.StringOutput {
	return o.ApplyT(func(v EtcdConfigurationProps) string { return v.DataDirectory }).(pulumi.StringOutput)
}

// Path to the etcd binary.
func (o EtcdConfigurationPropsOutput) EtcdPath() pulumi.StringOutput {
	return o.ApplyT(func(v EtcdConfigurationProps) string { return v.EtcdPath }).(pulumi.StringOutput)
}

// Internal IP of the etcd node.
func (o EtcdConfigurationPropsOutput) InternalIp() pulumi.StringOutput {
	return o.ApplyT(func(v EtcdConfigurationProps) string { return v.InternalIp }).(pulumi.StringOutput)
}

// Path to the private key file on the remote system.
func (o EtcdConfigurationPropsOutput) KeyFilePath() pulumi.StringOutput {
	return o.ApplyT(func(v EtcdConfigurationProps) string { return v.KeyFilePath }).(pulumi.StringOutput)
}

// Name of the etcd node.
func (o EtcdConfigurationPropsOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v EtcdConfigurationProps) string { return v.Name }).(pulumi.StringOutput)
}

type EtcdConfigurationPropsArrayOutput struct{ *pulumi.OutputState }

func (EtcdConfigurationPropsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EtcdConfigurationProps)(nil)).Elem()
}

func (o EtcdConfigurationPropsArrayOutput) ToEtcdConfigurationPropsArrayOutput() EtcdConfigurationPropsArrayOutput {
	return o
}

func (o EtcdConfigurationPropsArrayOutput) ToEtcdConfigurationPropsArrayOutputWithContext(ctx context.Context) EtcdConfigurationPropsArrayOutput {
	return o
}

func (o EtcdConfigurationPropsArrayOutput) Index(i pulumi.IntInput) EtcdConfigurationPropsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EtcdConfigurationProps {
		return vs[0].([]EtcdConfigurationProps)[vs[1].(int)]
	}).(EtcdConfigurationPropsOutput)
}

// https://www.freedesktop.org/software/systemd/man/latest/systemd.unit.html#%5BInstall%5D%20Section%20Options
type SystemdInstallSection struct {
	// A symbolic link is created in the .wants/, .requires/, or .upholds/ directory of each of the listed units when this unit is installed by systemctl enable.
	WantedBy []string `pulumi:"wantedBy"`
}

// SystemdInstallSectionInput is an input type that accepts SystemdInstallSectionArgs and SystemdInstallSectionOutput values.
// You can construct a concrete instance of `SystemdInstallSectionInput` via:
//
//	SystemdInstallSectionArgs{...}
type SystemdInstallSectionInput interface {
	pulumi.Input

	ToSystemdInstallSectionOutput() SystemdInstallSectionOutput
	ToSystemdInstallSectionOutputWithContext(context.Context) SystemdInstallSectionOutput
}

// https://www.freedesktop.org/software/systemd/man/latest/systemd.unit.html#%5BInstall%5D%20Section%20Options
type SystemdInstallSectionArgs struct {
	// A symbolic link is created in the .wants/, .requires/, or .upholds/ directory of each of the listed units when this unit is installed by systemctl enable.
	WantedBy pulumi.StringArrayInput `pulumi:"wantedBy"`
}

func (SystemdInstallSectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemdInstallSection)(nil)).Elem()
}

func (i SystemdInstallSectionArgs) ToSystemdInstallSectionOutput() SystemdInstallSectionOutput {
	return i.ToSystemdInstallSectionOutputWithContext(context.Background())
}

func (i SystemdInstallSectionArgs) ToSystemdInstallSectionOutputWithContext(ctx context.Context) SystemdInstallSectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemdInstallSectionOutput)
}

func (i SystemdInstallSectionArgs) ToSystemdInstallSectionPtrOutput() SystemdInstallSectionPtrOutput {
	return i.ToSystemdInstallSectionPtrOutputWithContext(context.Background())
}

func (i SystemdInstallSectionArgs) ToSystemdInstallSectionPtrOutputWithContext(ctx context.Context) SystemdInstallSectionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemdInstallSectionOutput).ToSystemdInstallSectionPtrOutputWithContext(ctx)
}

// SystemdInstallSectionPtrInput is an input type that accepts SystemdInstallSectionArgs, SystemdInstallSectionPtr and SystemdInstallSectionPtrOutput values.
// You can construct a concrete instance of `SystemdInstallSectionPtrInput` via:
//
//	        SystemdInstallSectionArgs{...}
//
//	or:
//
//	        nil
type SystemdInstallSectionPtrInput interface {
	pulumi.Input

	ToSystemdInstallSectionPtrOutput() SystemdInstallSectionPtrOutput
	ToSystemdInstallSectionPtrOutputWithContext(context.Context) SystemdInstallSectionPtrOutput
}

type systemdInstallSectionPtrType SystemdInstallSectionArgs

func SystemdInstallSectionPtr(v *SystemdInstallSectionArgs) SystemdInstallSectionPtrInput {
	return (*systemdInstallSectionPtrType)(v)
}

func (*systemdInstallSectionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemdInstallSection)(nil)).Elem()
}

func (i *systemdInstallSectionPtrType) ToSystemdInstallSectionPtrOutput() SystemdInstallSectionPtrOutput {
	return i.ToSystemdInstallSectionPtrOutputWithContext(context.Background())
}

func (i *systemdInstallSectionPtrType) ToSystemdInstallSectionPtrOutputWithContext(ctx context.Context) SystemdInstallSectionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemdInstallSectionPtrOutput)
}

// https://www.freedesktop.org/software/systemd/man/latest/systemd.unit.html#%5BInstall%5D%20Section%20Options
type SystemdInstallSectionOutput struct{ *pulumi.OutputState }

func (SystemdInstallSectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemdInstallSection)(nil)).Elem()
}

func (o SystemdInstallSectionOutput) ToSystemdInstallSectionOutput() SystemdInstallSectionOutput {
	return o
}

func (o SystemdInstallSectionOutput) ToSystemdInstallSectionOutputWithContext(ctx context.Context) SystemdInstallSectionOutput {
	return o
}

func (o SystemdInstallSectionOutput) ToSystemdInstallSectionPtrOutput() SystemdInstallSectionPtrOutput {
	return o.ToSystemdInstallSectionPtrOutputWithContext(context.Background())
}

func (o SystemdInstallSectionOutput) ToSystemdInstallSectionPtrOutputWithContext(ctx context.Context) SystemdInstallSectionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SystemdInstallSection) *SystemdInstallSection {
		return &v
	}).(SystemdInstallSectionPtrOutput)
}

// A symbolic link is created in the .wants/, .requires/, or .upholds/ directory of each of the listed units when this unit is installed by systemctl enable.
func (o SystemdInstallSectionOutput) WantedBy() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SystemdInstallSection) []string { return v.WantedBy }).(pulumi.StringArrayOutput)
}

type SystemdInstallSectionPtrOutput struct{ *pulumi.OutputState }

func (SystemdInstallSectionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemdInstallSection)(nil)).Elem()
}

func (o SystemdInstallSectionPtrOutput) ToSystemdInstallSectionPtrOutput() SystemdInstallSectionPtrOutput {
	return o
}

func (o SystemdInstallSectionPtrOutput) ToSystemdInstallSectionPtrOutputWithContext(ctx context.Context) SystemdInstallSectionPtrOutput {
	return o
}

func (o SystemdInstallSectionPtrOutput) Elem() SystemdInstallSectionOutput {
	return o.ApplyT(func(v *SystemdInstallSection) SystemdInstallSection {
		if v != nil {
			return *v
		}
		var ret SystemdInstallSection
		return ret
	}).(SystemdInstallSectionOutput)
}

// A symbolic link is created in the .wants/, .requires/, or .upholds/ directory of each of the listed units when this unit is installed by systemctl enable.
func (o SystemdInstallSectionPtrOutput) WantedBy() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SystemdInstallSection) []string {
		if v == nil {
			return nil
		}
		return v.WantedBy
	}).(pulumi.StringArrayOutput)
}

// https://www.freedesktop.org/software/systemd/man/latest/systemd.service.html#
type SystemdServiceSection struct {
	// Commands that are executed when this service is started.
	ExecStart *string `pulumi:"execStart"`
	// Specifies when the manager should consider the service to be finished.
	ExitType *SystemdServiceExitType `pulumi:"exitType"`
	// Configures whether the service shall be restarted when the service process exits, is killed, or a timeout is reached.
	Restart *SystemdServiceRestart `pulumi:"restart"`
	// Configures the time to sleep before restarting a service (as configured with Restart=).
	RestartSec *string `pulumi:"restartSec"`
	// Configures the mechanism via which the service notifies the manager that the service start-up has finished.
	Type *SystemdServiceType `pulumi:"type"`
}

// SystemdServiceSectionInput is an input type that accepts SystemdServiceSectionArgs and SystemdServiceSectionOutput values.
// You can construct a concrete instance of `SystemdServiceSectionInput` via:
//
//	SystemdServiceSectionArgs{...}
type SystemdServiceSectionInput interface {
	pulumi.Input

	ToSystemdServiceSectionOutput() SystemdServiceSectionOutput
	ToSystemdServiceSectionOutputWithContext(context.Context) SystemdServiceSectionOutput
}

// https://www.freedesktop.org/software/systemd/man/latest/systemd.service.html#
type SystemdServiceSectionArgs struct {
	// Commands that are executed when this service is started.
	ExecStart pulumi.StringPtrInput `pulumi:"execStart"`
	// Specifies when the manager should consider the service to be finished.
	ExitType SystemdServiceExitTypePtrInput `pulumi:"exitType"`
	// Configures whether the service shall be restarted when the service process exits, is killed, or a timeout is reached.
	Restart SystemdServiceRestartPtrInput `pulumi:"restart"`
	// Configures the time to sleep before restarting a service (as configured with Restart=).
	RestartSec pulumi.StringPtrInput `pulumi:"restartSec"`
	// Configures the mechanism via which the service notifies the manager that the service start-up has finished.
	Type SystemdServiceTypePtrInput `pulumi:"type"`
}

func (SystemdServiceSectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemdServiceSection)(nil)).Elem()
}

func (i SystemdServiceSectionArgs) ToSystemdServiceSectionOutput() SystemdServiceSectionOutput {
	return i.ToSystemdServiceSectionOutputWithContext(context.Background())
}

func (i SystemdServiceSectionArgs) ToSystemdServiceSectionOutputWithContext(ctx context.Context) SystemdServiceSectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemdServiceSectionOutput)
}

// https://www.freedesktop.org/software/systemd/man/latest/systemd.service.html#
type SystemdServiceSectionOutput struct{ *pulumi.OutputState }

func (SystemdServiceSectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemdServiceSection)(nil)).Elem()
}

func (o SystemdServiceSectionOutput) ToSystemdServiceSectionOutput() SystemdServiceSectionOutput {
	return o
}

func (o SystemdServiceSectionOutput) ToSystemdServiceSectionOutputWithContext(ctx context.Context) SystemdServiceSectionOutput {
	return o
}

// Commands that are executed when this service is started.
func (o SystemdServiceSectionOutput) ExecStart() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemdServiceSection) *string { return v.ExecStart }).(pulumi.StringPtrOutput)
}

// Specifies when the manager should consider the service to be finished.
func (o SystemdServiceSectionOutput) ExitType() SystemdServiceExitTypePtrOutput {
	return o.ApplyT(func(v SystemdServiceSection) *SystemdServiceExitType { return v.ExitType }).(SystemdServiceExitTypePtrOutput)
}

// Configures whether the service shall be restarted when the service process exits, is killed, or a timeout is reached.
func (o SystemdServiceSectionOutput) Restart() SystemdServiceRestartPtrOutput {
	return o.ApplyT(func(v SystemdServiceSection) *SystemdServiceRestart { return v.Restart }).(SystemdServiceRestartPtrOutput)
}

// Configures the time to sleep before restarting a service (as configured with Restart=).
func (o SystemdServiceSectionOutput) RestartSec() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemdServiceSection) *string { return v.RestartSec }).(pulumi.StringPtrOutput)
}

// Configures the mechanism via which the service notifies the manager that the service start-up has finished.
func (o SystemdServiceSectionOutput) Type() SystemdServiceTypePtrOutput {
	return o.ApplyT(func(v SystemdServiceSection) *SystemdServiceType { return v.Type }).(SystemdServiceTypePtrOutput)
}

// https://www.freedesktop.org/software/systemd/man/latest/systemd.unit.html#
type SystemdUnitSection struct {
	// Configures requirement dependencies, very similar in style to Requires=.
	BindsTo []string `pulumi:"bindsTo"`
	// A short human readable title of the unit.
	Description *string `pulumi:"description"`
	// A space-separated list of URIs referencing documentation for this unit or its configuration.
	Documentation []string `pulumi:"documentation"`
	// Similar to Wants=, but declares a stronger requirement dependency.
	Requires []string `pulumi:"requires"`
	// Similar to Requires=. However, if the units listed here are not started already, they will not be started and the starting of this unit will fail immediately.
	Requisite []string `pulumi:"requisite"`
	// Configures (weak) requirement dependencies on other units.
	Wants []string `pulumi:"wants"`
}

// SystemdUnitSectionInput is an input type that accepts SystemdUnitSectionArgs and SystemdUnitSectionOutput values.
// You can construct a concrete instance of `SystemdUnitSectionInput` via:
//
//	SystemdUnitSectionArgs{...}
type SystemdUnitSectionInput interface {
	pulumi.Input

	ToSystemdUnitSectionOutput() SystemdUnitSectionOutput
	ToSystemdUnitSectionOutputWithContext(context.Context) SystemdUnitSectionOutput
}

// https://www.freedesktop.org/software/systemd/man/latest/systemd.unit.html#
type SystemdUnitSectionArgs struct {
	// Configures requirement dependencies, very similar in style to Requires=.
	BindsTo pulumi.StringArrayInput `pulumi:"bindsTo"`
	// A short human readable title of the unit.
	Description pulumi.StringPtrInput `pulumi:"description"`
	// A space-separated list of URIs referencing documentation for this unit or its configuration.
	Documentation pulumi.StringArrayInput `pulumi:"documentation"`
	// Similar to Wants=, but declares a stronger requirement dependency.
	Requires pulumi.StringArrayInput `pulumi:"requires"`
	// Similar to Requires=. However, if the units listed here are not started already, they will not be started and the starting of this unit will fail immediately.
	Requisite pulumi.StringArrayInput `pulumi:"requisite"`
	// Configures (weak) requirement dependencies on other units.
	Wants pulumi.StringArrayInput `pulumi:"wants"`
}

func (SystemdUnitSectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemdUnitSection)(nil)).Elem()
}

func (i SystemdUnitSectionArgs) ToSystemdUnitSectionOutput() SystemdUnitSectionOutput {
	return i.ToSystemdUnitSectionOutputWithContext(context.Background())
}

func (i SystemdUnitSectionArgs) ToSystemdUnitSectionOutputWithContext(ctx context.Context) SystemdUnitSectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemdUnitSectionOutput)
}

func (i SystemdUnitSectionArgs) ToSystemdUnitSectionPtrOutput() SystemdUnitSectionPtrOutput {
	return i.ToSystemdUnitSectionPtrOutputWithContext(context.Background())
}

func (i SystemdUnitSectionArgs) ToSystemdUnitSectionPtrOutputWithContext(ctx context.Context) SystemdUnitSectionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemdUnitSectionOutput).ToSystemdUnitSectionPtrOutputWithContext(ctx)
}

// SystemdUnitSectionPtrInput is an input type that accepts SystemdUnitSectionArgs, SystemdUnitSectionPtr and SystemdUnitSectionPtrOutput values.
// You can construct a concrete instance of `SystemdUnitSectionPtrInput` via:
//
//	        SystemdUnitSectionArgs{...}
//
//	or:
//
//	        nil
type SystemdUnitSectionPtrInput interface {
	pulumi.Input

	ToSystemdUnitSectionPtrOutput() SystemdUnitSectionPtrOutput
	ToSystemdUnitSectionPtrOutputWithContext(context.Context) SystemdUnitSectionPtrOutput
}

type systemdUnitSectionPtrType SystemdUnitSectionArgs

func SystemdUnitSectionPtr(v *SystemdUnitSectionArgs) SystemdUnitSectionPtrInput {
	return (*systemdUnitSectionPtrType)(v)
}

func (*systemdUnitSectionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemdUnitSection)(nil)).Elem()
}

func (i *systemdUnitSectionPtrType) ToSystemdUnitSectionPtrOutput() SystemdUnitSectionPtrOutput {
	return i.ToSystemdUnitSectionPtrOutputWithContext(context.Background())
}

func (i *systemdUnitSectionPtrType) ToSystemdUnitSectionPtrOutputWithContext(ctx context.Context) SystemdUnitSectionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemdUnitSectionPtrOutput)
}

// https://www.freedesktop.org/software/systemd/man/latest/systemd.unit.html#
type SystemdUnitSectionOutput struct{ *pulumi.OutputState }

func (SystemdUnitSectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemdUnitSection)(nil)).Elem()
}

func (o SystemdUnitSectionOutput) ToSystemdUnitSectionOutput() SystemdUnitSectionOutput {
	return o
}

func (o SystemdUnitSectionOutput) ToSystemdUnitSectionOutputWithContext(ctx context.Context) SystemdUnitSectionOutput {
	return o
}

func (o SystemdUnitSectionOutput) ToSystemdUnitSectionPtrOutput() SystemdUnitSectionPtrOutput {
	return o.ToSystemdUnitSectionPtrOutputWithContext(context.Background())
}

func (o SystemdUnitSectionOutput) ToSystemdUnitSectionPtrOutputWithContext(ctx context.Context) SystemdUnitSectionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SystemdUnitSection) *SystemdUnitSection {
		return &v
	}).(SystemdUnitSectionPtrOutput)
}

// Configures requirement dependencies, very similar in style to Requires=.
func (o SystemdUnitSectionOutput) BindsTo() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SystemdUnitSection) []string { return v.BindsTo }).(pulumi.StringArrayOutput)
}

// A short human readable title of the unit.
func (o SystemdUnitSectionOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemdUnitSection) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// A space-separated list of URIs referencing documentation for this unit or its configuration.
func (o SystemdUnitSectionOutput) Documentation() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SystemdUnitSection) []string { return v.Documentation }).(pulumi.StringArrayOutput)
}

// Similar to Wants=, but declares a stronger requirement dependency.
func (o SystemdUnitSectionOutput) Requires() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SystemdUnitSection) []string { return v.Requires }).(pulumi.StringArrayOutput)
}

// Similar to Requires=. However, if the units listed here are not started already, they will not be started and the starting of this unit will fail immediately.
func (o SystemdUnitSectionOutput) Requisite() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SystemdUnitSection) []string { return v.Requisite }).(pulumi.StringArrayOutput)
}

// Configures (weak) requirement dependencies on other units.
func (o SystemdUnitSectionOutput) Wants() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SystemdUnitSection) []string { return v.Wants }).(pulumi.StringArrayOutput)
}

type SystemdUnitSectionPtrOutput struct{ *pulumi.OutputState }

func (SystemdUnitSectionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemdUnitSection)(nil)).Elem()
}

func (o SystemdUnitSectionPtrOutput) ToSystemdUnitSectionPtrOutput() SystemdUnitSectionPtrOutput {
	return o
}

func (o SystemdUnitSectionPtrOutput) ToSystemdUnitSectionPtrOutputWithContext(ctx context.Context) SystemdUnitSectionPtrOutput {
	return o
}

func (o SystemdUnitSectionPtrOutput) Elem() SystemdUnitSectionOutput {
	return o.ApplyT(func(v *SystemdUnitSection) SystemdUnitSection {
		if v != nil {
			return *v
		}
		var ret SystemdUnitSection
		return ret
	}).(SystemdUnitSectionOutput)
}

// Configures requirement dependencies, very similar in style to Requires=.
func (o SystemdUnitSectionPtrOutput) BindsTo() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SystemdUnitSection) []string {
		if v == nil {
			return nil
		}
		return v.BindsTo
	}).(pulumi.StringArrayOutput)
}

// A short human readable title of the unit.
func (o SystemdUnitSectionPtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemdUnitSection) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

// A space-separated list of URIs referencing documentation for this unit or its configuration.
func (o SystemdUnitSectionPtrOutput) Documentation() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SystemdUnitSection) []string {
		if v == nil {
			return nil
		}
		return v.Documentation
	}).(pulumi.StringArrayOutput)
}

// Similar to Wants=, but declares a stronger requirement dependency.
func (o SystemdUnitSectionPtrOutput) Requires() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SystemdUnitSection) []string {
		if v == nil {
			return nil
		}
		return v.Requires
	}).(pulumi.StringArrayOutput)
}

// Similar to Requires=. However, if the units listed here are not started already, they will not be started and the starting of this unit will fail immediately.
func (o SystemdUnitSectionPtrOutput) Requisite() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SystemdUnitSection) []string {
		if v == nil {
			return nil
		}
		return v.Requisite
	}).(pulumi.StringArrayOutput)
}

// Configures (weak) requirement dependencies on other units.
func (o SystemdUnitSectionPtrOutput) Wants() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SystemdUnitSection) []string {
		if v == nil {
			return nil
		}
		return v.Wants
	}).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EtcdConfigurationPropsInput)(nil)).Elem(), EtcdConfigurationPropsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EtcdConfigurationPropsArrayInput)(nil)).Elem(), EtcdConfigurationPropsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemdInstallSectionInput)(nil)).Elem(), SystemdInstallSectionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemdInstallSectionPtrInput)(nil)).Elem(), SystemdInstallSectionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemdServiceSectionInput)(nil)).Elem(), SystemdServiceSectionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemdUnitSectionInput)(nil)).Elem(), SystemdUnitSectionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemdUnitSectionPtrInput)(nil)).Elem(), SystemdUnitSectionArgs{})
	pulumi.RegisterOutputType(EtcdConfigurationPropsOutput{})
	pulumi.RegisterOutputType(EtcdConfigurationPropsArrayOutput{})
	pulumi.RegisterOutputType(SystemdInstallSectionOutput{})
	pulumi.RegisterOutputType(SystemdInstallSectionPtrOutput{})
	pulumi.RegisterOutputType(SystemdServiceSectionOutput{})
	pulumi.RegisterOutputType(SystemdUnitSectionOutput{})
	pulumi.RegisterOutputType(SystemdUnitSectionPtrOutput{})
}

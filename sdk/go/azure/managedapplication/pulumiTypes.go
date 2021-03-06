// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package managedapplication

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type DefinitionAuthorization struct {
	// Specifies a role definition identifier for the provider. This role will define all the permissions that the provider must have on the managed application's container resource group. This role definition cannot have permission to delete the resource group.
	RoleDefinitionId string `pulumi:"roleDefinitionId"`
	// Specifies a service principal identifier for the provider. This is the identity that the provider will use to call ARM to manage the managed application resources.
	ServicePrincipalId string `pulumi:"servicePrincipalId"`
}

// DefinitionAuthorizationInput is an input type that accepts DefinitionAuthorizationArgs and DefinitionAuthorizationOutput values.
// You can construct a concrete instance of `DefinitionAuthorizationInput` via:
//
// 		 DefinitionAuthorizationArgs{...}
//
type DefinitionAuthorizationInput interface {
	pulumi.Input

	ToDefinitionAuthorizationOutput() DefinitionAuthorizationOutput
	ToDefinitionAuthorizationOutputWithContext(context.Context) DefinitionAuthorizationOutput
}

type DefinitionAuthorizationArgs struct {
	// Specifies a role definition identifier for the provider. This role will define all the permissions that the provider must have on the managed application's container resource group. This role definition cannot have permission to delete the resource group.
	RoleDefinitionId pulumi.StringInput `pulumi:"roleDefinitionId"`
	// Specifies a service principal identifier for the provider. This is the identity that the provider will use to call ARM to manage the managed application resources.
	ServicePrincipalId pulumi.StringInput `pulumi:"servicePrincipalId"`
}

func (DefinitionAuthorizationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DefinitionAuthorization)(nil)).Elem()
}

func (i DefinitionAuthorizationArgs) ToDefinitionAuthorizationOutput() DefinitionAuthorizationOutput {
	return i.ToDefinitionAuthorizationOutputWithContext(context.Background())
}

func (i DefinitionAuthorizationArgs) ToDefinitionAuthorizationOutputWithContext(ctx context.Context) DefinitionAuthorizationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefinitionAuthorizationOutput)
}

// DefinitionAuthorizationArrayInput is an input type that accepts DefinitionAuthorizationArray and DefinitionAuthorizationArrayOutput values.
// You can construct a concrete instance of `DefinitionAuthorizationArrayInput` via:
//
// 		 DefinitionAuthorizationArray{ DefinitionAuthorizationArgs{...} }
//
type DefinitionAuthorizationArrayInput interface {
	pulumi.Input

	ToDefinitionAuthorizationArrayOutput() DefinitionAuthorizationArrayOutput
	ToDefinitionAuthorizationArrayOutputWithContext(context.Context) DefinitionAuthorizationArrayOutput
}

type DefinitionAuthorizationArray []DefinitionAuthorizationInput

func (DefinitionAuthorizationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DefinitionAuthorization)(nil)).Elem()
}

func (i DefinitionAuthorizationArray) ToDefinitionAuthorizationArrayOutput() DefinitionAuthorizationArrayOutput {
	return i.ToDefinitionAuthorizationArrayOutputWithContext(context.Background())
}

func (i DefinitionAuthorizationArray) ToDefinitionAuthorizationArrayOutputWithContext(ctx context.Context) DefinitionAuthorizationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefinitionAuthorizationArrayOutput)
}

type DefinitionAuthorizationOutput struct{ *pulumi.OutputState }

func (DefinitionAuthorizationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DefinitionAuthorization)(nil)).Elem()
}

func (o DefinitionAuthorizationOutput) ToDefinitionAuthorizationOutput() DefinitionAuthorizationOutput {
	return o
}

func (o DefinitionAuthorizationOutput) ToDefinitionAuthorizationOutputWithContext(ctx context.Context) DefinitionAuthorizationOutput {
	return o
}

// Specifies a role definition identifier for the provider. This role will define all the permissions that the provider must have on the managed application's container resource group. This role definition cannot have permission to delete the resource group.
func (o DefinitionAuthorizationOutput) RoleDefinitionId() pulumi.StringOutput {
	return o.ApplyT(func(v DefinitionAuthorization) string { return v.RoleDefinitionId }).(pulumi.StringOutput)
}

// Specifies a service principal identifier for the provider. This is the identity that the provider will use to call ARM to manage the managed application resources.
func (o DefinitionAuthorizationOutput) ServicePrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v DefinitionAuthorization) string { return v.ServicePrincipalId }).(pulumi.StringOutput)
}

type DefinitionAuthorizationArrayOutput struct{ *pulumi.OutputState }

func (DefinitionAuthorizationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DefinitionAuthorization)(nil)).Elem()
}

func (o DefinitionAuthorizationArrayOutput) ToDefinitionAuthorizationArrayOutput() DefinitionAuthorizationArrayOutput {
	return o
}

func (o DefinitionAuthorizationArrayOutput) ToDefinitionAuthorizationArrayOutputWithContext(ctx context.Context) DefinitionAuthorizationArrayOutput {
	return o
}

func (o DefinitionAuthorizationArrayOutput) Index(i pulumi.IntInput) DefinitionAuthorizationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DefinitionAuthorization {
		return vs[0].([]DefinitionAuthorization)[vs[1].(int)]
	}).(DefinitionAuthorizationOutput)
}

func init() {
	pulumi.RegisterOutputType(DefinitionAuthorizationOutput{})
	pulumi.RegisterOutputType(DefinitionAuthorizationArrayOutput{})
}

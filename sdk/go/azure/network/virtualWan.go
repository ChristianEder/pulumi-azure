// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Virtual WAN.
type VirtualWan struct {
	pulumi.CustomResourceState

	// Boolean flag to specify whether branch to branch traffic is allowed. Defaults to `true`.
	AllowBranchToBranchTraffic pulumi.BoolPtrOutput `pulumi:"allowBranchToBranchTraffic"`
	// Boolean flag to specify whether VNet to VNet traffic is allowed. Defaults to `false`.
	AllowVnetToVnetTraffic pulumi.BoolPtrOutput `pulumi:"allowVnetToVnetTraffic"`
	// Boolean flag to specify whether VPN encryption is disabled. Defaults to `false`.
	DisableVpnEncryption pulumi.BoolPtrOutput `pulumi:"disableVpnEncryption"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// Specifies the name of the Virtual WAN. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the Office365 local breakout category. Possible values include: `Optimize`, `OptimizeAndAllow`, `All`, `None`. Defaults to `None`.
	Office365LocalBreakoutCategory pulumi.StringPtrOutput `pulumi:"office365LocalBreakoutCategory"`
	// The name of the resource group in which to create the Virtual WAN. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the Virtual WAN.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Specifies the Virtual WAN type. Possible Values include: `Basic` and `Standard`. Defaults to `Standard`.
	Type pulumi.StringPtrOutput `pulumi:"type"`
}

// NewVirtualWan registers a new resource with the given unique name, arguments, and options.
func NewVirtualWan(ctx *pulumi.Context,
	name string, args *VirtualWanArgs, opts ...pulumi.ResourceOption) (*VirtualWan, error) {
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &VirtualWanArgs{}
	}
	var resource VirtualWan
	err := ctx.RegisterResource("azure:network/virtualWan:VirtualWan", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualWan gets an existing VirtualWan resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualWan(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualWanState, opts ...pulumi.ResourceOption) (*VirtualWan, error) {
	var resource VirtualWan
	err := ctx.ReadResource("azure:network/virtualWan:VirtualWan", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualWan resources.
type virtualWanState struct {
	// Boolean flag to specify whether branch to branch traffic is allowed. Defaults to `true`.
	AllowBranchToBranchTraffic *bool `pulumi:"allowBranchToBranchTraffic"`
	// Boolean flag to specify whether VNet to VNet traffic is allowed. Defaults to `false`.
	AllowVnetToVnetTraffic *bool `pulumi:"allowVnetToVnetTraffic"`
	// Boolean flag to specify whether VPN encryption is disabled. Defaults to `false`.
	DisableVpnEncryption *bool `pulumi:"disableVpnEncryption"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the Virtual WAN. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the Office365 local breakout category. Possible values include: `Optimize`, `OptimizeAndAllow`, `All`, `None`. Defaults to `None`.
	Office365LocalBreakoutCategory *string `pulumi:"office365LocalBreakoutCategory"`
	// The name of the resource group in which to create the Virtual WAN. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the Virtual WAN.
	Tags map[string]string `pulumi:"tags"`
	// Specifies the Virtual WAN type. Possible Values include: `Basic` and `Standard`. Defaults to `Standard`.
	Type *string `pulumi:"type"`
}

type VirtualWanState struct {
	// Boolean flag to specify whether branch to branch traffic is allowed. Defaults to `true`.
	AllowBranchToBranchTraffic pulumi.BoolPtrInput
	// Boolean flag to specify whether VNet to VNet traffic is allowed. Defaults to `false`.
	AllowVnetToVnetTraffic pulumi.BoolPtrInput
	// Boolean flag to specify whether VPN encryption is disabled. Defaults to `false`.
	DisableVpnEncryption pulumi.BoolPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the Virtual WAN. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the Office365 local breakout category. Possible values include: `Optimize`, `OptimizeAndAllow`, `All`, `None`. Defaults to `None`.
	Office365LocalBreakoutCategory pulumi.StringPtrInput
	// The name of the resource group in which to create the Virtual WAN. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// A mapping of tags to assign to the Virtual WAN.
	Tags pulumi.StringMapInput
	// Specifies the Virtual WAN type. Possible Values include: `Basic` and `Standard`. Defaults to `Standard`.
	Type pulumi.StringPtrInput
}

func (VirtualWanState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualWanState)(nil)).Elem()
}

type virtualWanArgs struct {
	// Boolean flag to specify whether branch to branch traffic is allowed. Defaults to `true`.
	AllowBranchToBranchTraffic *bool `pulumi:"allowBranchToBranchTraffic"`
	// Boolean flag to specify whether VNet to VNet traffic is allowed. Defaults to `false`.
	AllowVnetToVnetTraffic *bool `pulumi:"allowVnetToVnetTraffic"`
	// Boolean flag to specify whether VPN encryption is disabled. Defaults to `false`.
	DisableVpnEncryption *bool `pulumi:"disableVpnEncryption"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the Virtual WAN. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the Office365 local breakout category. Possible values include: `Optimize`, `OptimizeAndAllow`, `All`, `None`. Defaults to `None`.
	Office365LocalBreakoutCategory *string `pulumi:"office365LocalBreakoutCategory"`
	// The name of the resource group in which to create the Virtual WAN. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the Virtual WAN.
	Tags map[string]string `pulumi:"tags"`
	// Specifies the Virtual WAN type. Possible Values include: `Basic` and `Standard`. Defaults to `Standard`.
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a VirtualWan resource.
type VirtualWanArgs struct {
	// Boolean flag to specify whether branch to branch traffic is allowed. Defaults to `true`.
	AllowBranchToBranchTraffic pulumi.BoolPtrInput
	// Boolean flag to specify whether VNet to VNet traffic is allowed. Defaults to `false`.
	AllowVnetToVnetTraffic pulumi.BoolPtrInput
	// Boolean flag to specify whether VPN encryption is disabled. Defaults to `false`.
	DisableVpnEncryption pulumi.BoolPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the Virtual WAN. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the Office365 local breakout category. Possible values include: `Optimize`, `OptimizeAndAllow`, `All`, `None`. Defaults to `None`.
	Office365LocalBreakoutCategory pulumi.StringPtrInput
	// The name of the resource group in which to create the Virtual WAN. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// A mapping of tags to assign to the Virtual WAN.
	Tags pulumi.StringMapInput
	// Specifies the Virtual WAN type. Possible Values include: `Basic` and `Standard`. Defaults to `Standard`.
	Type pulumi.StringPtrInput
}

func (VirtualWanArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualWanArgs)(nil)).Elem()
}

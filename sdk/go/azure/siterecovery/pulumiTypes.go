// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package siterecovery

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type ReplicatedVMManagedDisk struct {
	// Id of disk that should be replicated.
	DiskId string `pulumi:"diskId"`
	// Storage account that should be used for caching.
	StagingStorageAccountId string `pulumi:"stagingStorageAccountId"`
	// What type should the disk be when a failover is done.
	TargetDiskType string `pulumi:"targetDiskType"`
	// What type should the disk be that holds the replication data.
	TargetReplicaDiskType string `pulumi:"targetReplicaDiskType"`
	// Resource group disk should belong to when a failover is done.
	TargetResourceGroupId string `pulumi:"targetResourceGroupId"`
}

// ReplicatedVMManagedDiskInput is an input type that accepts ReplicatedVMManagedDiskArgs and ReplicatedVMManagedDiskOutput values.
// You can construct a concrete instance of `ReplicatedVMManagedDiskInput` via:
//
// 		 ReplicatedVMManagedDiskArgs{...}
//
type ReplicatedVMManagedDiskInput interface {
	pulumi.Input

	ToReplicatedVMManagedDiskOutput() ReplicatedVMManagedDiskOutput
	ToReplicatedVMManagedDiskOutputWithContext(context.Context) ReplicatedVMManagedDiskOutput
}

type ReplicatedVMManagedDiskArgs struct {
	// Id of disk that should be replicated.
	DiskId pulumi.StringInput `pulumi:"diskId"`
	// Storage account that should be used for caching.
	StagingStorageAccountId pulumi.StringInput `pulumi:"stagingStorageAccountId"`
	// What type should the disk be when a failover is done.
	TargetDiskType pulumi.StringInput `pulumi:"targetDiskType"`
	// What type should the disk be that holds the replication data.
	TargetReplicaDiskType pulumi.StringInput `pulumi:"targetReplicaDiskType"`
	// Resource group disk should belong to when a failover is done.
	TargetResourceGroupId pulumi.StringInput `pulumi:"targetResourceGroupId"`
}

func (ReplicatedVMManagedDiskArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ReplicatedVMManagedDisk)(nil)).Elem()
}

func (i ReplicatedVMManagedDiskArgs) ToReplicatedVMManagedDiskOutput() ReplicatedVMManagedDiskOutput {
	return i.ToReplicatedVMManagedDiskOutputWithContext(context.Background())
}

func (i ReplicatedVMManagedDiskArgs) ToReplicatedVMManagedDiskOutputWithContext(ctx context.Context) ReplicatedVMManagedDiskOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicatedVMManagedDiskOutput)
}

// ReplicatedVMManagedDiskArrayInput is an input type that accepts ReplicatedVMManagedDiskArray and ReplicatedVMManagedDiskArrayOutput values.
// You can construct a concrete instance of `ReplicatedVMManagedDiskArrayInput` via:
//
// 		 ReplicatedVMManagedDiskArray{ ReplicatedVMManagedDiskArgs{...} }
//
type ReplicatedVMManagedDiskArrayInput interface {
	pulumi.Input

	ToReplicatedVMManagedDiskArrayOutput() ReplicatedVMManagedDiskArrayOutput
	ToReplicatedVMManagedDiskArrayOutputWithContext(context.Context) ReplicatedVMManagedDiskArrayOutput
}

type ReplicatedVMManagedDiskArray []ReplicatedVMManagedDiskInput

func (ReplicatedVMManagedDiskArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ReplicatedVMManagedDisk)(nil)).Elem()
}

func (i ReplicatedVMManagedDiskArray) ToReplicatedVMManagedDiskArrayOutput() ReplicatedVMManagedDiskArrayOutput {
	return i.ToReplicatedVMManagedDiskArrayOutputWithContext(context.Background())
}

func (i ReplicatedVMManagedDiskArray) ToReplicatedVMManagedDiskArrayOutputWithContext(ctx context.Context) ReplicatedVMManagedDiskArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicatedVMManagedDiskArrayOutput)
}

type ReplicatedVMManagedDiskOutput struct{ *pulumi.OutputState }

func (ReplicatedVMManagedDiskOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReplicatedVMManagedDisk)(nil)).Elem()
}

func (o ReplicatedVMManagedDiskOutput) ToReplicatedVMManagedDiskOutput() ReplicatedVMManagedDiskOutput {
	return o
}

func (o ReplicatedVMManagedDiskOutput) ToReplicatedVMManagedDiskOutputWithContext(ctx context.Context) ReplicatedVMManagedDiskOutput {
	return o
}

// Id of disk that should be replicated.
func (o ReplicatedVMManagedDiskOutput) DiskId() pulumi.StringOutput {
	return o.ApplyT(func(v ReplicatedVMManagedDisk) string { return v.DiskId }).(pulumi.StringOutput)
}

// Storage account that should be used for caching.
func (o ReplicatedVMManagedDiskOutput) StagingStorageAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v ReplicatedVMManagedDisk) string { return v.StagingStorageAccountId }).(pulumi.StringOutput)
}

// What type should the disk be when a failover is done.
func (o ReplicatedVMManagedDiskOutput) TargetDiskType() pulumi.StringOutput {
	return o.ApplyT(func(v ReplicatedVMManagedDisk) string { return v.TargetDiskType }).(pulumi.StringOutput)
}

// What type should the disk be that holds the replication data.
func (o ReplicatedVMManagedDiskOutput) TargetReplicaDiskType() pulumi.StringOutput {
	return o.ApplyT(func(v ReplicatedVMManagedDisk) string { return v.TargetReplicaDiskType }).(pulumi.StringOutput)
}

// Resource group disk should belong to when a failover is done.
func (o ReplicatedVMManagedDiskOutput) TargetResourceGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v ReplicatedVMManagedDisk) string { return v.TargetResourceGroupId }).(pulumi.StringOutput)
}

type ReplicatedVMManagedDiskArrayOutput struct{ *pulumi.OutputState }

func (ReplicatedVMManagedDiskArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ReplicatedVMManagedDisk)(nil)).Elem()
}

func (o ReplicatedVMManagedDiskArrayOutput) ToReplicatedVMManagedDiskArrayOutput() ReplicatedVMManagedDiskArrayOutput {
	return o
}

func (o ReplicatedVMManagedDiskArrayOutput) ToReplicatedVMManagedDiskArrayOutputWithContext(ctx context.Context) ReplicatedVMManagedDiskArrayOutput {
	return o
}

func (o ReplicatedVMManagedDiskArrayOutput) Index(i pulumi.IntInput) ReplicatedVMManagedDiskOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ReplicatedVMManagedDisk {
		return vs[0].([]ReplicatedVMManagedDisk)[vs[1].(int)]
	}).(ReplicatedVMManagedDiskOutput)
}

func init() {
	pulumi.RegisterOutputType(ReplicatedVMManagedDiskOutput{})
	pulumi.RegisterOutputType(ReplicatedVMManagedDiskArrayOutput{})
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iot

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Use this data source to access information about an existing IotHub Shared Access Policy
func LookupSharedAccessPolicy(ctx *pulumi.Context, args *LookupSharedAccessPolicyArgs, opts ...pulumi.InvokeOption) (*LookupSharedAccessPolicyResult, error) {
	var rv LookupSharedAccessPolicyResult
	err := ctx.Invoke("azure:iot/getSharedAccessPolicy:getSharedAccessPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSharedAccessPolicy.
type LookupSharedAccessPolicyArgs struct {
	// The name of the IoTHub to which this Shared Access Policy belongs.
	IothubName string `pulumi:"iothubName"`
	// Specifies the name of the IotHub Shared Access Policy resource.
	Name string `pulumi:"name"`
	// The name of the resource group under which the IotHub Shared Access Policy resource has to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A collection of values returned by getSharedAccessPolicy.
type LookupSharedAccessPolicyResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id         string `pulumi:"id"`
	IothubName string `pulumi:"iothubName"`
	Name       string `pulumi:"name"`
	// The primary connection string of the Shared Access Policy.
	PrimaryConnectionString string `pulumi:"primaryConnectionString"`
	// The primary key used to create the authentication token.
	PrimaryKey        string `pulumi:"primaryKey"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The secondary connection string of the Shared Access Policy.
	SecondaryConnectionString string `pulumi:"secondaryConnectionString"`
	// The secondary key used to create the authentication token.
	SecondaryKey string `pulumi:"secondaryKey"`
}

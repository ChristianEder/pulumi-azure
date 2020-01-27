// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apimanagement

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Uses this data source to access information about an API Version Set within an API Management Service.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/d/api_management_api_version_set.html.markdown.
func LookupApiVersionSet(ctx *pulumi.Context, args *GetApiVersionSetArgs) (*GetApiVersionSetResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["apiManagementName"] = args.ApiManagementName
		inputs["name"] = args.Name
		inputs["resourceGroupName"] = args.ResourceGroupName
	}
	outputs, err := ctx.Invoke("azure:apimanagement/getApiVersionSet:getApiVersionSet", inputs)
	if err != nil {
		return nil, err
	}
	return &GetApiVersionSetResult{
		ApiManagementName: outputs["apiManagementName"],
		Description: outputs["description"],
		DisplayName: outputs["displayName"],
		Name: outputs["name"],
		ResourceGroupName: outputs["resourceGroupName"],
		VersionHeaderName: outputs["versionHeaderName"],
		VersionQueryName: outputs["versionQueryName"],
		VersioningScheme: outputs["versioningScheme"],
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getApiVersionSet.
type GetApiVersionSetArgs struct {
	// The name of the API Management Service where the API Version Set exists.
	ApiManagementName interface{}
	// The name of the API Version Set.
	Name interface{}
	// The name of the Resource Group in which the parent API Management Service exists.
	ResourceGroupName interface{}
}

// A collection of values returned by getApiVersionSet.
type GetApiVersionSetResult struct {
	ApiManagementName interface{}
	// The description of API Version Set.
	Description interface{}
	// The display name of this API Version Set.
	DisplayName interface{}
	Name interface{}
	ResourceGroupName interface{}
	// The name of the Header which should be read from Inbound Requests which defines the API Version.
	VersionHeaderName interface{}
	// The name of the Query String which should be read from Inbound Requests which defines the API Version.
	VersionQueryName interface{}
	VersioningScheme interface{}
	// id is the provider-assigned unique ID for this managed resource.
	Id interface{}
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iot

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages an IotHub
// 
// > **NOTE:** Endpoints can be defined either directly on the `iot.IoTHub` resource, or using the `azurerm_iothub_endpoint_*` resources - but the two ways of defining the endpoints cannot be used together. If both are used against the same IoTHub, spurious changes will occur. Also, defining a `azurerm_iothub_endpoint_*` resource and another endpoint of a different type directly on the `iot.IoTHub` resource is not supported.
// 
// > **NOTE:** Routes can be defined either directly on the `iot.IoTHub` resource, or using the `iot.Route` resource - but the two cannot be used together. If both are used against the same IoTHub, spurious changes will occur.
// 
// > **NOTE:** Fallback route can be defined either directly on the `iot.IoTHub` resource, or using the `iot.FallbackRoute` resource - but the two cannot be used together. If both are used against the same IoTHub, spurious changes will occur.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/iothub.html.markdown.
type IoTHub struct {
	s *pulumi.ResourceState
}

// NewIoTHub registers a new resource with the given unique name, arguments, and options.
func NewIoTHub(ctx *pulumi.Context,
	name string, args *IoTHubArgs, opts ...pulumi.ResourceOpt) (*IoTHub, error) {
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.Sku == nil {
		return nil, errors.New("missing required argument 'Sku'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["endpoints"] = nil
		inputs["eventHubPartitionCount"] = nil
		inputs["eventHubRetentionInDays"] = nil
		inputs["fallbackRoute"] = nil
		inputs["fileUpload"] = nil
		inputs["ipFilterRules"] = nil
		inputs["location"] = nil
		inputs["name"] = nil
		inputs["resourceGroupName"] = nil
		inputs["routes"] = nil
		inputs["sku"] = nil
		inputs["tags"] = nil
	} else {
		inputs["endpoints"] = args.Endpoints
		inputs["eventHubPartitionCount"] = args.EventHubPartitionCount
		inputs["eventHubRetentionInDays"] = args.EventHubRetentionInDays
		inputs["fallbackRoute"] = args.FallbackRoute
		inputs["fileUpload"] = args.FileUpload
		inputs["ipFilterRules"] = args.IpFilterRules
		inputs["location"] = args.Location
		inputs["name"] = args.Name
		inputs["resourceGroupName"] = args.ResourceGroupName
		inputs["routes"] = args.Routes
		inputs["sku"] = args.Sku
		inputs["tags"] = args.Tags
	}
	inputs["eventHubEventsEndpoint"] = nil
	inputs["eventHubEventsPath"] = nil
	inputs["eventHubOperationsEndpoint"] = nil
	inputs["eventHubOperationsPath"] = nil
	inputs["hostname"] = nil
	inputs["sharedAccessPolicies"] = nil
	inputs["type"] = nil
	s, err := ctx.RegisterResource("azure:iot/ioTHub:IoTHub", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &IoTHub{s: s}, nil
}

// GetIoTHub gets an existing IoTHub resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIoTHub(ctx *pulumi.Context,
	name string, id pulumi.ID, state *IoTHubState, opts ...pulumi.ResourceOpt) (*IoTHub, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["endpoints"] = state.Endpoints
		inputs["eventHubEventsEndpoint"] = state.EventHubEventsEndpoint
		inputs["eventHubEventsPath"] = state.EventHubEventsPath
		inputs["eventHubOperationsEndpoint"] = state.EventHubOperationsEndpoint
		inputs["eventHubOperationsPath"] = state.EventHubOperationsPath
		inputs["eventHubPartitionCount"] = state.EventHubPartitionCount
		inputs["eventHubRetentionInDays"] = state.EventHubRetentionInDays
		inputs["fallbackRoute"] = state.FallbackRoute
		inputs["fileUpload"] = state.FileUpload
		inputs["hostname"] = state.Hostname
		inputs["ipFilterRules"] = state.IpFilterRules
		inputs["location"] = state.Location
		inputs["name"] = state.Name
		inputs["resourceGroupName"] = state.ResourceGroupName
		inputs["routes"] = state.Routes
		inputs["sharedAccessPolicies"] = state.SharedAccessPolicies
		inputs["sku"] = state.Sku
		inputs["tags"] = state.Tags
		inputs["type"] = state.Type
	}
	s, err := ctx.ReadResource("azure:iot/ioTHub:IoTHub", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &IoTHub{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *IoTHub) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *IoTHub) ID() pulumi.IDOutput {
	return r.s.ID()
}

// An `endpoint` block as defined below.
func (r *IoTHub) Endpoints() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["endpoints"])
}

// The EventHub compatible endpoint for events data
func (r *IoTHub) EventHubEventsEndpoint() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["eventHubEventsEndpoint"])
}

// The EventHub compatible path for events data
func (r *IoTHub) EventHubEventsPath() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["eventHubEventsPath"])
}

// The EventHub compatible endpoint for operational data
func (r *IoTHub) EventHubOperationsEndpoint() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["eventHubOperationsEndpoint"])
}

// The EventHub compatible path for operational data
func (r *IoTHub) EventHubOperationsPath() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["eventHubOperationsPath"])
}

// The number of device-to-cloud partitions used by backing event hubs. Must be between `2` and `128`.
func (r *IoTHub) EventHubPartitionCount() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["eventHubPartitionCount"])
}

// The event hub retention to use in days. Must be between `1` and `7`.
func (r *IoTHub) EventHubRetentionInDays() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["eventHubRetentionInDays"])
}

// A `fallbackRoute` block as defined below. If the fallback route is enabled, messages that don't match any of the supplied routes are automatically sent to this route. Defaults to messages/events.
func (r *IoTHub) FallbackRoute() pulumi.Output {
	return r.s.State["fallbackRoute"]
}

// A `fileUpload` block as defined below.
func (r *IoTHub) FileUpload() pulumi.Output {
	return r.s.State["fileUpload"]
}

// The hostname of the IotHub Resource.
func (r *IoTHub) Hostname() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["hostname"])
}

// One or more `ipFilterRule` blocks as defined below.
func (r *IoTHub) IpFilterRules() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["ipFilterRules"])
}

// Specifies the supported Azure location where the resource has to be createc. Changing this forces a new resource to be created.
func (r *IoTHub) Location() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["location"])
}

// Specifies the name of the IotHub resource. Changing this forces a new resource to be created.
func (r *IoTHub) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// The name of the resource group under which the IotHub resource has to be created. Changing this forces a new resource to be created.
func (r *IoTHub) ResourceGroupName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["resourceGroupName"])
}

// A `route` block as defined below.
func (r *IoTHub) Routes() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["routes"])
}

// One or more `sharedAccessPolicy` blocks as defined below.
func (r *IoTHub) SharedAccessPolicies() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["sharedAccessPolicies"])
}

// A `sku` block as defined below.
func (r *IoTHub) Sku() pulumi.Output {
	return r.s.State["sku"]
}

// A mapping of tags to assign to the resource.
func (r *IoTHub) Tags() pulumi.MapOutput {
	return (pulumi.MapOutput)(r.s.State["tags"])
}

func (r *IoTHub) Type() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["type"])
}

// Input properties used for looking up and filtering IoTHub resources.
type IoTHubState struct {
	// An `endpoint` block as defined below.
	Endpoints interface{}
	// The EventHub compatible endpoint for events data
	EventHubEventsEndpoint interface{}
	// The EventHub compatible path for events data
	EventHubEventsPath interface{}
	// The EventHub compatible endpoint for operational data
	EventHubOperationsEndpoint interface{}
	// The EventHub compatible path for operational data
	EventHubOperationsPath interface{}
	// The number of device-to-cloud partitions used by backing event hubs. Must be between `2` and `128`.
	EventHubPartitionCount interface{}
	// The event hub retention to use in days. Must be between `1` and `7`.
	EventHubRetentionInDays interface{}
	// A `fallbackRoute` block as defined below. If the fallback route is enabled, messages that don't match any of the supplied routes are automatically sent to this route. Defaults to messages/events.
	FallbackRoute interface{}
	// A `fileUpload` block as defined below.
	FileUpload interface{}
	// The hostname of the IotHub Resource.
	Hostname interface{}
	// One or more `ipFilterRule` blocks as defined below.
	IpFilterRules interface{}
	// Specifies the supported Azure location where the resource has to be createc. Changing this forces a new resource to be created.
	Location interface{}
	// Specifies the name of the IotHub resource. Changing this forces a new resource to be created.
	Name interface{}
	// The name of the resource group under which the IotHub resource has to be created. Changing this forces a new resource to be created.
	ResourceGroupName interface{}
	// A `route` block as defined below.
	Routes interface{}
	// One or more `sharedAccessPolicy` blocks as defined below.
	SharedAccessPolicies interface{}
	// A `sku` block as defined below.
	Sku interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
	Type interface{}
}

// The set of arguments for constructing a IoTHub resource.
type IoTHubArgs struct {
	// An `endpoint` block as defined below.
	Endpoints interface{}
	// The number of device-to-cloud partitions used by backing event hubs. Must be between `2` and `128`.
	EventHubPartitionCount interface{}
	// The event hub retention to use in days. Must be between `1` and `7`.
	EventHubRetentionInDays interface{}
	// A `fallbackRoute` block as defined below. If the fallback route is enabled, messages that don't match any of the supplied routes are automatically sent to this route. Defaults to messages/events.
	FallbackRoute interface{}
	// A `fileUpload` block as defined below.
	FileUpload interface{}
	// One or more `ipFilterRule` blocks as defined below.
	IpFilterRules interface{}
	// Specifies the supported Azure location where the resource has to be createc. Changing this forces a new resource to be created.
	Location interface{}
	// Specifies the name of the IotHub resource. Changing this forces a new resource to be created.
	Name interface{}
	// The name of the resource group under which the IotHub resource has to be created. Changing this forces a new resource to be created.
	ResourceGroupName interface{}
	// A `route` block as defined below.
	Routes interface{}
	// A `sku` block as defined below.
	Sku interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
}

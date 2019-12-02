// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iot

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages an IotHub ServiceBus Queue Endpoint
// 
// > **NOTE:** Endpoints can be defined either directly on the `iot.IoTHub` resource, or using the `azurerm_iothub_endpoint_*` resources - but the two ways of defining the endpoints cannot be used together. If both are used against the same IoTHub, spurious changes will occur. Also, defining a `azurerm_iothub_endpoint_*` resource and another endpoint of a different type directly on the `iot.IoTHub` resource is not supported.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/iothub_endpoint_servicebus_queue.html.markdown.
type EndpointServicebusQueue struct {
	s *pulumi.ResourceState
}

// NewEndpointServicebusQueue registers a new resource with the given unique name, arguments, and options.
func NewEndpointServicebusQueue(ctx *pulumi.Context,
	name string, args *EndpointServicebusQueueArgs, opts ...pulumi.ResourceOpt) (*EndpointServicebusQueue, error) {
	if args == nil || args.ConnectionString == nil {
		return nil, errors.New("missing required argument 'ConnectionString'")
	}
	if args == nil || args.IothubName == nil {
		return nil, errors.New("missing required argument 'IothubName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["connectionString"] = nil
		inputs["iothubName"] = nil
		inputs["name"] = nil
		inputs["resourceGroupName"] = nil
	} else {
		inputs["connectionString"] = args.ConnectionString
		inputs["iothubName"] = args.IothubName
		inputs["name"] = args.Name
		inputs["resourceGroupName"] = args.ResourceGroupName
	}
	s, err := ctx.RegisterResource("azure:iot/endpointServicebusQueue:EndpointServicebusQueue", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &EndpointServicebusQueue{s: s}, nil
}

// GetEndpointServicebusQueue gets an existing EndpointServicebusQueue resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEndpointServicebusQueue(ctx *pulumi.Context,
	name string, id pulumi.ID, state *EndpointServicebusQueueState, opts ...pulumi.ResourceOpt) (*EndpointServicebusQueue, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["connectionString"] = state.ConnectionString
		inputs["iothubName"] = state.IothubName
		inputs["name"] = state.Name
		inputs["resourceGroupName"] = state.ResourceGroupName
	}
	s, err := ctx.ReadResource("azure:iot/endpointServicebusQueue:EndpointServicebusQueue", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &EndpointServicebusQueue{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *EndpointServicebusQueue) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *EndpointServicebusQueue) ID() pulumi.IDOutput {
	return r.s.ID()
}

// The connection string for the endpoint.
func (r *EndpointServicebusQueue) ConnectionString() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["connectionString"])
}

func (r *EndpointServicebusQueue) IothubName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["iothubName"])
}

// The name of the endpoint. The name must be unique across endpoint types. The following names are reserved:  `events`, `operationsMonitoringEvents`, `fileNotifications` and `$default`.
func (r *EndpointServicebusQueue) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

func (r *EndpointServicebusQueue) ResourceGroupName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["resourceGroupName"])
}

// Input properties used for looking up and filtering EndpointServicebusQueue resources.
type EndpointServicebusQueueState struct {
	// The connection string for the endpoint.
	ConnectionString interface{}
	IothubName interface{}
	// The name of the endpoint. The name must be unique across endpoint types. The following names are reserved:  `events`, `operationsMonitoringEvents`, `fileNotifications` and `$default`.
	Name interface{}
	ResourceGroupName interface{}
}

// The set of arguments for constructing a EndpointServicebusQueue resource.
type EndpointServicebusQueueArgs struct {
	// The connection string for the endpoint.
	ConnectionString interface{}
	IothubName interface{}
	// The name of the endpoint. The name must be unique across endpoint types. The following names are reserved:  `events`, `operationsMonitoringEvents`, `fileNotifications` and `$default`.
	Name interface{}
	ResourceGroupName interface{}
}
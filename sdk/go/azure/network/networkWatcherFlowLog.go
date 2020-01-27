// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a Network Watcher Flow Log.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/network_watcher_flow_log.html.markdown.
type NetworkWatcherFlowLog struct {
	s *pulumi.ResourceState
}

// NewNetworkWatcherFlowLog registers a new resource with the given unique name, arguments, and options.
func NewNetworkWatcherFlowLog(ctx *pulumi.Context,
	name string, args *NetworkWatcherFlowLogArgs, opts ...pulumi.ResourceOpt) (*NetworkWatcherFlowLog, error) {
	if args == nil || args.Enabled == nil {
		return nil, errors.New("missing required argument 'Enabled'")
	}
	if args == nil || args.NetworkSecurityGroupId == nil {
		return nil, errors.New("missing required argument 'NetworkSecurityGroupId'")
	}
	if args == nil || args.NetworkWatcherName == nil {
		return nil, errors.New("missing required argument 'NetworkWatcherName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.RetentionPolicy == nil {
		return nil, errors.New("missing required argument 'RetentionPolicy'")
	}
	if args == nil || args.StorageAccountId == nil {
		return nil, errors.New("missing required argument 'StorageAccountId'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["enabled"] = nil
		inputs["networkSecurityGroupId"] = nil
		inputs["networkWatcherName"] = nil
		inputs["resourceGroupName"] = nil
		inputs["retentionPolicy"] = nil
		inputs["storageAccountId"] = nil
		inputs["trafficAnalytics"] = nil
		inputs["version"] = nil
	} else {
		inputs["enabled"] = args.Enabled
		inputs["networkSecurityGroupId"] = args.NetworkSecurityGroupId
		inputs["networkWatcherName"] = args.NetworkWatcherName
		inputs["resourceGroupName"] = args.ResourceGroupName
		inputs["retentionPolicy"] = args.RetentionPolicy
		inputs["storageAccountId"] = args.StorageAccountId
		inputs["trafficAnalytics"] = args.TrafficAnalytics
		inputs["version"] = args.Version
	}
	s, err := ctx.RegisterResource("azure:network/networkWatcherFlowLog:NetworkWatcherFlowLog", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &NetworkWatcherFlowLog{s: s}, nil
}

// GetNetworkWatcherFlowLog gets an existing NetworkWatcherFlowLog resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkWatcherFlowLog(ctx *pulumi.Context,
	name string, id pulumi.ID, state *NetworkWatcherFlowLogState, opts ...pulumi.ResourceOpt) (*NetworkWatcherFlowLog, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["enabled"] = state.Enabled
		inputs["networkSecurityGroupId"] = state.NetworkSecurityGroupId
		inputs["networkWatcherName"] = state.NetworkWatcherName
		inputs["resourceGroupName"] = state.ResourceGroupName
		inputs["retentionPolicy"] = state.RetentionPolicy
		inputs["storageAccountId"] = state.StorageAccountId
		inputs["trafficAnalytics"] = state.TrafficAnalytics
		inputs["version"] = state.Version
	}
	s, err := ctx.ReadResource("azure:network/networkWatcherFlowLog:NetworkWatcherFlowLog", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &NetworkWatcherFlowLog{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *NetworkWatcherFlowLog) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *NetworkWatcherFlowLog) ID() pulumi.IDOutput {
	return r.s.ID()
}

// Boolean flag to enable/disable traffic analytics.
func (r *NetworkWatcherFlowLog) Enabled() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["enabled"])
}

// The ID of the Network Security Group for which to enable flow logs for. Changing this forces a new resource to be created.
func (r *NetworkWatcherFlowLog) NetworkSecurityGroupId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["networkSecurityGroupId"])
}

// The name of the Network Watcher. Changing this forces a new resource to be created.
func (r *NetworkWatcherFlowLog) NetworkWatcherName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["networkWatcherName"])
}

// The name of the resource group in which the Network Watcher was deployed. Changing this forces a new resource to be created.
func (r *NetworkWatcherFlowLog) ResourceGroupName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["resourceGroupName"])
}

// A `retentionPolicy` block as documented below.
func (r *NetworkWatcherFlowLog) RetentionPolicy() pulumi.Output {
	return r.s.State["retentionPolicy"]
}

// The ID of the Storage Account where flow logs are stored.
func (r *NetworkWatcherFlowLog) StorageAccountId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["storageAccountId"])
}

// A `trafficAnalytics` block as documented below.
func (r *NetworkWatcherFlowLog) TrafficAnalytics() pulumi.Output {
	return r.s.State["trafficAnalytics"]
}

// The version (revision) of the flow log. Possible values are `1` and `2`.
func (r *NetworkWatcherFlowLog) Version() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["version"])
}

// Input properties used for looking up and filtering NetworkWatcherFlowLog resources.
type NetworkWatcherFlowLogState struct {
	// Boolean flag to enable/disable traffic analytics.
	Enabled interface{}
	// The ID of the Network Security Group for which to enable flow logs for. Changing this forces a new resource to be created.
	NetworkSecurityGroupId interface{}
	// The name of the Network Watcher. Changing this forces a new resource to be created.
	NetworkWatcherName interface{}
	// The name of the resource group in which the Network Watcher was deployed. Changing this forces a new resource to be created.
	ResourceGroupName interface{}
	// A `retentionPolicy` block as documented below.
	RetentionPolicy interface{}
	// The ID of the Storage Account where flow logs are stored.
	StorageAccountId interface{}
	// A `trafficAnalytics` block as documented below.
	TrafficAnalytics interface{}
	// The version (revision) of the flow log. Possible values are `1` and `2`.
	Version interface{}
}

// The set of arguments for constructing a NetworkWatcherFlowLog resource.
type NetworkWatcherFlowLogArgs struct {
	// Boolean flag to enable/disable traffic analytics.
	Enabled interface{}
	// The ID of the Network Security Group for which to enable flow logs for. Changing this forces a new resource to be created.
	NetworkSecurityGroupId interface{}
	// The name of the Network Watcher. Changing this forces a new resource to be created.
	NetworkWatcherName interface{}
	// The name of the resource group in which the Network Watcher was deployed. Changing this forces a new resource to be created.
	ResourceGroupName interface{}
	// A `retentionPolicy` block as documented below.
	RetentionPolicy interface{}
	// The ID of the Storage Account where flow logs are stored.
	StorageAccountId interface{}
	// A `trafficAnalytics` block as documented below.
	TrafficAnalytics interface{}
	// The version (revision) of the flow log. Possible values are `1` and `2`.
	Version interface{}
}

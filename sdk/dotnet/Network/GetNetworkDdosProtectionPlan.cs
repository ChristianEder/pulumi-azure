// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Network
{
    public static class GetNetworkDdosProtectionPlan
    {
        /// <summary>
        /// Use this data source to access information about an existing Azure Network DDoS Protection Plan.
        /// 
        /// {{% examples %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetNetworkDdosProtectionPlanResult> InvokeAsync(GetNetworkDdosProtectionPlanArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetNetworkDdosProtectionPlanResult>("azure:network/getNetworkDdosProtectionPlan:getNetworkDdosProtectionPlan", args ?? new GetNetworkDdosProtectionPlanArgs(), options.WithVersion());
    }


    public sealed class GetNetworkDdosProtectionPlanArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Network DDoS Protection Plan.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The name of the resource group where the Network DDoS Protection Plan exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        [Input("tags")]
        private Dictionary<string, string>? _tags;

        /// <summary>
        /// A mapping of tags assigned to the resource.
        /// </summary>
        public Dictionary<string, string> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, string>());
            set => _tags = value;
        }

        public GetNetworkDdosProtectionPlanArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetNetworkDdosProtectionPlanResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Specifies the supported Azure location where the resource exists.
        /// </summary>
        public readonly string Location;
        public readonly string Name;
        public readonly string ResourceGroupName;
        /// <summary>
        /// A mapping of tags assigned to the resource.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// The Resource ID list of the Virtual Networks associated with DDoS Protection Plan.
        /// </summary>
        public readonly ImmutableArray<string> VirtualNetworkIds;

        [OutputConstructor]
        private GetNetworkDdosProtectionPlanResult(
            string id,

            string location,

            string name,

            string resourceGroupName,

            ImmutableDictionary<string, string>? tags,

            ImmutableArray<string> virtualNetworkIds)
        {
            Id = id;
            Location = location;
            Name = name;
            ResourceGroupName = resourceGroupName;
            Tags = tags;
            VirtualNetworkIds = virtualNetworkIds;
        }
    }
}

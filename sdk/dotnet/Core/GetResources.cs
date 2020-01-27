// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Core
{
    public static partial class Invokes
    {
        /// <summary>
        /// Use this data source to access information about existing resources.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/d/resources.html.markdown.
        /// </summary>
        public static Task<GetResourcesResult> GetResources(GetResourcesArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetResourcesResult>("azure:core/getResources:getResources", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetResourcesArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Resource.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        [Input("requiredTags")]
        private Dictionary<string, string>? _requiredTags;

        /// <summary>
        /// A mapping of tags which the resource has to have in order to be included in the result.
        /// </summary>
        public Dictionary<string, string> RequiredTags
        {
            get => _requiredTags ?? (_requiredTags = new Dictionary<string, string>());
            set => _requiredTags = value;
        }

        /// <summary>
        /// The name of the Resource group where the Resources are located.
        /// </summary>
        [Input("resourceGroupName")]
        public string? ResourceGroupName { get; set; }

        /// <summary>
        /// The Resource Type of the Resources you want to list (e.g. `Microsoft.Network/virtualNetworks`). A full list of available Resource Types can be found [here](https://docs.microsoft.com/en-us/azure/azure-resource-manager/azure-services-resource-providers).
        /// </summary>
        [Input("type")]
        public string? Type { get; set; }

        public GetResourcesArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetResourcesResult
    {
        /// <summary>
        /// The name of this Resource.
        /// </summary>
        public readonly string Name;
        public readonly ImmutableDictionary<string, string> RequiredTags;
        public readonly string ResourceGroupName;
        /// <summary>
        /// One or more `resource` blocks as defined below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetResourcesResourcesResult> Resources;
        /// <summary>
        /// The type of this Resource. (e.g. `Microsoft.Network/virtualNetworks`).
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetResourcesResult(
            string name,
            ImmutableDictionary<string, string> requiredTags,
            string resourceGroupName,
            ImmutableArray<Outputs.GetResourcesResourcesResult> resources,
            string type,
            string id)
        {
            Name = name;
            RequiredTags = requiredTags;
            ResourceGroupName = resourceGroupName;
            Resources = resources;
            Type = type;
            Id = id;
        }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class GetResourcesResourcesResult
    {
        /// <summary>
        /// The ID of this Resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The Azure Region in which this Resource exists.
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// The name of the Resource.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// A map of tags assigned to this Resource.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Tags;
        /// <summary>
        /// The Resource Type of the Resources you want to list (e.g. `Microsoft.Network/virtualNetworks`). A full list of available Resource Types can be found [here](https://docs.microsoft.com/en-us/azure/azure-resource-manager/azure-services-resource-providers).
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetResourcesResourcesResult(
            string id,
            string location,
            string name,
            ImmutableDictionary<string, string> tags,
            string type)
        {
            Id = id;
            Location = location;
            Name = name;
            Tags = tags;
            Type = type;
        }
    }
    }
}

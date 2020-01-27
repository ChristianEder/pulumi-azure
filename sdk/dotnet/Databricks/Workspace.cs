// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.DataBricks
{
    /// <summary>
    /// Manages a Databricks Workspace
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/databricks_workspace.html.markdown.
    /// </summary>
    public partial class Workspace : Pulumi.CustomResource
    {
        /// <summary>
        /// A `custom_parameters` block as documented below.
        /// </summary>
        [Output("customParameters")]
        public Output<Outputs.WorkspaceCustomParameters> CustomParameters { get; private set; } = null!;

        /// <summary>
        /// Specifies the supported Azure location where the resource has to be created. Changing this forces a new resource to be created.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The ID of the Managed Resource Group created by the Databricks Workspace.
        /// </summary>
        [Output("managedResourceGroupId")]
        public Output<string> ManagedResourceGroupId { get; private set; } = null!;

        /// <summary>
        /// The name of the resource group where Azure should place the managed Databricks resources. Changing this forces a new resource to be created.
        /// </summary>
        [Output("managedResourceGroupName")]
        public Output<string> ManagedResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the Databricks Workspace resource. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The name of the Resource Group in which the Databricks Workspace should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// The `sku` to use for the Databricks Workspace. Possible values are `standard` or `premium`. Changing this forces a new resource to be created.
        /// </summary>
        [Output("sku")]
        public Output<string> Sku { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a Workspace resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Workspace(string name, WorkspaceArgs args, CustomResourceOptions? options = null)
            : base("azure:databricks/workspace:Workspace", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private Workspace(string name, Input<string> id, WorkspaceState? state = null, CustomResourceOptions? options = null)
            : base("azure:databricks/workspace:Workspace", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Workspace resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Workspace Get(string name, Input<string> id, WorkspaceState? state = null, CustomResourceOptions? options = null)
        {
            return new Workspace(name, id, state, options);
        }
    }

    public sealed class WorkspaceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A `custom_parameters` block as documented below.
        /// </summary>
        [Input("customParameters")]
        public Input<Inputs.WorkspaceCustomParametersArgs>? CustomParameters { get; set; }

        /// <summary>
        /// Specifies the supported Azure location where the resource has to be created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name of the resource group where Azure should place the managed Databricks resources. Changing this forces a new resource to be created.
        /// </summary>
        [Input("managedResourceGroupName")]
        public Input<string>? ManagedResourceGroupName { get; set; }

        /// <summary>
        /// Specifies the name of the Databricks Workspace resource. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the Resource Group in which the Databricks Workspace should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The `sku` to use for the Databricks Workspace. Possible values are `standard` or `premium`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("sku", required: true)]
        public Input<string> Sku { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public WorkspaceArgs()
        {
        }
    }

    public sealed class WorkspaceState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A `custom_parameters` block as documented below.
        /// </summary>
        [Input("customParameters")]
        public Input<Inputs.WorkspaceCustomParametersGetArgs>? CustomParameters { get; set; }

        /// <summary>
        /// Specifies the supported Azure location where the resource has to be created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The ID of the Managed Resource Group created by the Databricks Workspace.
        /// </summary>
        [Input("managedResourceGroupId")]
        public Input<string>? ManagedResourceGroupId { get; set; }

        /// <summary>
        /// The name of the resource group where Azure should place the managed Databricks resources. Changing this forces a new resource to be created.
        /// </summary>
        [Input("managedResourceGroupName")]
        public Input<string>? ManagedResourceGroupName { get; set; }

        /// <summary>
        /// Specifies the name of the Databricks Workspace resource. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the Resource Group in which the Databricks Workspace should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        /// <summary>
        /// The `sku` to use for the Databricks Workspace. Possible values are `standard` or `premium`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("sku")]
        public Input<string>? Sku { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public WorkspaceState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class WorkspaceCustomParametersArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Are public IP Addresses not allowed?
        /// </summary>
        [Input("noPublicIp")]
        public Input<bool>? NoPublicIp { get; set; }

        /// <summary>
        /// The name of the Private Subnet within the Virtual Network. Required if `virtual_network_id` is set.
        /// </summary>
        [Input("privateSubnetName")]
        public Input<string>? PrivateSubnetName { get; set; }

        /// <summary>
        /// The name of the Public Subnet within the Virtual Network. Required if `virtual_network_id` is set.
        /// </summary>
        [Input("publicSubnetName")]
        public Input<string>? PublicSubnetName { get; set; }

        /// <summary>
        /// The ID of a Virtual Network where this Databricks Cluster should be created.
        /// </summary>
        [Input("virtualNetworkId")]
        public Input<string>? VirtualNetworkId { get; set; }

        public WorkspaceCustomParametersArgs()
        {
        }
    }

    public sealed class WorkspaceCustomParametersGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Are public IP Addresses not allowed?
        /// </summary>
        [Input("noPublicIp")]
        public Input<bool>? NoPublicIp { get; set; }

        /// <summary>
        /// The name of the Private Subnet within the Virtual Network. Required if `virtual_network_id` is set.
        /// </summary>
        [Input("privateSubnetName")]
        public Input<string>? PrivateSubnetName { get; set; }

        /// <summary>
        /// The name of the Public Subnet within the Virtual Network. Required if `virtual_network_id` is set.
        /// </summary>
        [Input("publicSubnetName")]
        public Input<string>? PublicSubnetName { get; set; }

        /// <summary>
        /// The ID of a Virtual Network where this Databricks Cluster should be created.
        /// </summary>
        [Input("virtualNetworkId")]
        public Input<string>? VirtualNetworkId { get; set; }

        public WorkspaceCustomParametersGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class WorkspaceCustomParameters
    {
        /// <summary>
        /// Are public IP Addresses not allowed?
        /// </summary>
        public readonly bool? NoPublicIp;
        /// <summary>
        /// The name of the Private Subnet within the Virtual Network. Required if `virtual_network_id` is set.
        /// </summary>
        public readonly string? PrivateSubnetName;
        /// <summary>
        /// The name of the Public Subnet within the Virtual Network. Required if `virtual_network_id` is set.
        /// </summary>
        public readonly string? PublicSubnetName;
        /// <summary>
        /// The ID of a Virtual Network where this Databricks Cluster should be created.
        /// </summary>
        public readonly string? VirtualNetworkId;

        [OutputConstructor]
        private WorkspaceCustomParameters(
            bool? noPublicIp,
            string? privateSubnetName,
            string? publicSubnetName,
            string? virtualNetworkId)
        {
            NoPublicIp = noPublicIp;
            PrivateSubnetName = privateSubnetName;
            PublicSubnetName = publicSubnetName;
            VirtualNetworkId = virtualNetworkId;
        }
    }
    }
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Relay
{
    /// <summary>
    /// Manages an Azure Relay Namespace.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/relay_namespace.html.markdown.
    /// </summary>
    public partial class Namespace : Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies the supported Azure location where the Azure Relay Namespace exists. Changing this forces a new resource to be created.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The Identifier for Azure Insights metrics.
        /// </summary>
        [Output("metricId")]
        public Output<string> MetricId { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the Azure Relay Namespace. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The primary connection string for the authorization rule `RootManageSharedAccessKey`.
        /// </summary>
        [Output("primaryConnectionString")]
        public Output<string> PrimaryConnectionString { get; private set; } = null!;

        /// <summary>
        /// The primary access key for the authorization rule `RootManageSharedAccessKey`.
        /// </summary>
        [Output("primaryKey")]
        public Output<string> PrimaryKey { get; private set; } = null!;

        /// <summary>
        /// The name of the resource group in which to create the Azure Relay Namespace.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// The secondary connection string for the authorization rule `RootManageSharedAccessKey`.
        /// </summary>
        [Output("secondaryConnectionString")]
        public Output<string> SecondaryConnectionString { get; private set; } = null!;

        /// <summary>
        /// The secondary access key for the authorization rule `RootManageSharedAccessKey`.
        /// </summary>
        [Output("secondaryKey")]
        public Output<string> SecondaryKey { get; private set; } = null!;

        /// <summary>
        /// ) A `sku` block as described below.
        /// </summary>
        [Output("sku")]
        public Output<Outputs.NamespaceSku> Sku { get; private set; } = null!;

        /// <summary>
        /// The name of the SKU to use. At this time the only supported value is `Standard`.
        /// </summary>
        [Output("skuName")]
        public Output<string> SkuName { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a Namespace resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Namespace(string name, NamespaceArgs args, CustomResourceOptions? options = null)
            : base("azure:relay/namespace:Namespace", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private Namespace(string name, Input<string> id, NamespaceState? state = null, CustomResourceOptions? options = null)
            : base("azure:relay/namespace:Namespace", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Namespace resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Namespace Get(string name, Input<string> id, NamespaceState? state = null, CustomResourceOptions? options = null)
        {
            return new Namespace(name, id, state, options);
        }
    }

    public sealed class NamespaceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the supported Azure location where the Azure Relay Namespace exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Specifies the name of the Azure Relay Namespace. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the resource group in which to create the Azure Relay Namespace.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// ) A `sku` block as described below.
        /// </summary>
        [Input("sku")]
        public Input<Inputs.NamespaceSkuArgs>? Sku { get; set; }

        /// <summary>
        /// The name of the SKU to use. At this time the only supported value is `Standard`.
        /// </summary>
        [Input("skuName")]
        public Input<string>? SkuName { get; set; }

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

        public NamespaceArgs()
        {
        }
    }

    public sealed class NamespaceState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the supported Azure location where the Azure Relay Namespace exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The Identifier for Azure Insights metrics.
        /// </summary>
        [Input("metricId")]
        public Input<string>? MetricId { get; set; }

        /// <summary>
        /// Specifies the name of the Azure Relay Namespace. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The primary connection string for the authorization rule `RootManageSharedAccessKey`.
        /// </summary>
        [Input("primaryConnectionString")]
        public Input<string>? PrimaryConnectionString { get; set; }

        /// <summary>
        /// The primary access key for the authorization rule `RootManageSharedAccessKey`.
        /// </summary>
        [Input("primaryKey")]
        public Input<string>? PrimaryKey { get; set; }

        /// <summary>
        /// The name of the resource group in which to create the Azure Relay Namespace.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        /// <summary>
        /// The secondary connection string for the authorization rule `RootManageSharedAccessKey`.
        /// </summary>
        [Input("secondaryConnectionString")]
        public Input<string>? SecondaryConnectionString { get; set; }

        /// <summary>
        /// The secondary access key for the authorization rule `RootManageSharedAccessKey`.
        /// </summary>
        [Input("secondaryKey")]
        public Input<string>? SecondaryKey { get; set; }

        /// <summary>
        /// ) A `sku` block as described below.
        /// </summary>
        [Input("sku")]
        public Input<Inputs.NamespaceSkuGetArgs>? Sku { get; set; }

        /// <summary>
        /// The name of the SKU to use. At this time the only supported value is `Standard`.
        /// </summary>
        [Input("skuName")]
        public Input<string>? SkuName { get; set; }

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

        public NamespaceState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class NamespaceSkuArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the name of the Azure Relay Namespace. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public NamespaceSkuArgs()
        {
        }
    }

    public sealed class NamespaceSkuGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the name of the Azure Relay Namespace. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public NamespaceSkuGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class NamespaceSku
    {
        /// <summary>
        /// Specifies the name of the Azure Relay Namespace. Changing this forces a new resource to be created.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private NamespaceSku(string name)
        {
            Name = name;
        }
    }
    }
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Network
{
    /// <summary>
    /// Manages a Public IP Prefix.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/public_ip_prefix.html.markdown.
    /// </summary>
    public partial class PublicIpPrefix : Pulumi.CustomResource
    {
        /// <summary>
        /// The IP address prefix value that was allocated.
        /// </summary>
        [Output("ipPrefix")]
        public Output<string> IpPrefix { get; private set; } = null!;

        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the Public IP resource . Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Specifies the number of bits of the prefix. The value can be set between 24 (256 addresses) and 31 (2 addresses). Changing this forces a new resource to be created.
        /// </summary>
        [Output("prefixLength")]
        public Output<int?> PrefixLength { get; private set; } = null!;

        /// <summary>
        /// The name of the resource group in which to create the public IP.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// The SKU of the Public IP Prefix. Accepted values are `Standard`. Defaults to `Standard`. Changing this forces a new resource to be created.
        /// </summary>
        [Output("sku")]
        public Output<string?> Sku { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>> Tags { get; private set; } = null!;

        /// <summary>
        /// A collection containing the availability zone to allocate the Public IP in.
        /// </summary>
        [Output("zones")]
        public Output<string?> Zones { get; private set; } = null!;


        /// <summary>
        /// Create a PublicIpPrefix resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PublicIpPrefix(string name, PublicIpPrefixArgs args, CustomResourceOptions? options = null)
            : base("azure:network/publicIpPrefix:PublicIpPrefix", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private PublicIpPrefix(string name, Input<string> id, PublicIpPrefixState? state = null, CustomResourceOptions? options = null)
            : base("azure:network/publicIpPrefix:PublicIpPrefix", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing PublicIpPrefix resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PublicIpPrefix Get(string name, Input<string> id, PublicIpPrefixState? state = null, CustomResourceOptions? options = null)
        {
            return new PublicIpPrefix(name, id, state, options);
        }
    }

    public sealed class PublicIpPrefixArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Specifies the name of the Public IP resource . Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the number of bits of the prefix. The value can be set between 24 (256 addresses) and 31 (2 addresses). Changing this forces a new resource to be created.
        /// </summary>
        [Input("prefixLength")]
        public Input<int>? PrefixLength { get; set; }

        /// <summary>
        /// The name of the resource group in which to create the public IP.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The SKU of the Public IP Prefix. Accepted values are `Standard`. Defaults to `Standard`. Changing this forces a new resource to be created.
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

        /// <summary>
        /// A collection containing the availability zone to allocate the Public IP in.
        /// </summary>
        [Input("zones")]
        public Input<string>? Zones { get; set; }

        public PublicIpPrefixArgs()
        {
        }
    }

    public sealed class PublicIpPrefixState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The IP address prefix value that was allocated.
        /// </summary>
        [Input("ipPrefix")]
        public Input<string>? IpPrefix { get; set; }

        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Specifies the name of the Public IP resource . Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the number of bits of the prefix. The value can be set between 24 (256 addresses) and 31 (2 addresses). Changing this forces a new resource to be created.
        /// </summary>
        [Input("prefixLength")]
        public Input<int>? PrefixLength { get; set; }

        /// <summary>
        /// The name of the resource group in which to create the public IP.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        /// <summary>
        /// The SKU of the Public IP Prefix. Accepted values are `Standard`. Defaults to `Standard`. Changing this forces a new resource to be created.
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

        /// <summary>
        /// A collection containing the availability zone to allocate the Public IP in.
        /// </summary>
        [Input("zones")]
        public Input<string>? Zones { get; set; }

        public PublicIpPrefixState()
        {
        }
    }
}

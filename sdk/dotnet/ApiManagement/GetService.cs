// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.ApiManagement
{
    public static class GetService
    {
        /// <summary>
        /// Use this data source to access information about an existing API Management Service.
        /// 
        /// {{% examples %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetServiceResult> InvokeAsync(GetServiceArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetServiceResult>("azure:apimanagement/getService:getService", args ?? new GetServiceArgs(), options.WithVersion());
    }


    public sealed class GetServiceArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the API Management service.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The Name of the Resource Group in which the API Management Service exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetServiceArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetServiceResult
    {
        /// <summary>
        /// One or more `additional_location` blocks as defined below
        /// </summary>
        public readonly ImmutableArray<Outputs.GetServiceAdditionalLocationResult> AdditionalLocations;
        /// <summary>
        /// Gateway URL of the API Management service in the Region.
        /// </summary>
        public readonly string GatewayRegionalUrl;
        /// <summary>
        /// The URL for the API Management Service's Gateway.
        /// </summary>
        public readonly string GatewayUrl;
        /// <summary>
        /// A `hostname_configuration` block as defined below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetServiceHostnameConfigurationResult> HostnameConfigurations;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The location name of the additional region among Azure Data center regions.
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// The URL for the Management API.
        /// </summary>
        public readonly string ManagementApiUrl;
        /// <summary>
        /// Specifies the plan's pricing tier.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The email address from which the notification will be sent.
        /// </summary>
        public readonly string NotificationSenderEmail;
        /// <summary>
        /// The URL of the Publisher Portal.
        /// </summary>
        public readonly string PortalUrl;
        /// <summary>
        /// Public Static Load Balanced IP addresses of the API Management service in the additional location. Available only for Basic, Standard and Premium SKU.
        /// </summary>
        public readonly ImmutableArray<string> PublicIpAddresses;
        /// <summary>
        /// The email of Publisher/Company of the API Management Service.
        /// </summary>
        public readonly string PublisherEmail;
        /// <summary>
        /// The name of the Publisher/Company of the API Management Service.
        /// </summary>
        public readonly string PublisherName;
        public readonly string ResourceGroupName;
        /// <summary>
        /// The SCM (Source Code Management) endpoint.
        /// </summary>
        public readonly string ScmUrl;
        public readonly string SkuName;
        /// <summary>
        /// A mapping of tags assigned to the resource.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Tags;

        [OutputConstructor]
        private GetServiceResult(
            ImmutableArray<Outputs.GetServiceAdditionalLocationResult> additionalLocations,

            string gatewayRegionalUrl,

            string gatewayUrl,

            ImmutableArray<Outputs.GetServiceHostnameConfigurationResult> hostnameConfigurations,

            string id,

            string location,

            string managementApiUrl,

            string name,

            string notificationSenderEmail,

            string portalUrl,

            ImmutableArray<string> publicIpAddresses,

            string publisherEmail,

            string publisherName,

            string resourceGroupName,

            string scmUrl,

            string skuName,

            ImmutableDictionary<string, string> tags)
        {
            AdditionalLocations = additionalLocations;
            GatewayRegionalUrl = gatewayRegionalUrl;
            GatewayUrl = gatewayUrl;
            HostnameConfigurations = hostnameConfigurations;
            Id = id;
            Location = location;
            ManagementApiUrl = managementApiUrl;
            Name = name;
            NotificationSenderEmail = notificationSenderEmail;
            PortalUrl = portalUrl;
            PublicIpAddresses = publicIpAddresses;
            PublisherEmail = publisherEmail;
            PublisherName = publisherName;
            ResourceGroupName = resourceGroupName;
            ScmUrl = scmUrl;
            SkuName = skuName;
            Tags = tags;
        }
    }
}

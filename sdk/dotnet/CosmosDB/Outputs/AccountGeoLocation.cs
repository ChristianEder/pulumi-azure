// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.CosmosDB.Outputs
{

    [OutputType]
    public sealed class AccountGeoLocation
    {
        /// <summary>
        /// The failover priority of the region. A failover priority of `0` indicates a write region. The maximum value for a failover priority = (total number of regions - 1). Failover priority values must be unique for each of the regions in which the database account exists. Changing this causes the location to be re-provisioned and cannot be changed for the location with failover priority `0`.
        /// </summary>
        public readonly int FailoverPriority;
        /// <summary>
        /// The ID of the virtual network subnet.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// The name of the Azure region to host replicated data.
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// The string used to generate the document endpoints for this region. If not specified it defaults to `${cosmosdb_account.name}-${location}`. Changing this causes the location to be deleted and re-provisioned and cannot be changed for the location with failover priority `0`.
        /// </summary>
        public readonly string? Prefix;

        [OutputConstructor]
        private AccountGeoLocation(
            int failoverPriority,

            string? id,

            string location,

            string? prefix)
        {
            FailoverPriority = failoverPriority;
            Id = id;
            Location = location;
            Prefix = prefix;
        }
    }
}
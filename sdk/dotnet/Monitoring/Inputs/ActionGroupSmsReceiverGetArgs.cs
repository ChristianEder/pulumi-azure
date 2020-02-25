// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Monitoring.Inputs
{

    public sealed class ActionGroupSmsReceiverGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The country code of the SMS receiver.
        /// </summary>
        [Input("countryCode", required: true)]
        public Input<string> CountryCode { get; set; } = null!;

        /// <summary>
        /// The name of the SMS receiver. Names must be unique (case-insensitive) across all receivers within an action group.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The phone number of the SMS receiver.
        /// </summary>
        [Input("phoneNumber", required: true)]
        public Input<string> PhoneNumber { get; set; } = null!;

        public ActionGroupSmsReceiverGetArgs()
        {
        }
    }
}
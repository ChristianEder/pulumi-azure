// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.DevTest.Outputs
{

    [OutputType]
    public sealed class ScheduleDailyRecurrence
    {
        /// <summary>
        /// The time each day when the schedule takes effect.
        /// </summary>
        public readonly string Time;

        [OutputConstructor]
        private ScheduleDailyRecurrence(string time)
        {
            Time = time;
        }
    }
}
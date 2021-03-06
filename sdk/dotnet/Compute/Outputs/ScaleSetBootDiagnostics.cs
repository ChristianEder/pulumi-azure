// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Compute.Outputs
{

    [OutputType]
    public sealed class ScaleSetBootDiagnostics
    {
        public readonly bool? Enabled;
        public readonly string StorageUri;

        [OutputConstructor]
        private ScaleSetBootDiagnostics(
            bool? enabled,

            string storageUri)
        {
            Enabled = enabled;
            StorageUri = storageUri;
        }
    }
}

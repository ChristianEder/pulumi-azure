// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Manages a Data Lake Gen2 File System within an Azure Storage Account.
 * 
 * > **NOTE:** This Resource requires using Azure Active Directory to connect to Azure Storage, which in turn requires the `Storage` specific roles - which are not granted by default.
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/storage_data_lake_gen2_filesystem.html.markdown.
 */
export class DataLakeGen2Filesystem extends pulumi.CustomResource {
    /**
     * Get an existing DataLakeGen2Filesystem resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DataLakeGen2FilesystemState, opts?: pulumi.CustomResourceOptions): DataLakeGen2Filesystem {
        return new DataLakeGen2Filesystem(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:storage/dataLakeGen2Filesystem:DataLakeGen2Filesystem';

    /**
     * Returns true if the given object is an instance of DataLakeGen2Filesystem.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DataLakeGen2Filesystem {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DataLakeGen2Filesystem.__pulumiType;
    }

    /**
     * The name of the Data Lake Gen2 File System which should be created within the Storage Account. Must be unique within the storage account the queue is located. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * A mapping of Key to Base64-Encoded Values which should be assigned to this Data Lake Gen2 File System. 
     */
    public readonly properties!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Specifies the ID of the Storage Account in which the Data Lake Gen2 File System should exist. Changing this forces a new resource to be created.
     */
    public readonly storageAccountId!: pulumi.Output<string>;

    /**
     * Create a DataLakeGen2Filesystem resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DataLakeGen2FilesystemArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DataLakeGen2FilesystemArgs | DataLakeGen2FilesystemState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as DataLakeGen2FilesystemState | undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["properties"] = state ? state.properties : undefined;
            inputs["storageAccountId"] = state ? state.storageAccountId : undefined;
        } else {
            const args = argsOrState as DataLakeGen2FilesystemArgs | undefined;
            if (!args || args.storageAccountId === undefined) {
                throw new Error("Missing required property 'storageAccountId'");
            }
            inputs["name"] = args ? args.name : undefined;
            inputs["properties"] = args ? args.properties : undefined;
            inputs["storageAccountId"] = args ? args.storageAccountId : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(DataLakeGen2Filesystem.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DataLakeGen2Filesystem resources.
 */
export interface DataLakeGen2FilesystemState {
    /**
     * The name of the Data Lake Gen2 File System which should be created within the Storage Account. Must be unique within the storage account the queue is located. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A mapping of Key to Base64-Encoded Values which should be assigned to this Data Lake Gen2 File System. 
     */
    readonly properties?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Specifies the ID of the Storage Account in which the Data Lake Gen2 File System should exist. Changing this forces a new resource to be created.
     */
    readonly storageAccountId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DataLakeGen2Filesystem resource.
 */
export interface DataLakeGen2FilesystemArgs {
    /**
     * The name of the Data Lake Gen2 File System which should be created within the Storage Account. Must be unique within the storage account the queue is located. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A mapping of Key to Base64-Encoded Values which should be assigned to this Data Lake Gen2 File System. 
     */
    readonly properties?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Specifies the ID of the Storage Account in which the Data Lake Gen2 File System should exist. Changing this forces a new resource to be created.
     */
    readonly storageAccountId: pulumi.Input<string>;
}

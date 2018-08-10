// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";

/**
 * Gets information about the specified Storage Account.
 */
export function getAccount(args: GetAccountArgs, opts?: pulumi.InvokeOptions): Promise<GetAccountResult> {
    return pulumi.runtime.invoke("azure:storage/getAccount:getAccount", {
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

/**
 * A collection of arguments for invoking getAccount.
 */
export interface GetAccountArgs {
    /**
     * Specifies the name of the Storage Account
     */
    readonly name: string;
    /**
     * Specifies the name of the resource group the Storage Account is located in.
     */
    readonly resourceGroupName: string;
}

/**
 * A collection of values returned by getAccount.
 */
export interface GetAccountResult {
    /**
     * The access tier for `BlobStorage` accounts.
     */
    readonly accessTier: string;
    /**
     * The Encryption Source for this Storage Account.
     */
    readonly accountEncryptionSource: string;
    /**
     * The Kind of account.
     */
    readonly accountKind: string;
    /**
     * The type of replication used for this storage account.
     */
    readonly accountReplicationType: string;
    /**
     * The Tier of this storage account.
     */
    readonly accountTier: string;
    /**
     * A `custom_domain` block as documented below.
     */
    readonly customDomain: { name: string };
    /**
     * Are Encryption Services are enabled for Blob storage? See [here](https://azure.microsoft.com/en-us/documentation/articles/storage-service-encryption/)
     * for more information.
     */
    readonly enableBlobEncryption: boolean;
    /**
     * Are Encryption Services are enabled for File storage? See [here](https://azure.microsoft.com/en-us/documentation/articles/storage-service-encryption/)
     * for more information.
     */
    readonly enableFileEncryption: boolean;
    /**
     * Is traffic only allowed via HTTPS? See [here](https://docs.microsoft.com/en-us/azure/storage/storage-require-secure-transfer/)
     * for more information.
     */
    readonly enableHttpsTrafficOnly: boolean;
    /**
     * The Azure location where the Storage Account exists
     */
    readonly location: string;
    /**
     * The primary access key for the Storage Account.
     */
    readonly primaryAccessKey: string;
    /**
     * The connection string associated with the primary blob location
     */
    readonly primaryBlobConnectionString: string;
    /**
     * The endpoint URL for blob storage in the primary location.
     */
    readonly primaryBlobEndpoint: string;
    /**
     * The connection string associated with the primary location
     */
    readonly primaryConnectionString: string;
    /**
     * The endpoint URL for file storage in the primary location.
     */
    readonly primaryFileEndpoint: string;
    /**
     * The primary location of the Storage Account.
     */
    readonly primaryLocation: string;
    /**
     * The endpoint URL for queue storage in the primary location.
     */
    readonly primaryQueueEndpoint: string;
    /**
     * The endpoint URL for table storage in the primary location.
     */
    readonly primaryTableEndpoint: string;
    /**
     * The secondary access key for the Storage Account.
     */
    readonly secondaryAccessKey: string;
    /**
     * The connection string associated with the secondary blob location
     */
    readonly secondaryBlobConnectionString: string;
    /**
     * The endpoint URL for blob storage in the secondary location.
     */
    readonly secondaryBlobEndpoint: string;
    /**
     * The connection string associated with the secondary location
     */
    readonly secondaryConnectionString: string;
    /**
     * The secondary location of the Storage Account.
     */
    readonly secondaryLocation: string;
    /**
     * The endpoint URL for queue storage in the secondary location.
     */
    readonly secondaryQueueEndpoint: string;
    /**
     * The endpoint URL for table storage in the secondary location.
     */
    readonly secondaryTableEndpoint: string;
    /**
     * A mapping of tags to assigned to the resource.
     */
    readonly tags: {[key: string]: any};
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}

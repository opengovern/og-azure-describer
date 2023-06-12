package describer

import (
	"context"
	"fmt"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	azblobOld "github.com/Azure/azure-storage-blob-go/azblob"
	"github.com/kaytu-io/kaytu-util/pkg/concurrency"

	"github.com/Azure/azure-sdk-for-go/services/preview/monitor/mgmt/2022-10-01-preview/insights"
	"github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2019-06-01/storage"
	"github.com/Azure/go-autorest/autorest"
	"github.com/kaytu-io/kaytu-azure-describer/azure/model"
	"github.com/tombuildsstuff/giovanni/storage/2018-11-09/queue/queues"
	"github.com/tombuildsstuff/giovanni/storage/2019-12-12/blob/accounts"
)

func StorageContainer(ctx context.Context, authorizer autorest.Authorizer, subscription string, stream *StreamSender) ([]Resource, error) {
	client := storage.NewBlobContainersClient(subscription)
	client.Authorizer = authorizer

	storageClient := storage.NewAccountsClient(subscription)
	storageClient.Authorizer = authorizer

	resultAccounts, err := storageClient.List(ctx)
	if err != nil {
		return nil, err
	}

	wpe := concurrency.NewWorkPool(4)
	var values []Resource
	for {
		for _, ac := range resultAccounts.Values() {
			account := ac
			wpe.AddJob(func() (interface{}, error) {
				resourceGroup := &strings.Split(string(*account.ID), "/")[4]
				result, err := client.List(ctx, *resourceGroup, *account.Name, "", "", "")
				if err != nil {
					return nil, err
				}

				wp := concurrency.NewWorkPool(8)
				for {
					for _, vl := range result.Values() {
						v := vl
						acc := account
						wp.AddJob(func() (interface{}, error) {
							resourceGroup := strings.Split(*v.ID, "/")[4]
							accountName := strings.Split(*v.ID, "/")[8]

							op, err := client.GetImmutabilityPolicy(ctx, resourceGroup, accountName, *v.Name, "")
							if err != nil {
								return nil, err
							}

							return Resource{
								ID:       *v.ID,
								Name:     *v.Name,
								Location: "global",
								Description: JSONAllFieldsMarshaller{
									model.StorageContainerDescription{
										AccountName:        *acc.Name,
										ListContainerItem:  v,
										ImmutabilityPolicy: op,
										ResourceGroup:      resourceGroup,
									},
								},
							}, nil
						})
					}

					if !result.NotDone() {
						break
					}

					err = result.NextWithContext(ctx)
					if err != nil {
						return nil, err
					}
				}
				results := wp.Run()
				var vvv []Resource
				for _, r := range results {
					if r.Error != nil {
						return nil, err
					}
					if r.Value == nil {
						continue
					}
					vvv = append(vvv, r.Value.(Resource))
				}
				return vvv, nil
			})
		}

		if !resultAccounts.NotDone() {
			break
		}

		err = resultAccounts.NextWithContext(ctx)
		if err != nil {
			return nil, err
		}
	}

	results := wpe.Run()
	for _, r := range results {
		if r.Error != nil {
			return nil, err
		}
		if r.Value == nil {
			continue
		}
		values = append(values, r.Value.([]Resource)...)
	}
	if stream != nil {
		for _, resource := range values {
			if err := (*stream)(resource); err != nil {
				return values, err
			}
		}
		values = nil
	}
	return values, nil
}

func StorageAccount(ctx context.Context, authorizer autorest.Authorizer, subscription string, stream *StreamSender) ([]Resource, error) {
	encryptionScopesStorageClient := storage.NewEncryptionScopesClient(subscription)
	encryptionScopesStorageClient.Authorizer = authorizer

	client := insights.NewDiagnosticSettingsClient(subscription)
	client.Authorizer = authorizer

	fileServicesStorageClient := storage.NewFileServicesClient(subscription)
	fileServicesStorageClient.Authorizer = authorizer

	blobServicesStorageClient := storage.NewBlobServicesClient(subscription)
	blobServicesStorageClient.Authorizer = authorizer

	managementPoliciesStorageClient := storage.NewManagementPoliciesClient(subscription)
	managementPoliciesStorageClient.Authorizer = authorizer

	storageClient := storage.NewAccountsClient(subscription)
	storageClient.Authorizer = authorizer

	result, err := storageClient.List(ctx)
	if err != nil {
		return nil, err
	}

	var values []Resource
	for {
		for _, account := range result.Values() {
			resourceGroup := &strings.Split(*account.ID, "/")[4]

			var managementPolicy *storage.ManagementPolicy
			storageGetOp, err := managementPoliciesStorageClient.Get(ctx, *resourceGroup, *account.Name)
			if err != nil {
				if !strings.Contains(err.Error(), "ManagementPolicyNotFound") {
					return nil, err
				}
			} else {
				managementPolicy = &storageGetOp
			}

			var blobServicesProperties *storage.BlobServiceProperties
			if account.Kind != "FileStorage" {
				blobServicesPropertiesOp, err := blobServicesStorageClient.GetServiceProperties(ctx, *resourceGroup, *account.Name)
				if err != nil {
					if strings.Contains(err.Error(), "ContainerOperationFailure") {
						return nil, err
					}
				} else {
					blobServicesProperties = &blobServicesPropertiesOp
				}
			}

			var logging *accounts.Logging
			if account.Kind != "FileStorage" {
				v, err := storageClient.ListKeys(ctx, *resourceGroup, *account.Name, "")
				if err != nil {
					if !strings.Contains(err.Error(), "ScopeLocked") {
						return nil, err
					}
				} else {
					if *v.Keys != nil || len(*v.Keys) > 0 {
						key := (*v.Keys)[0]

						storageAuth, err := autorest.NewSharedKeyAuthorizer(*account.Name, *key.Value, autorest.SharedKeyLite)
						if err != nil {
							return nil, err
						}

						client := accounts.New()
						client.Client.Authorizer = storageAuth
						client.BaseURI = storage.DefaultBaseURI

						resp, err := client.GetServiceProperties(ctx, *account.Name)
						if err != nil {
							if !strings.Contains(err.Error(), "FeatureNotSupportedForAccount") {
								return nil, err
							}
						} else {
							logging = resp.StorageServiceProperties.Logging
						}
					}
				}
			}

			var storageGetServicePropertiesOp *storage.FileServiceProperties
			if account.Kind != "BlobStorage" {
				v, err := fileServicesStorageClient.GetServiceProperties(ctx, *resourceGroup, *account.Name)
				if err != nil {
					if !strings.Contains(err.Error(), "FeatureNotSupportedForAccount") {
						return nil, err
					}
				}
				storageGetServicePropertiesOp = &v
			}

			diagSettingsOp, err := client.List(ctx, *account.ID)
			if err != nil {
				return nil, err
			}

			storageListEncryptionScope, err := encryptionScopesStorageClient.List(ctx, *resourceGroup, *account.Name)
			if err != nil {
				return nil, err
			}
			vsop := storageListEncryptionScope.Values()
			for storageListEncryptionScope.NotDone() {
				err := storageListEncryptionScope.NextWithContext(ctx)
				if err != nil {
					return nil, err
				}

				vsop = append(vsop, storageListEncryptionScope.Values()...)
			}

			var storageProperties *queues.StorageServiceProperties
			if account.Sku.Tier == "Standard" && (account.Kind == "Storage" || account.Kind == "StorageV2") {
				accountKeys, err := storageClient.ListKeys(ctx, *resourceGroup, *account.Name, "")
				if err != nil {
					if !strings.Contains(err.Error(), "ScopeLocked") {
						return nil, err
					}
				} else {
					if *accountKeys.Keys != nil || len(*accountKeys.Keys) > 0 {
						key := (*accountKeys.Keys)[0]
						storageAuth, err := autorest.NewSharedKeyAuthorizer(*account.Name, *key.Value, autorest.SharedKeyLite)
						if err != nil {
							return nil, err
						}

						queuesClient := queues.New()
						queuesClient.Client.Authorizer = storageAuth
						queuesClient.BaseURI = storage.DefaultBaseURI

						resp, err := queuesClient.GetServiceProperties(ctx, *account.Name)

						if err != nil {
							if !strings.Contains(err.Error(), "FeatureNotSupportedForAccount") {
								return nil, err
							}
						} else {
							storageProperties = &resp.StorageServiceProperties
						}
					}
				}
			}

			resource := Resource{
				ID:       *account.ID,
				Name:     *account.Name,
				Location: *account.Location,
				Description: JSONAllFieldsMarshaller{
					model.StorageAccountDescription{
						Account:                     account,
						ManagementPolicy:            managementPolicy,
						BlobServiceProperties:       blobServicesProperties,
						Logging:                     logging,
						StorageServiceProperties:    storageProperties,
						FileServiceProperties:       storageGetServicePropertiesOp,
						DiagnosticSettingsResources: diagSettingsOp.Value,
						EncryptionScopes:            vsop,
						ResourceGroup:               *resourceGroup,
					},
				},
			}
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}
		}
		if !result.NotDone() {
			break
		}
		err = result.NextWithContext(ctx)
		if err != nil {
			return nil, err
		}
	}
	return values, nil
}

func StorageBlob(ctx context.Context, authorizer autorest.Authorizer, subscription string, stream *StreamSender) ([]Resource, error) {
	accountClient := storage.NewAccountsClient(subscription)
	accountClient.Authorizer = authorizer

	containerClient := storage.NewBlobContainersClient(subscription)
	containerClient.Authorizer = authorizer

	storageAccounts, err := accountClient.List(ctx)
	if err != nil {
		return nil, err
	}

	resourceGroups, err := listResourceGroups(ctx, authorizer, subscription)
	if err != nil {
		return nil, err
	}

	var values []Resource
	for {
		for _, storageAccount := range storageAccounts.Values() {
			for _, resourceGroup := range resourceGroups {
				keys, err := accountClient.ListKeys(ctx, *resourceGroup.Name, *storageAccount.Name, "")
				if err != nil {
					return nil, err
				}

				credential, err := azblob.NewSharedKeyCredential(*storageAccount.Name, *(*(keys.Keys))[0].Value)
				if err != nil {
					return nil, err
				}
				baseUrl := fmt.Sprintf("https://%s.blob.core.windows.net", *storageAccount.Name)
				blobClient, err := azblob.NewClientWithSharedKeyCredential(baseUrl, credential, nil)
				if err != nil {
					return nil, err
				}

				containers, err := containerClient.List(ctx, *resourceGroup.Name, *storageAccount.Name, "", "", "")
				if err != nil {
					return nil, err
				}
				for {
					for _, container := range containers.Values() {
						blobPager := blobClient.NewListBlobsFlatPager(*container.Name, nil)
						for blobPager.More() {
							flatResponse, err := blobPager.NextPage(ctx)
							if err != nil {
								return nil, err
							}
							for _, blob := range flatResponse.Segment.BlobItems {
								metadata := azblobOld.Metadata{}
								for k, v := range blob.Metadata {
									metadata[k] = *v
								}

								blobTags := &azblobOld.BlobTags{
									BlobTagSet: []azblobOld.BlobTag{},
								}
								if blob.BlobTags != nil {
									for _, tag := range blob.BlobTags.BlobTagSet {
										blobTags.BlobTagSet = append(blobTags.BlobTagSet, azblobOld.BlobTag{
											Key:   *tag.Key,
											Value: *tag.Value,
										})
									}
								} else {
									blobTags = nil
								}

								resource := Resource{
									ID:       *blob.Name,
									Name:     *blob.Name,
									Location: *storageAccount.Location,
									Description: JSONAllFieldsMarshaller{
										model.StorageBlobDescription{
											Blob: azblobOld.BlobItemInternal{
												Name:             *blob.Name,
												Deleted:          *blob.Deleted,
												Snapshot:         *blob.Snapshot,
												VersionID:        blob.VersionID,
												IsCurrentVersion: blob.IsCurrentVersion,
												Properties: azblobOld.BlobProperties{
													CreationTime:              blob.Properties.CreationTime,
													LastModified:              *blob.Properties.LastModified,
													Etag:                      azblobOld.ETag(*blob.Properties.ETag),
													ContentLength:             blob.Properties.ContentLength,
													ContentType:               blob.Properties.ContentType,
													ContentEncoding:           blob.Properties.ContentEncoding,
													ContentLanguage:           blob.Properties.ContentLanguage,
													ContentMD5:                blob.Properties.ContentMD5,
													ContentDisposition:        blob.Properties.ContentDisposition,
													CacheControl:              blob.Properties.CacheControl,
													BlobSequenceNumber:        blob.Properties.BlobSequenceNumber,
													BlobType:                  azblobOld.BlobType(*blob.Properties.BlobType),
													LeaseStatus:               azblobOld.LeaseStatusType(*blob.Properties.LeaseStatus),
													LeaseState:                azblobOld.LeaseStateType(*blob.Properties.LeaseState),
													LeaseDuration:             azblobOld.LeaseDurationType(*blob.Properties.LeaseDuration),
													CopyID:                    blob.Properties.CopyID,
													CopyStatus:                azblobOld.CopyStatusType(*blob.Properties.CopyStatus),
													CopySource:                blob.Properties.CopySource,
													CopyProgress:              blob.Properties.CopyProgress,
													CopyCompletionTime:        blob.Properties.CopyCompletionTime,
													CopyStatusDescription:     blob.Properties.CopyStatusDescription,
													ServerEncrypted:           blob.Properties.ServerEncrypted,
													IncrementalCopy:           blob.Properties.IncrementalCopy,
													DestinationSnapshot:       blob.Properties.DestinationSnapshot,
													DeletedTime:               blob.Properties.DeletedTime,
													RemainingRetentionDays:    blob.Properties.RemainingRetentionDays,
													AccessTier:                azblobOld.AccessTierType(*blob.Properties.AccessTier),
													AccessTierInferred:        blob.Properties.AccessTierInferred,
													ArchiveStatus:             azblobOld.ArchiveStatusType(*blob.Properties.ArchiveStatus),
													CustomerProvidedKeySha256: blob.Properties.CustomerProvidedKeySHA256,
													EncryptionScope:           blob.Properties.EncryptionScope,
													AccessTierChangeTime:      blob.Properties.AccessTierChangeTime,
													TagCount:                  blob.Properties.TagCount,
													ExpiresOn:                 blob.Properties.ExpiresOn,
													IsSealed:                  blob.Properties.IsSealed,
													RehydratePriority:         azblobOld.RehydratePriorityType(*blob.Properties.RehydratePriority),
													LastAccessedOn:            blob.Properties.LastAccessedOn,
												},
												Metadata: metadata,
												BlobTags: blobTags,
											},
											AccountName:   *storageAccount.Name,
											ContainerName: *container.Name,
											ResourceGroup: *resourceGroup.Name,
											IsSnapshot:    len(*blob.Snapshot) > 0,
										},
									},
								}
								if stream != nil {
									if err := (*stream)(resource); err != nil {
										return nil, err
									}
								} else {
									values = append(values, resource)
								}

							}
						}
					}

					if !containers.NotDone() {
						break
					}
					err := containers.NextWithContext(ctx)
					if err != nil {
						return nil, err
					}
				}
			}
		}
		if !storageAccounts.NotDone() {
			break
		}
		err := storageAccounts.NextWithContext(ctx)
		if err != nil {
			return nil, err
		}
	}

	return values, nil
}

func StorageBlobService(ctx context.Context, authorizer autorest.Authorizer, subscription string, stream *StreamSender) ([]Resource, error) {
	accountClient := storage.NewAccountsClient(subscription)
	accountClient.Authorizer = authorizer

	storageClient := storage.NewBlobServicesClient(subscription)
	storageClient.Authorizer = authorizer

	resourceGroups, err := listResourceGroups(ctx, authorizer, subscription)
	if err != nil {
		return nil, err
	}

	storageAccounts, err := accountClient.List(ctx)
	if err != nil {
		return nil, err
	}

	var values []Resource
	for {
		for _, account := range storageAccounts.Values() {
			for _, resourceGroup := range resourceGroups {
				blobServices, err := storageClient.List(ctx, *resourceGroup.Name, *account.Name)
				if err != nil {
					if strings.Contains(err.Error(), "ParentResourceNotFound") ||
						strings.Contains(err.Error(), "ContainerOperationFailure") ||
						strings.Contains(err.Error(), "FeatureNotSupportedForAccount") {
						continue
					}
					return nil, err
				}
				for _, blobService := range *blobServices.Value {
					resource := Resource{
						ID:       *blobService.ID,
						Name:     *blobService.Name,
						Location: *account.Location,
						Description: JSONAllFieldsMarshaller{
							model.StorageBlobServiceDescription{
								BlobService:   blobService,
								AccountName:   *account.Name,
								Location:      *account.Location,
								ResourceGroup: *resourceGroup.Name,
							},
						},
					}
					if stream != nil {
						if err := (*stream)(resource); err != nil {
							return nil, err
						}
					} else {
						values = append(values, resource)
					}
				}
			}
		}

		if !storageAccounts.NotDone() {
			break
		}
		err := storageAccounts.NextWithContext(ctx)
		if err != nil {
			return nil, err
		}
	}

	return values, nil
}

func StorageQueue(ctx context.Context, authorizer autorest.Authorizer, subscription string, stream *StreamSender) ([]Resource, error) {
	accountClient := storage.NewAccountsClient(subscription)
	accountClient.Authorizer = authorizer

	storageClient := storage.NewQueueClient(subscription)
	storageClient.Authorizer = authorizer

	resourceGroups, err := listResourceGroups(ctx, authorizer, subscription)
	if err != nil {
		return nil, err
	}

	storageAccounts, err := accountClient.List(ctx)
	if err != nil {
		return nil, err
	}

	var values []Resource
	for {
		for _, account := range storageAccounts.Values() {
			if account.Kind == "FileStorage" || account.Kind == "BlockBlobStorage" {
				continue
			}

			for _, resourceGroup := range resourceGroups {
				queuesRes, err := storageClient.List(ctx, *resourceGroup.Name, *account.Name, "", "")
				if err != nil {
					/*
					* For storage account type 'Page Blob' we are getting the kind value as 'StorageV2'.
					* Storage account type 'Page Blob' does not support table, so we are getting 'FeatureNotSupportedForAccount'/'OperationNotAllowedOnKind' error.
					* With same kind(StorageV2) of storage account, we my have different type(File Share) of storage account so we need to handle this particular error.
					 */
					if strings.Contains(err.Error(), "FeatureNotSupportedForAccount") ||
						strings.Contains(err.Error(), "AccountIsDisabled") ||
						strings.Contains(err.Error(), "ParentResourceNotFound") ||
						strings.Contains(err.Error(), "OperationNotAllowedOnKind") {
						continue
					} else {
						return nil, err
					}
				}
				for {
					for _, queue := range queuesRes.Values() {
						resource := Resource{
							ID:       *queue.ID,
							Name:     *queue.Name,
							Location: *account.Location,
							Description: JSONAllFieldsMarshaller{
								model.StorageQueueDescription{
									Queue:         queue,
									AccountName:   *account.Name,
									Location:      *account.Location,
									ResourceGroup: *resourceGroup.Name,
								},
							},
						}
						if stream != nil {
							if err := (*stream)(resource); err != nil {
								return nil, err
							}
						} else {
							values = append(values, resource)
						}
					}
					if !queuesRes.NotDone() {
						break
					}
					err := queuesRes.NextWithContext(ctx)
					if err != nil {
						return nil, err
					}
				}
			}
		}

		if !storageAccounts.NotDone() {
			break
		}
		err := storageAccounts.NextWithContext(ctx)
		if err != nil {
			return nil, err
		}
	}

	return values, nil
}

func StorageFileShare(ctx context.Context, authorizer autorest.Authorizer, subscription string, stream *StreamSender) ([]Resource, error) {
	accountClient := storage.NewAccountsClient(subscription)
	accountClient.Authorizer = authorizer

	storageClient := storage.NewFileSharesClient(subscription)
	storageClient.Authorizer = authorizer

	resourceGroups, err := listResourceGroups(ctx, authorizer, subscription)
	if err != nil {
		return nil, err
	}

	storageAccounts, err := accountClient.List(ctx)
	if err != nil {
		return nil, err
	}

	var values []Resource
	for {
		for _, account := range storageAccounts.Values() {
			for _, resourceGroup := range resourceGroups {
				fileShares, err := storageClient.List(ctx, *resourceGroup.Name, *account.Name, "", "", "")
				if err != nil {
					if strings.Contains(err.Error(), "ParentResourceNotFound") ||
						strings.Contains(err.Error(), "FeatureNotSupportedForAccount") ||
						strings.Contains(err.Error(), "AccountIsDisabled") {
						continue
					}
					return nil, err
				}
				for {
					for _, fileShareItem := range fileShares.Values() {
						resource := Resource{
							ID:       *fileShareItem.ID,
							Name:     *fileShareItem.Name,
							Location: *account.Location,
							Description: JSONAllFieldsMarshaller{
								model.StorageFileShareDescription{
									FileShare:     fileShareItem,
									AccountName:   *account.Name,
									Location:      *account.Location,
									ResourceGroup: *resourceGroup.Name,
								},
							},
						}
						if stream != nil {
							if err := (*stream)(resource); err != nil {
								return nil, err
							}
						} else {
							values = append(values, resource)
						}
					}
					if !fileShares.NotDone() {
						break
					}
					err := fileShares.NextWithContext(ctx)
					if err != nil {
						return nil, err
					}
				}
			}
		}

		if !storageAccounts.NotDone() {
			break
		}
		err := storageAccounts.NextWithContext(ctx)
		if err != nil {
			return nil, err
		}
	}

	return values, nil
}

func StorageTable(ctx context.Context, authorizer autorest.Authorizer, subscription string, stream *StreamSender) ([]Resource, error) {
	accountClient := storage.NewAccountsClient(subscription)
	accountClient.Authorizer = authorizer

	storageClient := storage.NewTableClient(subscription)
	storageClient.Authorizer = authorizer

	resourceGroups, err := listResourceGroups(ctx, authorizer, subscription)
	if err != nil {
		return nil, err
	}

	storageAccounts, err := accountClient.List(ctx)
	if err != nil {
		return nil, err
	}

	var values []Resource
	for {
		for _, account := range storageAccounts.Values() {
			if account.Kind == "FileStorage" || account.Kind == "BlockBlobStorage" {
				continue
			}
			for _, resourceGroup := range resourceGroups {
				tables, err := storageClient.List(ctx, *resourceGroup.Name, *account.Name)
				if err != nil {
					/*
					* For storage account type 'Page Blob' we are getting the kind value as 'StorageV2'.
					* Storage account type 'Page Blob' does not support table, so we are getting 'FeatureNotSupportedForAccount'/'OperationNotAllowedOnKind' error.
					* With same kind(StorageV2) of storage account, we my have different type(File Share) of storage account so we need to handle this particular error.
					 */
					if strings.Contains(err.Error(), "FeatureNotSupportedForAccount") ||
						strings.Contains(err.Error(), "OperationNotAllowedOnKind") ||
						strings.Contains(err.Error(), "ParentResourceNotFound") {
						continue
					}
					return nil, err
				}
				for {
					for _, table := range tables.Values() {
						resource := Resource{
							ID:       *table.ID,
							Name:     *table.Name,
							Location: *account.Location,
							Description: JSONAllFieldsMarshaller{
								model.StorageTableDescription{
									Table:         table,
									AccountName:   *account.Name,
									Location:      *account.Location,
									ResourceGroup: *resourceGroup.Name,
								},
							},
						}
						if stream != nil {
							if err := (*stream)(resource); err != nil {
								return nil, err
							}
						} else {
							values = append(values, resource)
						}
					}
					if !tables.NotDone() {
						break
					}
					err := tables.NextWithContext(ctx)
					if err != nil {
						return nil, err
					}
				}
			}
		}

		if !storageAccounts.NotDone() {
			break
		}
		err := storageAccounts.NextWithContext(ctx)
		if err != nil {
			return nil, err
		}
	}

	return values, nil
}

func StorageTableService(ctx context.Context, authorizer autorest.Authorizer, subscription string, stream *StreamSender) ([]Resource, error) {
	accountClient := storage.NewAccountsClient(subscription)
	accountClient.Authorizer = authorizer

	storageClient := storage.NewTableServicesClient(subscription)
	storageClient.Authorizer = authorizer

	resourceGroups, err := listResourceGroups(ctx, authorizer, subscription)
	if err != nil {
		return nil, err
	}

	storageAccounts, err := accountClient.List(ctx)
	if err != nil {
		return nil, err
	}

	var values []Resource
	for {
		for _, account := range storageAccounts.Values() {
			if account.Kind == "FileStorage" {
				continue
			}
			for _, resourceGroup := range resourceGroups {
				tableServices, err := storageClient.List(ctx, *resourceGroup.Name, *account.Name)
				if err != nil {
					if strings.Contains(err.Error(), "ParentResourceNotFound") ||
						strings.Contains(err.Error(), "FeatureNotSupportedForAccount") {
						continue
					}
					return nil, err
				}

				for _, tableService := range *tableServices.Value {
					resource := Resource{
						ID:       *tableService.ID,
						Name:     *tableService.Name,
						Location: *account.Location,
						Description: JSONAllFieldsMarshaller{
							model.StorageTableServiceDescription{
								TableService:  tableService,
								AccountName:   *account.Name,
								Location:      *account.Location,
								ResourceGroup: *resourceGroup.Name,
							},
						},
					}
					if stream != nil {
						if err := (*stream)(resource); err != nil {
							return nil, err
						}
					} else {
						values = append(values, resource)
					}
				}
			}
		}

		if !storageAccounts.NotDone() {
			break
		}
		err := storageAccounts.NextWithContext(ctx)
		if err != nil {
			return nil, err
		}
	}

	return values, nil
}

//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armstoragecache_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagecache/armstoragecache/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d4205894880b989ede35d62d97c8e901ed14fb5a/specification/storagecache/resource-manager/Microsoft.StorageCache/preview/2023-11-01-preview/examples/StorageTargets_DnsRefresh.json
func ExampleStorageTargetsClient_BeginDNSRefresh() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstoragecache.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewStorageTargetsClient().BeginDNSRefresh(ctx, "scgroup", "sc", "st1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d4205894880b989ede35d62d97c8e901ed14fb5a/specification/storagecache/resource-manager/Microsoft.StorageCache/preview/2023-11-01-preview/examples/StorageTargets_ListByCache.json
func ExampleStorageTargetsClient_NewListByCachePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstoragecache.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewStorageTargetsClient().NewListByCachePager("scgroup", "sc1", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.StorageTargetsResult = armstoragecache.StorageTargetsResult{
		// 	Value: []*armstoragecache.StorageTarget{
		// 		{
		// 			Name: to.Ptr("st1"),
		// 			Type: to.Ptr("Microsoft.StorageCache/Cache/StorageTarget"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/scgroup/providers/Microsoft.StorageCache/caches/sc1/storagetargets/st1"),
		// 			SystemData: &armstoragecache.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.123Z"); return t}()),
		// 				CreatedBy: to.Ptr("user1"),
		// 				CreatedByType: to.Ptr(armstoragecache.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.123Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("user2"),
		// 				LastModifiedByType: to.Ptr(armstoragecache.CreatedByTypeUser),
		// 			},
		// 			Properties: &armstoragecache.StorageTargetProperties{
		// 				AllocationPercentage: to.Ptr[int32](25),
		// 				Junctions: []*armstoragecache.NamespaceJunction{
		// 					{
		// 						NamespacePath: to.Ptr("/path/on/cache"),
		// 						NfsAccessPolicy: to.Ptr("default"),
		// 						NfsExport: to.Ptr("exp1"),
		// 						TargetPath: to.Ptr("/path/on/exp1"),
		// 					},
		// 					{
		// 						NamespacePath: to.Ptr("/path2/on/cache"),
		// 						NfsAccessPolicy: to.Ptr("default"),
		// 						NfsExport: to.Ptr("exp2"),
		// 						TargetPath: to.Ptr("/path2/on/exp2"),
		// 				}},
		// 				Nfs3: &armstoragecache.Nfs3Target{
		// 					Target: to.Ptr("10.0.44.44"),
		// 					UsageModel: to.Ptr("READ_ONLY"),
		// 				},
		// 				State: to.Ptr(armstoragecache.OperationalStateTypeReady),
		// 				TargetType: to.Ptr(armstoragecache.StorageTargetTypeNfs3),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("st2"),
		// 			Type: to.Ptr("Microsoft.StorageCache/Cache/StorageTarget"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/scgroup/providers/Microsoft.StorageCache/caches/sc1/storagetargets/st2"),
		// 			SystemData: &armstoragecache.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.123Z"); return t}()),
		// 				CreatedBy: to.Ptr("user1"),
		// 				CreatedByType: to.Ptr(armstoragecache.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.123Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("user2"),
		// 				LastModifiedByType: to.Ptr(armstoragecache.CreatedByTypeUser),
		// 			},
		// 			Properties: &armstoragecache.StorageTargetProperties{
		// 				AllocationPercentage: to.Ptr[int32](50),
		// 				Clfs: &armstoragecache.ClfsTarget{
		// 					Target: to.Ptr("https://contoso123.blob.core.windows.net/contoso123"),
		// 				},
		// 				Junctions: []*armstoragecache.NamespaceJunction{
		// 					{
		// 						NamespacePath: to.Ptr("/some/arbitrary/place/on/cache"),
		// 						NfsAccessPolicy: to.Ptr("default"),
		// 						TargetPath: to.Ptr("/"),
		// 				}},
		// 				State: to.Ptr(armstoragecache.OperationalStateTypeReady),
		// 				TargetType: to.Ptr(armstoragecache.StorageTargetTypeClfs),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("st3"),
		// 			Type: to.Ptr("Microsoft.StorageCache/Cache/StorageTarget"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/scgroup/providers/Microsoft.StorageCache/caches/sc1/storagetargets/st3"),
		// 			SystemData: &armstoragecache.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.123Z"); return t}()),
		// 				CreatedBy: to.Ptr("user1"),
		// 				CreatedByType: to.Ptr(armstoragecache.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.123Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("user2"),
		// 				LastModifiedByType: to.Ptr(armstoragecache.CreatedByTypeUser),
		// 			},
		// 			Properties: &armstoragecache.StorageTargetProperties{
		// 				AllocationPercentage: to.Ptr[int32](25),
		// 				Junctions: []*armstoragecache.NamespaceJunction{
		// 					{
		// 						NamespacePath: to.Ptr("/some/crazier/place/on/cache"),
		// 						NfsAccessPolicy: to.Ptr("default"),
		// 						NfsExport: to.Ptr(""),
		// 						TargetPath: to.Ptr("/"),
		// 				}},
		// 				State: to.Ptr(armstoragecache.OperationalStateTypeReady),
		// 				TargetType: to.Ptr(armstoragecache.StorageTargetTypeUnknown),
		// 				Unknown: &armstoragecache.UnknownTarget{
		// 					Attributes: map[string]*string{
		// 						"foo": to.Ptr("bar"),
		// 						"foo2": to.Ptr("test"),
		// 					},
		// 				},
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d4205894880b989ede35d62d97c8e901ed14fb5a/specification/storagecache/resource-manager/Microsoft.StorageCache/preview/2023-11-01-preview/examples/StorageTargets_Delete.json
func ExampleStorageTargetsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstoragecache.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewStorageTargetsClient().BeginDelete(ctx, "scgroup", "sc1", "st1", &armstoragecache.StorageTargetsClientBeginDeleteOptions{Force: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d4205894880b989ede35d62d97c8e901ed14fb5a/specification/storagecache/resource-manager/Microsoft.StorageCache/preview/2023-11-01-preview/examples/StorageTargets_Get.json
func ExampleStorageTargetsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstoragecache.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewStorageTargetsClient().Get(ctx, "scgroup", "sc1", "st1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.StorageTarget = armstoragecache.StorageTarget{
	// 	Name: to.Ptr("st1"),
	// 	Type: to.Ptr("Microsoft.StorageCache/Cache/StorageTarget"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/scgroup/providers/Microsoft.StorageCache/caches/sc1/storagetargets/st1"),
	// 	SystemData: &armstoragecache.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.123Z"); return t}()),
	// 		CreatedBy: to.Ptr("user1"),
	// 		CreatedByType: to.Ptr(armstoragecache.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.123Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user2"),
	// 		LastModifiedByType: to.Ptr(armstoragecache.CreatedByTypeUser),
	// 	},
	// 	Properties: &armstoragecache.StorageTargetProperties{
	// 		AllocationPercentage: to.Ptr[int32](25),
	// 		Junctions: []*armstoragecache.NamespaceJunction{
	// 			{
	// 				NamespacePath: to.Ptr("/path/on/cache"),
	// 				NfsAccessPolicy: to.Ptr("default"),
	// 				NfsExport: to.Ptr("exp1"),
	// 				TargetPath: to.Ptr("/path/on/exp1"),
	// 			},
	// 			{
	// 				NamespacePath: to.Ptr("/path2/on/cache"),
	// 				NfsAccessPolicy: to.Ptr("default"),
	// 				NfsExport: to.Ptr("exp2"),
	// 				TargetPath: to.Ptr("/path2/on/exp2"),
	// 		}},
	// 		Nfs3: &armstoragecache.Nfs3Target{
	// 			Target: to.Ptr("10.0.44.44"),
	// 			UsageModel: to.Ptr("READ_HEAVY_FREQ"),
	// 		},
	// 		State: to.Ptr(armstoragecache.OperationalStateTypeReady),
	// 		TargetType: to.Ptr(armstoragecache.StorageTargetTypeNfs3),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d4205894880b989ede35d62d97c8e901ed14fb5a/specification/storagecache/resource-manager/Microsoft.StorageCache/preview/2023-11-01-preview/examples/StorageTargets_CreateOrUpdate.json
func ExampleStorageTargetsClient_BeginCreateOrUpdate_storageTargetsCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstoragecache.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewStorageTargetsClient().BeginCreateOrUpdate(ctx, "scgroup", "sc1", "st1", armstoragecache.StorageTarget{
		Properties: &armstoragecache.StorageTargetProperties{
			Junctions: []*armstoragecache.NamespaceJunction{
				{
					NamespacePath:   to.Ptr("/path/on/cache"),
					NfsAccessPolicy: to.Ptr("default"),
					NfsExport:       to.Ptr("exp1"),
					TargetPath:      to.Ptr("/path/on/exp1"),
				},
				{
					NamespacePath:   to.Ptr("/path2/on/cache"),
					NfsAccessPolicy: to.Ptr("rootSquash"),
					NfsExport:       to.Ptr("exp2"),
					TargetPath:      to.Ptr("/path2/on/exp2"),
				}},
			Nfs3: &armstoragecache.Nfs3Target{
				Target:            to.Ptr("10.0.44.44"),
				UsageModel:        to.Ptr("READ_ONLY"),
				VerificationTimer: to.Ptr[int32](30),
			},
			TargetType: to.Ptr(armstoragecache.StorageTargetTypeNfs3),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.StorageTarget = armstoragecache.StorageTarget{
	// 	Name: to.Ptr("st1"),
	// 	Type: to.Ptr("Microsoft.StorageCache/Cache/StorageTarget"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/scgroup/providers/Microsoft.StorageCache/caches/sc1/storagetargets/st1"),
	// 	SystemData: &armstoragecache.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.123Z"); return t}()),
	// 		CreatedBy: to.Ptr("user1"),
	// 		CreatedByType: to.Ptr(armstoragecache.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.123Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user2"),
	// 		LastModifiedByType: to.Ptr(armstoragecache.CreatedByTypeUser),
	// 	},
	// 	Properties: &armstoragecache.StorageTargetProperties{
	// 		Junctions: []*armstoragecache.NamespaceJunction{
	// 			{
	// 				NamespacePath: to.Ptr("/path/on/cache"),
	// 				NfsAccessPolicy: to.Ptr("default"),
	// 				NfsExport: to.Ptr("exp1"),
	// 				TargetPath: to.Ptr("/path/on/exp1"),
	// 			},
	// 			{
	// 				NamespacePath: to.Ptr("/path2/on/cache"),
	// 				NfsAccessPolicy: to.Ptr("rootSquash"),
	// 				NfsExport: to.Ptr("exp2"),
	// 				TargetPath: to.Ptr("/path2/on/exp2"),
	// 		}},
	// 		Nfs3: &armstoragecache.Nfs3Target{
	// 			Target: to.Ptr("10.0.44.44"),
	// 			UsageModel: to.Ptr("READ_ONLY"),
	// 			VerificationTimer: to.Ptr[int32](30),
	// 		},
	// 		State: to.Ptr(armstoragecache.OperationalStateTypeReady),
	// 		TargetType: to.Ptr(armstoragecache.StorageTargetTypeNfs3),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d4205894880b989ede35d62d97c8e901ed14fb5a/specification/storagecache/resource-manager/Microsoft.StorageCache/preview/2023-11-01-preview/examples/StorageTargets_CreateOrUpdate_BlobNfs.json
func ExampleStorageTargetsClient_BeginCreateOrUpdate_storageTargetsCreateOrUpdateBlobNfs() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstoragecache.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewStorageTargetsClient().BeginCreateOrUpdate(ctx, "scgroup", "sc1", "st1", armstoragecache.StorageTarget{
		Properties: &armstoragecache.StorageTargetProperties{
			BlobNfs: &armstoragecache.BlobNfsTarget{
				Target:            to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/scgroup/providers/Microsoft.Storage/storageAccounts/blofnfs/blobServices/default/containers/blobnfs"),
				UsageModel:        to.Ptr("READ_WRITE"),
				VerificationTimer: to.Ptr[int32](28800),
				WriteBackTimer:    to.Ptr[int32](3600),
			},
			Junctions: []*armstoragecache.NamespaceJunction{
				{
					NamespacePath: to.Ptr("/blobnfs"),
				}},
			TargetType: to.Ptr(armstoragecache.StorageTargetTypeBlobNfs),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.StorageTarget = armstoragecache.StorageTarget{
	// 	Name: to.Ptr("st1"),
	// 	Type: to.Ptr("Microsoft.StorageCache/Cache/StorageTarget"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/scgroup/providers/Microsoft.StorageCache/caches/sc1/storagetargets/st1"),
	// 	SystemData: &armstoragecache.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.123Z"); return t}()),
	// 		CreatedBy: to.Ptr("user1"),
	// 		CreatedByType: to.Ptr(armstoragecache.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.123Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user2"),
	// 		LastModifiedByType: to.Ptr(armstoragecache.CreatedByTypeUser),
	// 	},
	// 	Properties: &armstoragecache.StorageTargetProperties{
	// 		BlobNfs: &armstoragecache.BlobNfsTarget{
	// 			Target: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/scgroup/providers/Microsoft.Storage/storageAccounts/blofnfs/blobServices/default/containers/blobnfs"),
	// 			UsageModel: to.Ptr("READ_WRITE"),
	// 			VerificationTimer: to.Ptr[int32](28800),
	// 			WriteBackTimer: to.Ptr[int32](3600),
	// 		},
	// 		Junctions: []*armstoragecache.NamespaceJunction{
	// 			{
	// 				NamespacePath: to.Ptr("/blobnfs"),
	// 		}},
	// 		State: to.Ptr(armstoragecache.OperationalStateTypeReady),
	// 		TargetType: to.Ptr(armstoragecache.StorageTargetTypeBlobNfs),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d4205894880b989ede35d62d97c8e901ed14fb5a/specification/storagecache/resource-manager/Microsoft.StorageCache/preview/2023-11-01-preview/examples/StorageTargets_CreateOrUpdate_NoJunctions.json
func ExampleStorageTargetsClient_BeginCreateOrUpdate_storageTargetsCreateOrUpdateNoJunctions() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstoragecache.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewStorageTargetsClient().BeginCreateOrUpdate(ctx, "scgroup", "sc1", "st1", armstoragecache.StorageTarget{
		Properties: &armstoragecache.StorageTargetProperties{
			Nfs3: &armstoragecache.Nfs3Target{
				Target:            to.Ptr("10.0.44.44"),
				UsageModel:        to.Ptr("READ_ONLY"),
				VerificationTimer: to.Ptr[int32](30),
			},
			TargetType: to.Ptr(armstoragecache.StorageTargetTypeNfs3),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.StorageTarget = armstoragecache.StorageTarget{
	// 	Name: to.Ptr("st1"),
	// 	Type: to.Ptr("Microsoft.StorageCache/Cache/StorageTarget"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/scgroup/providers/Microsoft.StorageCache/caches/sc1/storagetargets/st1"),
	// 	SystemData: &armstoragecache.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.123Z"); return t}()),
	// 		CreatedBy: to.Ptr("user1"),
	// 		CreatedByType: to.Ptr(armstoragecache.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.123Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user2"),
	// 		LastModifiedByType: to.Ptr(armstoragecache.CreatedByTypeUser),
	// 	},
	// 	Properties: &armstoragecache.StorageTargetProperties{
	// 		Nfs3: &armstoragecache.Nfs3Target{
	// 			Target: to.Ptr("10.0.44.44"),
	// 			UsageModel: to.Ptr("READ_ONLY"),
	// 			VerificationTimer: to.Ptr[int32](30),
	// 		},
	// 		State: to.Ptr(armstoragecache.OperationalStateTypeReady),
	// 		TargetType: to.Ptr(armstoragecache.StorageTargetTypeNfs3),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d4205894880b989ede35d62d97c8e901ed14fb5a/specification/storagecache/resource-manager/Microsoft.StorageCache/preview/2023-11-01-preview/examples/StorageTargets_RestoreDefaults.json
func ExampleStorageTargetsClient_BeginRestoreDefaults() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstoragecache.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewStorageTargetsClient().BeginRestoreDefaults(ctx, "scgroup", "sc", "st1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

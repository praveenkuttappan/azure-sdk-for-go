// +build go1.9

// Copyright 2020 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package storagecache

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/storagecache/mgmt/2020-03-01/storagecache"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type CacheIdentityType = original.CacheIdentityType

const (
	None           CacheIdentityType = original.None
	SystemAssigned CacheIdentityType = original.SystemAssigned
)

type CreatedByType = original.CreatedByType

const (
	Application     CreatedByType = original.Application
	Key             CreatedByType = original.Key
	ManagedIdentity CreatedByType = original.ManagedIdentity
	User            CreatedByType = original.User
)

type FirmwareStatusType = original.FirmwareStatusType

const (
	Available   FirmwareStatusType = original.Available
	Unavailable FirmwareStatusType = original.Unavailable
)

type HealthStateType = original.HealthStateType

const (
	Degraded      HealthStateType = original.Degraded
	Down          HealthStateType = original.Down
	Flushing      HealthStateType = original.Flushing
	Healthy       HealthStateType = original.Healthy
	Stopped       HealthStateType = original.Stopped
	Stopping      HealthStateType = original.Stopping
	Transitioning HealthStateType = original.Transitioning
	Unknown       HealthStateType = original.Unknown
	Upgrading     HealthStateType = original.Upgrading
)

type MetricAggregationType = original.MetricAggregationType

const (
	MetricAggregationTypeAverage      MetricAggregationType = original.MetricAggregationTypeAverage
	MetricAggregationTypeCount        MetricAggregationType = original.MetricAggregationTypeCount
	MetricAggregationTypeMaximum      MetricAggregationType = original.MetricAggregationTypeMaximum
	MetricAggregationTypeMinimum      MetricAggregationType = original.MetricAggregationTypeMinimum
	MetricAggregationTypeNone         MetricAggregationType = original.MetricAggregationTypeNone
	MetricAggregationTypeNotSpecified MetricAggregationType = original.MetricAggregationTypeNotSpecified
	MetricAggregationTypeTotal        MetricAggregationType = original.MetricAggregationTypeTotal
)

type ProvisioningStateType = original.ProvisioningStateType

const (
	Cancelled ProvisioningStateType = original.Cancelled
	Creating  ProvisioningStateType = original.Creating
	Deleting  ProvisioningStateType = original.Deleting
	Failed    ProvisioningStateType = original.Failed
	Succeeded ProvisioningStateType = original.Succeeded
	Updating  ProvisioningStateType = original.Updating
)

type ReasonCode = original.ReasonCode

const (
	NotAvailableForSubscription ReasonCode = original.NotAvailableForSubscription
	QuotaID                     ReasonCode = original.QuotaID
)

type TargetType = original.TargetType

const (
	TargetTypeClfs                    TargetType = original.TargetTypeClfs
	TargetTypeNfs3                    TargetType = original.TargetTypeNfs3
	TargetTypeStorageTargetProperties TargetType = original.TargetTypeStorageTargetProperties
	TargetTypeUnknown                 TargetType = original.TargetTypeUnknown
)

type APIOperation = original.APIOperation
type APIOperationDisplay = original.APIOperationDisplay
type APIOperationListResult = original.APIOperationListResult
type APIOperationListResultIterator = original.APIOperationListResultIterator
type APIOperationListResultPage = original.APIOperationListResultPage
type APIOperationProperties = original.APIOperationProperties
type APIOperationPropertiesServiceSpecification = original.APIOperationPropertiesServiceSpecification
type AscOperation = original.AscOperation
type AscOperationsClient = original.AscOperationsClient
type BaseClient = original.BaseClient
type BasicStorageTargetProperties = original.BasicStorageTargetProperties
type Cache = original.Cache
type CacheEncryptionSettings = original.CacheEncryptionSettings
type CacheHealth = original.CacheHealth
type CacheIdentity = original.CacheIdentity
type CacheNetworkSettings = original.CacheNetworkSettings
type CacheProperties = original.CacheProperties
type CacheSecuritySettings = original.CacheSecuritySettings
type CacheSku = original.CacheSku
type CacheUpgradeStatus = original.CacheUpgradeStatus
type CachesClient = original.CachesClient
type CachesCreateOrUpdateFuture = original.CachesCreateOrUpdateFuture
type CachesDeleteFuture = original.CachesDeleteFuture
type CachesFlushFuture = original.CachesFlushFuture
type CachesListResult = original.CachesListResult
type CachesListResultIterator = original.CachesListResultIterator
type CachesListResultPage = original.CachesListResultPage
type CachesStartFuture = original.CachesStartFuture
type CachesStopFuture = original.CachesStopFuture
type CachesUpgradeFirmwareFuture = original.CachesUpgradeFirmwareFuture
type ClfsTarget = original.ClfsTarget
type ClfsTargetProperties = original.ClfsTargetProperties
type CloudError = original.CloudError
type CloudErrorBody = original.CloudErrorBody
type ErrorResponse = original.ErrorResponse
type KeyVaultKeyReference = original.KeyVaultKeyReference
type KeyVaultKeyReferenceSourceVault = original.KeyVaultKeyReferenceSourceVault
type MetricDimension = original.MetricDimension
type MetricSpecification = original.MetricSpecification
type NamespaceJunction = original.NamespaceJunction
type Nfs3Target = original.Nfs3Target
type Nfs3TargetProperties = original.Nfs3TargetProperties
type OperationsClient = original.OperationsClient
type ResourceSku = original.ResourceSku
type ResourceSkuCapabilities = original.ResourceSkuCapabilities
type ResourceSkuLocationInfo = original.ResourceSkuLocationInfo
type ResourceSkusResult = original.ResourceSkusResult
type ResourceSkusResultIterator = original.ResourceSkusResultIterator
type ResourceSkusResultPage = original.ResourceSkusResultPage
type Restriction = original.Restriction
type SetObject = original.SetObject
type SkusClient = original.SkusClient
type StorageTarget = original.StorageTarget
type StorageTargetProperties = original.StorageTargetProperties
type StorageTargetResource = original.StorageTargetResource
type StorageTargetsClient = original.StorageTargetsClient
type StorageTargetsCreateOrUpdateFuture = original.StorageTargetsCreateOrUpdateFuture
type StorageTargetsDeleteFuture = original.StorageTargetsDeleteFuture
type StorageTargetsResult = original.StorageTargetsResult
type StorageTargetsResultIterator = original.StorageTargetsResultIterator
type StorageTargetsResultPage = original.StorageTargetsResultPage
type SystemData = original.SystemData
type UnknownTarget = original.UnknownTarget
type UnknownTargetProperties = original.UnknownTargetProperties
type UsageModel = original.UsageModel
type UsageModelDisplay = original.UsageModelDisplay
type UsageModelsClient = original.UsageModelsClient
type UsageModelsResult = original.UsageModelsResult
type UsageModelsResultIterator = original.UsageModelsResultIterator
type UsageModelsResultPage = original.UsageModelsResultPage

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewAPIOperationListResultIterator(page APIOperationListResultPage) APIOperationListResultIterator {
	return original.NewAPIOperationListResultIterator(page)
}
func NewAPIOperationListResultPage(cur APIOperationListResult, getNextPage func(context.Context, APIOperationListResult) (APIOperationListResult, error)) APIOperationListResultPage {
	return original.NewAPIOperationListResultPage(cur, getNextPage)
}
func NewAscOperationsClient(subscriptionID string) AscOperationsClient {
	return original.NewAscOperationsClient(subscriptionID)
}
func NewAscOperationsClientWithBaseURI(baseURI string, subscriptionID string) AscOperationsClient {
	return original.NewAscOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewCachesClient(subscriptionID string) CachesClient {
	return original.NewCachesClient(subscriptionID)
}
func NewCachesClientWithBaseURI(baseURI string, subscriptionID string) CachesClient {
	return original.NewCachesClientWithBaseURI(baseURI, subscriptionID)
}
func NewCachesListResultIterator(page CachesListResultPage) CachesListResultIterator {
	return original.NewCachesListResultIterator(page)
}
func NewCachesListResultPage(cur CachesListResult, getNextPage func(context.Context, CachesListResult) (CachesListResult, error)) CachesListResultPage {
	return original.NewCachesListResultPage(cur, getNextPage)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewResourceSkusResultIterator(page ResourceSkusResultPage) ResourceSkusResultIterator {
	return original.NewResourceSkusResultIterator(page)
}
func NewResourceSkusResultPage(cur ResourceSkusResult, getNextPage func(context.Context, ResourceSkusResult) (ResourceSkusResult, error)) ResourceSkusResultPage {
	return original.NewResourceSkusResultPage(cur, getNextPage)
}
func NewSkusClient(subscriptionID string) SkusClient {
	return original.NewSkusClient(subscriptionID)
}
func NewSkusClientWithBaseURI(baseURI string, subscriptionID string) SkusClient {
	return original.NewSkusClientWithBaseURI(baseURI, subscriptionID)
}
func NewStorageTargetsClient(subscriptionID string) StorageTargetsClient {
	return original.NewStorageTargetsClient(subscriptionID)
}
func NewStorageTargetsClientWithBaseURI(baseURI string, subscriptionID string) StorageTargetsClient {
	return original.NewStorageTargetsClientWithBaseURI(baseURI, subscriptionID)
}
func NewStorageTargetsResultIterator(page StorageTargetsResultPage) StorageTargetsResultIterator {
	return original.NewStorageTargetsResultIterator(page)
}
func NewStorageTargetsResultPage(cur StorageTargetsResult, getNextPage func(context.Context, StorageTargetsResult) (StorageTargetsResult, error)) StorageTargetsResultPage {
	return original.NewStorageTargetsResultPage(cur, getNextPage)
}
func NewUsageModelsClient(subscriptionID string) UsageModelsClient {
	return original.NewUsageModelsClient(subscriptionID)
}
func NewUsageModelsClientWithBaseURI(baseURI string, subscriptionID string) UsageModelsClient {
	return original.NewUsageModelsClientWithBaseURI(baseURI, subscriptionID)
}
func NewUsageModelsResultIterator(page UsageModelsResultPage) UsageModelsResultIterator {
	return original.NewUsageModelsResultIterator(page)
}
func NewUsageModelsResultPage(cur UsageModelsResult, getNextPage func(context.Context, UsageModelsResult) (UsageModelsResult, error)) UsageModelsResultPage {
	return original.NewUsageModelsResultPage(cur, getNextPage)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleCacheIdentityTypeValues() []CacheIdentityType {
	return original.PossibleCacheIdentityTypeValues()
}
func PossibleCreatedByTypeValues() []CreatedByType {
	return original.PossibleCreatedByTypeValues()
}
func PossibleFirmwareStatusTypeValues() []FirmwareStatusType {
	return original.PossibleFirmwareStatusTypeValues()
}
func PossibleHealthStateTypeValues() []HealthStateType {
	return original.PossibleHealthStateTypeValues()
}
func PossibleMetricAggregationTypeValues() []MetricAggregationType {
	return original.PossibleMetricAggregationTypeValues()
}
func PossibleProvisioningStateTypeValues() []ProvisioningStateType {
	return original.PossibleProvisioningStateTypeValues()
}
func PossibleReasonCodeValues() []ReasonCode {
	return original.PossibleReasonCodeValues()
}
func PossibleTargetTypeValues() []TargetType {
	return original.PossibleTargetTypeValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}

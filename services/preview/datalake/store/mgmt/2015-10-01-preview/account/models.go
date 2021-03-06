package account

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"encoding/json"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/Azure/go-autorest/tracing"
	"github.com/satori/go.uuid"
	"net/http"
)

// The package's fully qualified name.
const fqdn = "github.com/Azure/azure-sdk-for-go/services/preview/datalake/store/mgmt/2015-10-01-preview/account"

// AzureAsyncOperationResult the response body contains the status of the specified asynchronous operation,
// indicating whether it has succeeded, is in progress, or has failed. Note that this status is distinct
// from the HTTP status code returned for the Get Operation Status operation itself. If the asynchronous
// operation succeeded, the response body includes the HTTP status code for the successful request. If the
// asynchronous operation failed, the response body includes the HTTP status code for the failed request
// and error information regarding the failure.
type AzureAsyncOperationResult struct {
	// Status - READ-ONLY; the status of the AzureAsyncOperation. Possible values include: 'OperationStatusInProgress', 'OperationStatusSucceeded', 'OperationStatusFailed'
	Status OperationStatus `json:"status,omitempty"`
	// Error - READ-ONLY
	Error *Error `json:"error,omitempty"`
}

// CreateFuture an abstraction for monitoring and retrieving the results of a long-running operation.
type CreateFuture struct {
	azure.FutureAPI
	// Result returns the result of the asynchronous operation.
	// If the operation has not completed it will return an error.
	Result func(Client) (DataLakeStoreAccount, error)
}

// DataLakeStoreAccount data Lake Store account information
type DataLakeStoreAccount struct {
	autorest.Response `json:"-"`
	// Location - the account regional location.
	Location *string `json:"location,omitempty"`
	// Name - the account name.
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; the namespace and type of the account.
	Type *string `json:"type,omitempty"`
	// ID - READ-ONLY; the account subscription ID.
	ID *string `json:"id,omitempty"`
	// Identity - The Key vault encryption identity, if any.
	Identity *EncryptionIdentity `json:"identity,omitempty"`
	// Tags - the value of custom properties.
	Tags map[string]*string `json:"tags"`
	// Properties - the Data Lake Store account properties.
	Properties *DataLakeStoreAccountProperties `json:"properties,omitempty"`
}

// MarshalJSON is the custom marshaler for DataLakeStoreAccount.
func (dlsa DataLakeStoreAccount) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if dlsa.Location != nil {
		objectMap["location"] = dlsa.Location
	}
	if dlsa.Name != nil {
		objectMap["name"] = dlsa.Name
	}
	if dlsa.Identity != nil {
		objectMap["identity"] = dlsa.Identity
	}
	if dlsa.Tags != nil {
		objectMap["tags"] = dlsa.Tags
	}
	if dlsa.Properties != nil {
		objectMap["properties"] = dlsa.Properties
	}
	return json.Marshal(objectMap)
}

// DataLakeStoreAccountListResult data Lake Store account list information response.
type DataLakeStoreAccountListResult struct {
	autorest.Response `json:"-"`
	// Value - READ-ONLY; the results of the list operation
	Value *[]DataLakeStoreAccount `json:"value,omitempty"`
	// NextLink - READ-ONLY; the link (url) to the next page of results.
	NextLink *string `json:"nextLink,omitempty"`
	// Count - READ-ONLY; the total count of results that are available, but might not be returned in the current page.
	Count *int64 `json:"count,omitempty"`
}

// DataLakeStoreAccountListResultIterator provides access to a complete listing of DataLakeStoreAccount
// values.
type DataLakeStoreAccountListResultIterator struct {
	i    int
	page DataLakeStoreAccountListResultPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *DataLakeStoreAccountListResultIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DataLakeStoreAccountListResultIterator.NextWithContext")
		defer func() {
			sc := -1
			if iter.Response().Response.Response != nil {
				sc = iter.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err = iter.page.NextWithContext(ctx)
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (iter *DataLakeStoreAccountListResultIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter DataLakeStoreAccountListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter DataLakeStoreAccountListResultIterator) Response() DataLakeStoreAccountListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter DataLakeStoreAccountListResultIterator) Value() DataLakeStoreAccount {
	if !iter.page.NotDone() {
		return DataLakeStoreAccount{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the DataLakeStoreAccountListResultIterator type.
func NewDataLakeStoreAccountListResultIterator(page DataLakeStoreAccountListResultPage) DataLakeStoreAccountListResultIterator {
	return DataLakeStoreAccountListResultIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (dlsalr DataLakeStoreAccountListResult) IsEmpty() bool {
	return dlsalr.Value == nil || len(*dlsalr.Value) == 0
}

// hasNextLink returns true if the NextLink is not empty.
func (dlsalr DataLakeStoreAccountListResult) hasNextLink() bool {
	return dlsalr.NextLink != nil && len(*dlsalr.NextLink) != 0
}

// dataLakeStoreAccountListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (dlsalr DataLakeStoreAccountListResult) dataLakeStoreAccountListResultPreparer(ctx context.Context) (*http.Request, error) {
	if !dlsalr.hasNextLink() {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(dlsalr.NextLink)))
}

// DataLakeStoreAccountListResultPage contains a page of DataLakeStoreAccount values.
type DataLakeStoreAccountListResultPage struct {
	fn     func(context.Context, DataLakeStoreAccountListResult) (DataLakeStoreAccountListResult, error)
	dlsalr DataLakeStoreAccountListResult
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *DataLakeStoreAccountListResultPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DataLakeStoreAccountListResultPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	for {
		next, err := page.fn(ctx, page.dlsalr)
		if err != nil {
			return err
		}
		page.dlsalr = next
		if !next.hasNextLink() || !next.IsEmpty() {
			break
		}
	}
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *DataLakeStoreAccountListResultPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page DataLakeStoreAccountListResultPage) NotDone() bool {
	return !page.dlsalr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page DataLakeStoreAccountListResultPage) Response() DataLakeStoreAccountListResult {
	return page.dlsalr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page DataLakeStoreAccountListResultPage) Values() []DataLakeStoreAccount {
	if page.dlsalr.IsEmpty() {
		return nil
	}
	return *page.dlsalr.Value
}

// Creates a new instance of the DataLakeStoreAccountListResultPage type.
func NewDataLakeStoreAccountListResultPage(cur DataLakeStoreAccountListResult, getNextPage func(context.Context, DataLakeStoreAccountListResult) (DataLakeStoreAccountListResult, error)) DataLakeStoreAccountListResultPage {
	return DataLakeStoreAccountListResultPage{
		fn:     getNextPage,
		dlsalr: cur,
	}
}

// DataLakeStoreAccountProperties data Lake Store account properties information
type DataLakeStoreAccountProperties struct {
	// ProvisioningState - READ-ONLY; the status of the Data Lake Store account while being provisioned. Possible values include: 'Failed', 'Creating', 'Running', 'Succeeded', 'Patching', 'Suspending', 'Resuming', 'Deleting', 'Deleted'
	ProvisioningState DataLakeStoreAccountStatus `json:"provisioningState,omitempty"`
	// State - READ-ONLY; the status of the Data Lake Store account after provisioning has completed. Possible values include: 'Active', 'Suspended'
	State DataLakeStoreAccountState `json:"state,omitempty"`
	// CreationTime - READ-ONLY; the account creation time.
	CreationTime *date.Time `json:"creationTime,omitempty"`
	// EncryptionState - The current state of encryption for this Data Lake store account. Possible values include: 'Enabled', 'Disabled'
	EncryptionState EncryptionState `json:"encryptionState,omitempty"`
	// EncryptionProvisioningState - READ-ONLY; The current state of encryption provisioning for this Data Lake store account. Possible values include: 'EncryptionProvisioningStateCreating', 'EncryptionProvisioningStateSucceeded'
	EncryptionProvisioningState EncryptionProvisioningState `json:"encryptionProvisioningState,omitempty"`
	// EncryptionConfig - The Key vault encryption configuration.
	EncryptionConfig *EncryptionConfig `json:"encryptionConfig,omitempty"`
	// LastModifiedTime - READ-ONLY; the account last modified time.
	LastModifiedTime *date.Time `json:"lastModifiedTime,omitempty"`
	// Endpoint - the gateway host.
	Endpoint *string `json:"endpoint,omitempty"`
	// DefaultGroup - the default owner group for all new folders and files created in the Data Lake Store account.
	DefaultGroup *string `json:"defaultGroup,omitempty"`
}

// MarshalJSON is the custom marshaler for DataLakeStoreAccountProperties.
func (dlsap DataLakeStoreAccountProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if dlsap.EncryptionState != "" {
		objectMap["encryptionState"] = dlsap.EncryptionState
	}
	if dlsap.EncryptionConfig != nil {
		objectMap["encryptionConfig"] = dlsap.EncryptionConfig
	}
	if dlsap.Endpoint != nil {
		objectMap["endpoint"] = dlsap.Endpoint
	}
	if dlsap.DefaultGroup != nil {
		objectMap["defaultGroup"] = dlsap.DefaultGroup
	}
	return json.Marshal(objectMap)
}

// DataLakeStoreFirewallRuleListResult data Lake Store firewall rule list information.
type DataLakeStoreFirewallRuleListResult struct {
	autorest.Response `json:"-"`
	// Value - READ-ONLY; the results of the list operation
	Value *[]FirewallRule `json:"value,omitempty"`
	// NextLink - READ-ONLY; the link (url) to the next page of results.
	NextLink *string `json:"nextLink,omitempty"`
	// Count - READ-ONLY; the total count of results that are available, but might not be returned in the current page.
	Count *int64 `json:"count,omitempty"`
}

// DataLakeStoreFirewallRuleListResultIterator provides access to a complete listing of FirewallRule
// values.
type DataLakeStoreFirewallRuleListResultIterator struct {
	i    int
	page DataLakeStoreFirewallRuleListResultPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *DataLakeStoreFirewallRuleListResultIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DataLakeStoreFirewallRuleListResultIterator.NextWithContext")
		defer func() {
			sc := -1
			if iter.Response().Response.Response != nil {
				sc = iter.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err = iter.page.NextWithContext(ctx)
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (iter *DataLakeStoreFirewallRuleListResultIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter DataLakeStoreFirewallRuleListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter DataLakeStoreFirewallRuleListResultIterator) Response() DataLakeStoreFirewallRuleListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter DataLakeStoreFirewallRuleListResultIterator) Value() FirewallRule {
	if !iter.page.NotDone() {
		return FirewallRule{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the DataLakeStoreFirewallRuleListResultIterator type.
func NewDataLakeStoreFirewallRuleListResultIterator(page DataLakeStoreFirewallRuleListResultPage) DataLakeStoreFirewallRuleListResultIterator {
	return DataLakeStoreFirewallRuleListResultIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (dlsfrlr DataLakeStoreFirewallRuleListResult) IsEmpty() bool {
	return dlsfrlr.Value == nil || len(*dlsfrlr.Value) == 0
}

// hasNextLink returns true if the NextLink is not empty.
func (dlsfrlr DataLakeStoreFirewallRuleListResult) hasNextLink() bool {
	return dlsfrlr.NextLink != nil && len(*dlsfrlr.NextLink) != 0
}

// dataLakeStoreFirewallRuleListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (dlsfrlr DataLakeStoreFirewallRuleListResult) dataLakeStoreFirewallRuleListResultPreparer(ctx context.Context) (*http.Request, error) {
	if !dlsfrlr.hasNextLink() {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(dlsfrlr.NextLink)))
}

// DataLakeStoreFirewallRuleListResultPage contains a page of FirewallRule values.
type DataLakeStoreFirewallRuleListResultPage struct {
	fn      func(context.Context, DataLakeStoreFirewallRuleListResult) (DataLakeStoreFirewallRuleListResult, error)
	dlsfrlr DataLakeStoreFirewallRuleListResult
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *DataLakeStoreFirewallRuleListResultPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DataLakeStoreFirewallRuleListResultPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	for {
		next, err := page.fn(ctx, page.dlsfrlr)
		if err != nil {
			return err
		}
		page.dlsfrlr = next
		if !next.hasNextLink() || !next.IsEmpty() {
			break
		}
	}
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *DataLakeStoreFirewallRuleListResultPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page DataLakeStoreFirewallRuleListResultPage) NotDone() bool {
	return !page.dlsfrlr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page DataLakeStoreFirewallRuleListResultPage) Response() DataLakeStoreFirewallRuleListResult {
	return page.dlsfrlr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page DataLakeStoreFirewallRuleListResultPage) Values() []FirewallRule {
	if page.dlsfrlr.IsEmpty() {
		return nil
	}
	return *page.dlsfrlr.Value
}

// Creates a new instance of the DataLakeStoreFirewallRuleListResultPage type.
func NewDataLakeStoreFirewallRuleListResultPage(cur DataLakeStoreFirewallRuleListResult, getNextPage func(context.Context, DataLakeStoreFirewallRuleListResult) (DataLakeStoreFirewallRuleListResult, error)) DataLakeStoreFirewallRuleListResultPage {
	return DataLakeStoreFirewallRuleListResultPage{
		fn:      getNextPage,
		dlsfrlr: cur,
	}
}

// DeleteFuture an abstraction for monitoring and retrieving the results of a long-running operation.
type DeleteFuture struct {
	azure.FutureAPI
	// Result returns the result of the asynchronous operation.
	// If the operation has not completed it will return an error.
	Result func(Client) (autorest.Response, error)
}

// EncryptionConfig ...
type EncryptionConfig struct {
	// Type - The type of encryption configuration being used. Currently the only supported types are 'UserManaged' and 'ServiceManaged'. Possible values include: 'UserManaged', 'ServiceManaged'
	Type EncryptionConfigType `json:"type,omitempty"`
	// KeyVaultMetaInfo - The Key vault information for connecting to user managed encryption keys.
	KeyVaultMetaInfo *KeyVaultMetaInfo `json:"keyVaultMetaInfo,omitempty"`
}

// EncryptionIdentity ...
type EncryptionIdentity struct {
	// Type - The type of encryption being used. Currently the only supported type is 'SystemAssigned'. Possible values include: 'SystemAssigned'
	Type EncryptionIdentityType `json:"type,omitempty"`
	// PrincipalID - READ-ONLY; The principal identifier associated with the encryption.
	PrincipalID *uuid.UUID `json:"principalId,omitempty"`
	// TenantID - READ-ONLY; The tenant identifier associated with the encryption.
	TenantID *uuid.UUID `json:"tenantId,omitempty"`
}

// MarshalJSON is the custom marshaler for EncryptionIdentity.
func (ei EncryptionIdentity) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if ei.Type != "" {
		objectMap["type"] = ei.Type
	}
	return json.Marshal(objectMap)
}

// Error data Lake Store error information
type Error struct {
	// Code - READ-ONLY; the HTTP status code or error code associated with this error
	Code *string `json:"code,omitempty"`
	// Message - READ-ONLY; the error message to display.
	Message *string `json:"message,omitempty"`
	// Target - READ-ONLY; the target of the error.
	Target *string `json:"target,omitempty"`
	// Details - READ-ONLY; the list of error details
	Details *[]ErrorDetails `json:"details,omitempty"`
	// InnerError - READ-ONLY; the inner exceptions or errors, if any
	InnerError *InnerError `json:"innerError,omitempty"`
}

// ErrorDetails data Lake Store error details information
type ErrorDetails struct {
	// Code - READ-ONLY; the HTTP status code or error code associated with this error
	Code *string `json:"code,omitempty"`
	// Message - READ-ONLY; the error message localized based on Accept-Language
	Message *string `json:"message,omitempty"`
	// Target - READ-ONLY; the target of the particular error (for example, the name of the property in error).
	Target *string `json:"target,omitempty"`
}

// FirewallRule data Lake Store firewall rule information
type FirewallRule struct {
	autorest.Response `json:"-"`
	// Name - the firewall rule's name.
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; the namespace and type of the firewall Rule.
	Type *string `json:"type,omitempty"`
	// ID - the firewall rule's subscription ID.
	ID *string `json:"id,omitempty"`
	// Location - the firewall rule's regional location.
	Location *string `json:"location,omitempty"`
	// Properties - the properties of the firewall rule.
	Properties *FirewallRuleProperties `json:"properties,omitempty"`
}

// MarshalJSON is the custom marshaler for FirewallRule.
func (fr FirewallRule) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if fr.Name != nil {
		objectMap["name"] = fr.Name
	}
	if fr.ID != nil {
		objectMap["id"] = fr.ID
	}
	if fr.Location != nil {
		objectMap["location"] = fr.Location
	}
	if fr.Properties != nil {
		objectMap["properties"] = fr.Properties
	}
	return json.Marshal(objectMap)
}

// FirewallRuleProperties data Lake Store firewall rule properties information
type FirewallRuleProperties struct {
	// StartIPAddress - the start IP address for the firewall rule.
	StartIPAddress *string `json:"startIpAddress,omitempty"`
	// EndIPAddress - the end IP address for the firewall rule.
	EndIPAddress *string `json:"endIpAddress,omitempty"`
}

// InnerError data Lake Store inner error information
type InnerError struct {
	// Trace - READ-ONLY; the stack trace for the error
	Trace *string `json:"trace,omitempty"`
	// Context - READ-ONLY; the context for the error message
	Context *string `json:"context,omitempty"`
}

// KeyVaultMetaInfo ...
type KeyVaultMetaInfo struct {
	// KeyVaultResourceID - The resource identifier for the user managed Key Vault being used to encrypt.
	KeyVaultResourceID *string `json:"keyVaultResourceId,omitempty"`
	// EncryptionKeyName - The name of the user managed encryption key.
	EncryptionKeyName *string `json:"encryptionKeyName,omitempty"`
	// EncryptionKeyVersion - The version of the user managed encryption key.
	EncryptionKeyVersion *string `json:"encryptionKeyVersion,omitempty"`
}

// UpdateFuture an abstraction for monitoring and retrieving the results of a long-running operation.
type UpdateFuture struct {
	azure.FutureAPI
	// Result returns the result of the asynchronous operation.
	// If the operation has not completed it will return an error.
	Result func(Client) (DataLakeStoreAccount, error)
}

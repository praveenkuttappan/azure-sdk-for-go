//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armm365securityandcompliance

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// PrivateLinkServicesForM365SecurityCenterClient contains the methods for the PrivateLinkServicesForM365SecurityCenter group.
// Don't use this type directly, use NewPrivateLinkServicesForM365SecurityCenterClient() instead.
type PrivateLinkServicesForM365SecurityCenterClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewPrivateLinkServicesForM365SecurityCenterClient creates a new instance of PrivateLinkServicesForM365SecurityCenterClient with the specified values.
// subscriptionID - The subscription identifier.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewPrivateLinkServicesForM365SecurityCenterClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*PrivateLinkServicesForM365SecurityCenterClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublicCloud.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &PrivateLinkServicesForM365SecurityCenterClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create or update the metadata of a privateLinkServicesForM365SecurityCenter instance.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the service instance.
// resourceName - The name of the service instance.
// privateLinkServicesForM365SecurityCenterDescription - The service instance metadata.
// options - PrivateLinkServicesForM365SecurityCenterClientBeginCreateOrUpdateOptions contains the optional parameters for
// the PrivateLinkServicesForM365SecurityCenterClient.BeginCreateOrUpdate method.
func (client *PrivateLinkServicesForM365SecurityCenterClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, resourceName string, privateLinkServicesForM365SecurityCenterDescription PrivateLinkServicesForM365SecurityCenterDescription, options *PrivateLinkServicesForM365SecurityCenterClientBeginCreateOrUpdateOptions) (*armruntime.Poller[PrivateLinkServicesForM365SecurityCenterClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, resourceName, privateLinkServicesForM365SecurityCenterDescription, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller(resp, client.pl, &armruntime.NewPollerOptions[PrivateLinkServicesForM365SecurityCenterClientCreateOrUpdateResponse]{
			FinalStateVia: armruntime.FinalStateViaLocation,
		})
	} else {
		return armruntime.NewPollerFromResumeToken[PrivateLinkServicesForM365SecurityCenterClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Create or update the metadata of a privateLinkServicesForM365SecurityCenter instance.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *PrivateLinkServicesForM365SecurityCenterClient) createOrUpdate(ctx context.Context, resourceGroupName string, resourceName string, privateLinkServicesForM365SecurityCenterDescription PrivateLinkServicesForM365SecurityCenterDescription, options *PrivateLinkServicesForM365SecurityCenterClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, resourceName, privateLinkServicesForM365SecurityCenterDescription, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *PrivateLinkServicesForM365SecurityCenterClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, privateLinkServicesForM365SecurityCenterDescription PrivateLinkServicesForM365SecurityCenterDescription, options *PrivateLinkServicesForM365SecurityCenterClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.M365SecurityAndCompliance/privateLinkServicesForM365SecurityCenter/{resourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-25-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, privateLinkServicesForM365SecurityCenterDescription)
}

// BeginDelete - Delete a service instance.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the service instance.
// resourceName - The name of the service instance.
// options - PrivateLinkServicesForM365SecurityCenterClientBeginDeleteOptions contains the optional parameters for the PrivateLinkServicesForM365SecurityCenterClient.BeginDelete
// method.
func (client *PrivateLinkServicesForM365SecurityCenterClient) BeginDelete(ctx context.Context, resourceGroupName string, resourceName string, options *PrivateLinkServicesForM365SecurityCenterClientBeginDeleteOptions) (*armruntime.Poller[PrivateLinkServicesForM365SecurityCenterClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, resourceName, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller(resp, client.pl, &armruntime.NewPollerOptions[PrivateLinkServicesForM365SecurityCenterClientDeleteResponse]{
			FinalStateVia: armruntime.FinalStateViaLocation,
		})
	} else {
		return armruntime.NewPollerFromResumeToken[PrivateLinkServicesForM365SecurityCenterClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Delete a service instance.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *PrivateLinkServicesForM365SecurityCenterClient) deleteOperation(ctx context.Context, resourceGroupName string, resourceName string, options *PrivateLinkServicesForM365SecurityCenterClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, resourceName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *PrivateLinkServicesForM365SecurityCenterClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, options *PrivateLinkServicesForM365SecurityCenterClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.M365SecurityAndCompliance/privateLinkServicesForM365SecurityCenter/{resourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-25-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Get the metadata of a privateLinkServicesForM365SecurityCenter resource.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the service instance.
// resourceName - The name of the service instance.
// options - PrivateLinkServicesForM365SecurityCenterClientGetOptions contains the optional parameters for the PrivateLinkServicesForM365SecurityCenterClient.Get
// method.
func (client *PrivateLinkServicesForM365SecurityCenterClient) Get(ctx context.Context, resourceGroupName string, resourceName string, options *PrivateLinkServicesForM365SecurityCenterClientGetOptions) (PrivateLinkServicesForM365SecurityCenterClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, resourceName, options)
	if err != nil {
		return PrivateLinkServicesForM365SecurityCenterClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PrivateLinkServicesForM365SecurityCenterClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PrivateLinkServicesForM365SecurityCenterClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *PrivateLinkServicesForM365SecurityCenterClient) getCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, options *PrivateLinkServicesForM365SecurityCenterClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.M365SecurityAndCompliance/privateLinkServicesForM365SecurityCenter/{resourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-25-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *PrivateLinkServicesForM365SecurityCenterClient) getHandleResponse(resp *http.Response) (PrivateLinkServicesForM365SecurityCenterClientGetResponse, error) {
	result := PrivateLinkServicesForM365SecurityCenterClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateLinkServicesForM365SecurityCenterDescription); err != nil {
		return PrivateLinkServicesForM365SecurityCenterClientGetResponse{}, err
	}
	return result, nil
}

// List - Get all the privateLinkServicesForM365SecurityCenter instances in a subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// options - PrivateLinkServicesForM365SecurityCenterClientListOptions contains the optional parameters for the PrivateLinkServicesForM365SecurityCenterClient.List
// method.
func (client *PrivateLinkServicesForM365SecurityCenterClient) List(options *PrivateLinkServicesForM365SecurityCenterClientListOptions) *runtime.Pager[PrivateLinkServicesForM365SecurityCenterClientListResponse] {
	return runtime.NewPager(runtime.PageProcessor[PrivateLinkServicesForM365SecurityCenterClientListResponse]{
		More: func(page PrivateLinkServicesForM365SecurityCenterClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *PrivateLinkServicesForM365SecurityCenterClientListResponse) (PrivateLinkServicesForM365SecurityCenterClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return PrivateLinkServicesForM365SecurityCenterClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return PrivateLinkServicesForM365SecurityCenterClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return PrivateLinkServicesForM365SecurityCenterClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *PrivateLinkServicesForM365SecurityCenterClient) listCreateRequest(ctx context.Context, options *PrivateLinkServicesForM365SecurityCenterClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.M365SecurityAndCompliance/privateLinkServicesForM365SecurityCenter"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-25-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *PrivateLinkServicesForM365SecurityCenterClient) listHandleResponse(resp *http.Response) (PrivateLinkServicesForM365SecurityCenterClientListResponse, error) {
	result := PrivateLinkServicesForM365SecurityCenterClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateLinkServicesForM365SecurityCenterDescriptionListResult); err != nil {
		return PrivateLinkServicesForM365SecurityCenterClientListResponse{}, err
	}
	return result, nil
}

// ListByResourceGroup - Get all the service instances in a resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the service instance.
// options - PrivateLinkServicesForM365SecurityCenterClientListByResourceGroupOptions contains the optional parameters for
// the PrivateLinkServicesForM365SecurityCenterClient.ListByResourceGroup method.
func (client *PrivateLinkServicesForM365SecurityCenterClient) ListByResourceGroup(resourceGroupName string, options *PrivateLinkServicesForM365SecurityCenterClientListByResourceGroupOptions) *runtime.Pager[PrivateLinkServicesForM365SecurityCenterClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PageProcessor[PrivateLinkServicesForM365SecurityCenterClientListByResourceGroupResponse]{
		More: func(page PrivateLinkServicesForM365SecurityCenterClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *PrivateLinkServicesForM365SecurityCenterClientListByResourceGroupResponse) (PrivateLinkServicesForM365SecurityCenterClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return PrivateLinkServicesForM365SecurityCenterClientListByResourceGroupResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return PrivateLinkServicesForM365SecurityCenterClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return PrivateLinkServicesForM365SecurityCenterClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *PrivateLinkServicesForM365SecurityCenterClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *PrivateLinkServicesForM365SecurityCenterClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.M365SecurityAndCompliance/privateLinkServicesForM365SecurityCenter"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-25-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *PrivateLinkServicesForM365SecurityCenterClient) listByResourceGroupHandleResponse(resp *http.Response) (PrivateLinkServicesForM365SecurityCenterClientListByResourceGroupResponse, error) {
	result := PrivateLinkServicesForM365SecurityCenterClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateLinkServicesForM365SecurityCenterDescriptionListResult); err != nil {
		return PrivateLinkServicesForM365SecurityCenterClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Update the metadata of a privateLinkServicesForM365SecurityCenter instance.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the service instance.
// resourceName - The name of the service instance.
// servicePatchDescription - The service instance metadata and security metadata.
// options - PrivateLinkServicesForM365SecurityCenterClientBeginUpdateOptions contains the optional parameters for the PrivateLinkServicesForM365SecurityCenterClient.BeginUpdate
// method.
func (client *PrivateLinkServicesForM365SecurityCenterClient) BeginUpdate(ctx context.Context, resourceGroupName string, resourceName string, servicePatchDescription ServicesPatchDescription, options *PrivateLinkServicesForM365SecurityCenterClientBeginUpdateOptions) (*armruntime.Poller[PrivateLinkServicesForM365SecurityCenterClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, resourceName, servicePatchDescription, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller(resp, client.pl, &armruntime.NewPollerOptions[PrivateLinkServicesForM365SecurityCenterClientUpdateResponse]{
			FinalStateVia: armruntime.FinalStateViaLocation,
		})
	} else {
		return armruntime.NewPollerFromResumeToken[PrivateLinkServicesForM365SecurityCenterClientUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Update - Update the metadata of a privateLinkServicesForM365SecurityCenter instance.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *PrivateLinkServicesForM365SecurityCenterClient) update(ctx context.Context, resourceGroupName string, resourceName string, servicePatchDescription ServicesPatchDescription, options *PrivateLinkServicesForM365SecurityCenterClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, resourceName, servicePatchDescription, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *PrivateLinkServicesForM365SecurityCenterClient) updateCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, servicePatchDescription ServicesPatchDescription, options *PrivateLinkServicesForM365SecurityCenterClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.M365SecurityAndCompliance/privateLinkServicesForM365SecurityCenter/{resourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-25-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, servicePatchDescription)
}

//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armvmwarecloudsimple

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// ResourcePoolsClient contains the methods for the ResourcePools group.
// Don't use this type directly, use NewResourcePoolsClient() instead.
type ResourcePoolsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewResourcePoolsClient creates a new instance of ResourcePoolsClient with the specified values.
func NewResourcePoolsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ResourcePoolsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &ResourcePoolsClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// Get - Returns resource pool templates by its name
// If the operation fails it returns the *CSRPError error type.
func (client *ResourcePoolsClient) Get(ctx context.Context, regionID string, pcName string, resourcePoolName string, options *ResourcePoolsGetOptions) (ResourcePoolsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, regionID, pcName, resourcePoolName, options)
	if err != nil {
		return ResourcePoolsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ResourcePoolsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ResourcePoolsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ResourcePoolsClient) getCreateRequest(ctx context.Context, regionID string, pcName string, resourcePoolName string, options *ResourcePoolsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.VMwareCloudSimple/locations/{regionId}/privateClouds/{pcName}/resourcePools/{resourcePoolName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if regionID == "" {
		return nil, errors.New("parameter regionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{regionId}", url.PathEscape(regionID))
	if pcName == "" {
		return nil, errors.New("parameter pcName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{pcName}", url.PathEscape(pcName))
	if resourcePoolName == "" {
		return nil, errors.New("parameter resourcePoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourcePoolName}", url.PathEscape(resourcePoolName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ResourcePoolsClient) getHandleResponse(resp *http.Response) (ResourcePoolsGetResponse, error) {
	result := ResourcePoolsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ResourcePool); err != nil {
		return ResourcePoolsGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *ResourcePoolsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CSRPError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// List - Returns list of resource pools in region for private cloud
// If the operation fails it returns the *CSRPError error type.
func (client *ResourcePoolsClient) List(regionID string, pcName string, options *ResourcePoolsListOptions) *ResourcePoolsListPager {
	return &ResourcePoolsListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, regionID, pcName, options)
		},
		advancer: func(ctx context.Context, resp ResourcePoolsListResponseEnvelope) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ResourcePoolsListResponse.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *ResourcePoolsClient) listCreateRequest(ctx context.Context, regionID string, pcName string, options *ResourcePoolsListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.VMwareCloudSimple/locations/{regionId}/privateClouds/{pcName}/resourcePools"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if regionID == "" {
		return nil, errors.New("parameter regionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{regionId}", url.PathEscape(regionID))
	if pcName == "" {
		return nil, errors.New("parameter pcName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{pcName}", url.PathEscape(pcName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ResourcePoolsClient) listHandleResponse(resp *http.Response) (ResourcePoolsListResponseEnvelope, error) {
	result := ResourcePoolsListResponseEnvelope{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ResourcePoolsListResponse); err != nil {
		return ResourcePoolsListResponseEnvelope{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *ResourcePoolsClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CSRPError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
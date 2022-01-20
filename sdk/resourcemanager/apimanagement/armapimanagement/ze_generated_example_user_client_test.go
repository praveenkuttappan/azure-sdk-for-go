//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armapimanagement_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement"
)

// x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementListUsers.json
func ExampleUserClient_ListByService() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armapimanagement.NewUserClient("<subscription-id>", cred, nil)
	pager := client.ListByService("<resource-group-name>",
		"<service-name>",
		&armapimanagement.UserClientListByServiceOptions{Filter: nil,
			Top:          nil,
			Skip:         nil,
			ExpandGroups: nil,
		})
	for {
		nextResult := pager.NextPage(ctx)
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		if !nextResult {
			break
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("Pager result: %#v\n", v)
		}
	}
}

// x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementHeadUser.json
func ExampleUserClient_GetEntityTag() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armapimanagement.NewUserClient("<subscription-id>", cred, nil)
	_, err = client.GetEntityTag(ctx,
		"<resource-group-name>",
		"<service-name>",
		"<user-id>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementGetUser.json
func ExampleUserClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armapimanagement.NewUserClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<service-name>",
		"<user-id>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.UserClientGetResult)
}

// x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementCreateUser.json
func ExampleUserClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armapimanagement.NewUserClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<service-name>",
		"<user-id>",
		armapimanagement.UserCreateParameters{
			Properties: &armapimanagement.UserCreateParameterProperties{
				Confirmation: armapimanagement.Confirmation("signup").ToPtr(),
				Email:        to.StringPtr("<email>"),
				FirstName:    to.StringPtr("<first-name>"),
				LastName:     to.StringPtr("<last-name>"),
			},
		},
		&armapimanagement.UserClientCreateOrUpdateOptions{Notify: nil,
			IfMatch: nil,
		})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.UserClientCreateOrUpdateResult)
}

// x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementUpdateUser.json
func ExampleUserClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armapimanagement.NewUserClient("<subscription-id>", cred, nil)
	res, err := client.Update(ctx,
		"<resource-group-name>",
		"<service-name>",
		"<user-id>",
		"<if-match>",
		armapimanagement.UserUpdateParameters{
			Properties: &armapimanagement.UserUpdateParametersProperties{
				Email:     to.StringPtr("<email>"),
				FirstName: to.StringPtr("<first-name>"),
				LastName:  to.StringPtr("<last-name>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.UserClientUpdateResult)
}

// x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementDeleteUser.json
func ExampleUserClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armapimanagement.NewUserClient("<subscription-id>", cred, nil)
	_, err = client.Delete(ctx,
		"<resource-group-name>",
		"<service-name>",
		"<user-id>",
		"<if-match>",
		&armapimanagement.UserClientDeleteOptions{DeleteSubscriptions: nil,
			Notify:  nil,
			AppType: nil,
		})
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementUserToken.json
func ExampleUserClient_GetSharedAccessToken() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armapimanagement.NewUserClient("<subscription-id>", cred, nil)
	res, err := client.GetSharedAccessToken(ctx,
		"<resource-group-name>",
		"<service-name>",
		"<user-id>",
		armapimanagement.UserTokenParameters{
			Properties: &armapimanagement.UserTokenParameterProperties{
				Expiry:  to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-04-21T00:44:24.2845269Z"); return t }()),
				KeyType: armapimanagement.KeyTypePrimary.ToPtr(),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.UserClientGetSharedAccessTokenResult)
}
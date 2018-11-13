package blueprintapi

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
	"github.com/Azure/azure-sdk-for-go/services/preview/blueprint/mgmt/2017-11-11-preview/blueprint"
)

// BlueprintsClientAPI contains the set of methods on the BlueprintsClient type.
type BlueprintsClientAPI interface {
	CreateOrUpdate(ctx context.Context, managementGroupName string, blueprintName string, blueprint blueprint.Model) (result blueprint.Model, err error)
	Delete(ctx context.Context, managementGroupName string, blueprintName string) (result blueprint.Model, err error)
	Get(ctx context.Context, managementGroupName string, blueprintName string) (result blueprint.Model, err error)
	List(ctx context.Context, managementGroupName string) (result blueprint.ListPage, err error)
}

var _ BlueprintsClientAPI = (*blueprint.BlueprintsClient)(nil)

// ArtifactsClientAPI contains the set of methods on the ArtifactsClient type.
type ArtifactsClientAPI interface {
	CreateOrUpdate(ctx context.Context, managementGroupName string, blueprintName string, artifactName string, artifact blueprint.BasicArtifact) (result blueprint.ArtifactModel, err error)
	Delete(ctx context.Context, managementGroupName string, blueprintName string, artifactName string) (result blueprint.ArtifactModel, err error)
	Get(ctx context.Context, managementGroupName string, blueprintName string, artifactName string) (result blueprint.ArtifactModel, err error)
	List(ctx context.Context, managementGroupName string, blueprintName string) (result blueprint.ArtifactListPage, err error)
}

var _ ArtifactsClientAPI = (*blueprint.ArtifactsClient)(nil)

// PublishedBlueprintsClientAPI contains the set of methods on the PublishedBlueprintsClient type.
type PublishedBlueprintsClientAPI interface {
	Create(ctx context.Context, managementGroupName string, blueprintName string, versionID string) (result blueprint.PublishedBlueprint, err error)
	Delete(ctx context.Context, managementGroupName string, blueprintName string, versionID string) (result blueprint.PublishedBlueprint, err error)
	Get(ctx context.Context, managementGroupName string, blueprintName string, versionID string) (result blueprint.PublishedBlueprint, err error)
	List(ctx context.Context, managementGroupName string, blueprintName string) (result blueprint.PublishedBlueprintListPage, err error)
}

var _ PublishedBlueprintsClientAPI = (*blueprint.PublishedBlueprintsClient)(nil)

// PublishedArtifactsClientAPI contains the set of methods on the PublishedArtifactsClient type.
type PublishedArtifactsClientAPI interface {
	Get(ctx context.Context, managementGroupName string, blueprintName string, versionID string, artifactName string) (result blueprint.ArtifactModel, err error)
	List(ctx context.Context, managementGroupName string, blueprintName string, versionID string) (result blueprint.ArtifactListPage, err error)
}

var _ PublishedArtifactsClientAPI = (*blueprint.PublishedArtifactsClient)(nil)

// AssignmentsClientAPI contains the set of methods on the AssignmentsClient type.
type AssignmentsClientAPI interface {
	CreateOrUpdate(ctx context.Context, subscriptionID string, assignmentName string, assignment blueprint.Assignment) (result blueprint.Assignment, err error)
	Delete(ctx context.Context, subscriptionID string, assignmentName string) (result blueprint.Assignment, err error)
	Get(ctx context.Context, subscriptionID string, assignmentName string) (result blueprint.Assignment, err error)
	List(ctx context.Context, subscriptionID string) (result blueprint.AssignmentListPage, err error)
}

var _ AssignmentsClientAPI = (*blueprint.AssignmentsClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result blueprint.ResourceProviderOperationList, err error)
}

var _ OperationsClientAPI = (*blueprint.OperationsClient)(nil)

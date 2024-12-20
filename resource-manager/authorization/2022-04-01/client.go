package v2022_04_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/resource-manager/authorization/2022-04-01/denyassignments"
	"github.com/redtenant/go-azure-sdk/resource-manager/authorization/2022-04-01/permissions"
	"github.com/redtenant/go-azure-sdk/resource-manager/authorization/2022-04-01/provideroperationsmetadata"
	"github.com/redtenant/go-azure-sdk/resource-manager/authorization/2022-04-01/roleassignments"
	"github.com/redtenant/go-azure-sdk/resource-manager/authorization/2022-04-01/roledefinitions"
	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

type Client struct {
	DenyAssignments            *denyassignments.DenyAssignmentsClient
	Permissions                *permissions.PermissionsClient
	ProviderOperationsMetadata *provideroperationsmetadata.ProviderOperationsMetadataClient
	RoleAssignments            *roleassignments.RoleAssignmentsClient
	RoleDefinitions            *roledefinitions.RoleDefinitionsClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	denyAssignmentsClient, err := denyassignments.NewDenyAssignmentsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DenyAssignments client: %+v", err)
	}
	configureFunc(denyAssignmentsClient.Client)

	permissionsClient, err := permissions.NewPermissionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Permissions client: %+v", err)
	}
	configureFunc(permissionsClient.Client)

	providerOperationsMetadataClient, err := provideroperationsmetadata.NewProviderOperationsMetadataClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ProviderOperationsMetadata client: %+v", err)
	}
	configureFunc(providerOperationsMetadataClient.Client)

	roleAssignmentsClient, err := roleassignments.NewRoleAssignmentsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleAssignments client: %+v", err)
	}
	configureFunc(roleAssignmentsClient.Client)

	roleDefinitionsClient, err := roledefinitions.NewRoleDefinitionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleDefinitions client: %+v", err)
	}
	configureFunc(roleDefinitionsClient.Client)

	return &Client{
		DenyAssignments:            denyAssignmentsClient,
		Permissions:                permissionsClient,
		ProviderOperationsMetadata: providerOperationsMetadataClient,
		RoleAssignments:            roleAssignmentsClient,
		RoleDefinitions:            roleDefinitionsClient,
	}, nil
}

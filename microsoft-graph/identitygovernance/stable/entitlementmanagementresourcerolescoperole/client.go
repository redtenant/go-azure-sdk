package entitlementmanagementresourcerolescoperole

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementResourceRoleScopeRoleClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementResourceRoleScopeRoleClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementResourceRoleScopeRoleClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementresourcerolescoperole", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementResourceRoleScopeRoleClient: %+v", err)
	}

	return &EntitlementManagementResourceRoleScopeRoleClient{
		Client: client,
	}, nil
}

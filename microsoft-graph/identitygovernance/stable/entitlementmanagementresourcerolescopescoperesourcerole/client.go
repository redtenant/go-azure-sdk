package entitlementmanagementresourcerolescopescoperesourcerole

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementResourceRoleScopeScopeResourceRoleClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementResourceRoleScopeScopeResourceRoleClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementResourceRoleScopeScopeResourceRoleClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementresourcerolescopescoperesourcerole", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementResourceRoleScopeScopeResourceRoleClient: %+v", err)
	}

	return &EntitlementManagementResourceRoleScopeScopeResourceRoleClient{
		Client: client,
	}, nil
}

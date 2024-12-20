package entitlementmanagementresourcerolescopescoperesourceroleresourceenvironment

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementResourceRoleScopeScopeResourceRoleResourceEnvironmentClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementResourceRoleScopeScopeResourceRoleResourceEnvironmentClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementResourceRoleScopeScopeResourceRoleResourceEnvironmentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementresourcerolescopescoperesourceroleresourceenvironment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementResourceRoleScopeScopeResourceRoleResourceEnvironmentClient: %+v", err)
	}

	return &EntitlementManagementResourceRoleScopeScopeResourceRoleResourceEnvironmentClient{
		Client: client,
	}, nil
}

package entitlementmanagementcatalogresourceroleresourcescoperesourceenvironment

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementCatalogResourceRoleResourceScopeResourceEnvironmentClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementCatalogResourceRoleResourceScopeResourceEnvironmentClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementCatalogResourceRoleResourceScopeResourceEnvironmentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementcatalogresourceroleresourcescoperesourceenvironment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementCatalogResourceRoleResourceScopeResourceEnvironmentClient: %+v", err)
	}

	return &EntitlementManagementCatalogResourceRoleResourceScopeResourceEnvironmentClient{
		Client: client,
	}, nil
}

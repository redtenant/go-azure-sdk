package entitlementmanagementcatalogresourceroleresourcescoperesource

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementCatalogResourceRoleResourceScopeResourceClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementCatalogResourceRoleResourceScopeResourceClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementCatalogResourceRoleResourceScopeResourceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementcatalogresourceroleresourcescoperesource", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementCatalogResourceRoleResourceScopeResourceClient: %+v", err)
	}

	return &EntitlementManagementCatalogResourceRoleResourceScopeResourceClient{
		Client: client,
	}, nil
}

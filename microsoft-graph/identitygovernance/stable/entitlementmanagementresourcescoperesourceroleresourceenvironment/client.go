package entitlementmanagementresourcescoperesourceroleresourceenvironment

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementResourceScopeResourceRoleResourceEnvironmentClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementResourceScopeResourceRoleResourceEnvironmentClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementResourceScopeResourceRoleResourceEnvironmentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementresourcescoperesourceroleresourceenvironment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementResourceScopeResourceRoleResourceEnvironmentClient: %+v", err)
	}

	return &EntitlementManagementResourceScopeResourceRoleResourceEnvironmentClient{
		Client: client,
	}, nil
}

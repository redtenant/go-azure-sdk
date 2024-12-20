package entitlementmanagementaccesspackageresourcerolescopescoperesourceenvironment

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementAccessPackageResourceRoleScopeScopeResourceEnvironmentClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementAccessPackageResourceRoleScopeScopeResourceEnvironmentClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementAccessPackageResourceRoleScopeScopeResourceEnvironmentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementaccesspackageresourcerolescopescoperesourceenvironment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementAccessPackageResourceRoleScopeScopeResourceEnvironmentClient: %+v", err)
	}

	return &EntitlementManagementAccessPackageResourceRoleScopeScopeResourceEnvironmentClient{
		Client: client,
	}, nil
}

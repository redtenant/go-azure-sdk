package entitlementmanagementaccesspackageresourcerolescoperoleresourceenvironment

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementAccessPackageResourceRoleScopeRoleResourceEnvironmentClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementAccessPackageResourceRoleScopeRoleResourceEnvironmentClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementAccessPackageResourceRoleScopeRoleResourceEnvironmentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementaccesspackageresourcerolescoperoleresourceenvironment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementAccessPackageResourceRoleScopeRoleResourceEnvironmentClient: %+v", err)
	}

	return &EntitlementManagementAccessPackageResourceRoleScopeRoleResourceEnvironmentClient{
		Client: client,
	}, nil
}

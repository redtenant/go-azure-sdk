package entitlementmanagementaccesspackageresourcerolescopeaccesspackageresourcescopeaccesspackageresourceaccesspackageresourceroleaccesspackageresourceaccesspackageresourceenvironment

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementAccessPackageResourceRoleScopeAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceEnvironmentClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementAccessPackageResourceRoleScopeAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceEnvironmentClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementAccessPackageResourceRoleScopeAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceEnvironmentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementaccesspackageresourcerolescopeaccesspackageresourcescopeaccesspackageresourceaccesspackageresourceroleaccesspackageresourceaccesspackageresourceenvironment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementAccessPackageResourceRoleScopeAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceEnvironmentClient: %+v", err)
	}

	return &EntitlementManagementAccessPackageResourceRoleScopeAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRoleAccessPackageResourceAccessPackageResourceEnvironmentClient{
		Client: client,
	}, nil
}

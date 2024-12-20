package entitlementmanagementaccesspackageresourcerolescopeaccesspackageresourcescopeaccesspackageresourceaccesspackageresourcerole

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementAccessPackageResourceRoleScopeAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRoleClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementAccessPackageResourceRoleScopeAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRoleClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementAccessPackageResourceRoleScopeAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRoleClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementaccesspackageresourcerolescopeaccesspackageresourcescopeaccesspackageresourceaccesspackageresourcerole", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementAccessPackageResourceRoleScopeAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRoleClient: %+v", err)
	}

	return &EntitlementManagementAccessPackageResourceRoleScopeAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRoleClient{
		Client: client,
	}, nil
}

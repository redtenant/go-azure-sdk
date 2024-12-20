package entitlementmanagementaccesspackageresourceaccesspackageresourcescopeaccesspackageresourceaccesspackageresourceenvironment

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementAccessPackageResourceAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceEnvironmentClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementAccessPackageResourceAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceEnvironmentClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementAccessPackageResourceAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceEnvironmentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementaccesspackageresourceaccesspackageresourcescopeaccesspackageresourceaccesspackageresourceenvironment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementAccessPackageResourceAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceEnvironmentClient: %+v", err)
	}

	return &EntitlementManagementAccessPackageResourceAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceEnvironmentClient{
		Client: client,
	}, nil
}

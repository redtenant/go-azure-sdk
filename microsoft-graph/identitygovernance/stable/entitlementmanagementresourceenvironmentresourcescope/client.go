package entitlementmanagementresourceenvironmentresourcescope

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementResourceEnvironmentResourceScopeClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementResourceEnvironmentResourceScopeClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementResourceEnvironmentResourceScopeClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementresourceenvironmentresourcescope", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementResourceEnvironmentResourceScopeClient: %+v", err)
	}

	return &EntitlementManagementResourceEnvironmentResourceScopeClient{
		Client: client,
	}, nil
}

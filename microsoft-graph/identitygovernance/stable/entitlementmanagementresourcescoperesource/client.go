package entitlementmanagementresourcescoperesource

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementResourceScopeResourceClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementResourceScopeResourceClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementResourceScopeResourceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementresourcescoperesource", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementResourceScopeResourceClient: %+v", err)
	}

	return &EntitlementManagementResourceScopeResourceClient{
		Client: client,
	}, nil
}

package entitlementmanagementresourcerequestresourcescoperesource

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementResourceRequestResourceScopeResourceClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementResourceRequestResourceScopeResourceClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementResourceRequestResourceScopeResourceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementresourcerequestresourcescoperesource", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementResourceRequestResourceScopeResourceClient: %+v", err)
	}

	return &EntitlementManagementResourceRequestResourceScopeResourceClient{
		Client: client,
	}, nil
}

package entitlementmanagementassignmentpolicycatalog

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementAssignmentPolicyCatalogClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementAssignmentPolicyCatalogClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementAssignmentPolicyCatalogClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementassignmentpolicycatalog", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementAssignmentPolicyCatalogClient: %+v", err)
	}

	return &EntitlementManagementAssignmentPolicyCatalogClient{
		Client: client,
	}, nil
}

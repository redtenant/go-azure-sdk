package entitlementmanagementassignment

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementAssignmentClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementAssignmentClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementAssignmentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementassignment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementAssignmentClient: %+v", err)
	}

	return &EntitlementManagementAssignmentClient{
		Client: client,
	}, nil
}

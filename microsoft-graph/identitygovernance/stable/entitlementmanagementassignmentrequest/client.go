package entitlementmanagementassignmentrequest

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementAssignmentRequestClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementAssignmentRequestClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementAssignmentRequestClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementassignmentrequest", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementAssignmentRequestClient: %+v", err)
	}

	return &EntitlementManagementAssignmentRequestClient{
		Client: client,
	}, nil
}

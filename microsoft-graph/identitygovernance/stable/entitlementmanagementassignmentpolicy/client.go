package entitlementmanagementassignmentpolicy

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementAssignmentPolicyClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementAssignmentPolicyClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementAssignmentPolicyClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementassignmentpolicy", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementAssignmentPolicyClient: %+v", err)
	}

	return &EntitlementManagementAssignmentPolicyClient{
		Client: client,
	}, nil
}

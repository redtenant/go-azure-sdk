package entitlementmanagementroleassignmentschedulerequestactivatedusing

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementRoleAssignmentScheduleRequestActivatedUsingClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementRoleAssignmentScheduleRequestActivatedUsingClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementRoleAssignmentScheduleRequestActivatedUsingClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementroleassignmentschedulerequestactivatedusing", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementRoleAssignmentScheduleRequestActivatedUsingClient: %+v", err)
	}

	return &EntitlementManagementRoleAssignmentScheduleRequestActivatedUsingClient{
		Client: client,
	}, nil
}

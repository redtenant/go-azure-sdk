package entitlementmanagementroleeligibilityschedulerequestroledefinition

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementRoleEligibilityScheduleRequestRoleDefinitionClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementRoleEligibilityScheduleRequestRoleDefinitionClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementRoleEligibilityScheduleRequestRoleDefinitionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementroleeligibilityschedulerequestroledefinition", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementRoleEligibilityScheduleRequestRoleDefinitionClient: %+v", err)
	}

	return &EntitlementManagementRoleEligibilityScheduleRequestRoleDefinitionClient{
		Client: client,
	}, nil
}

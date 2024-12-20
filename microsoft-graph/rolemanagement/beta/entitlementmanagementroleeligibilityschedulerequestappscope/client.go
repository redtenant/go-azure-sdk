package entitlementmanagementroleeligibilityschedulerequestappscope

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementRoleEligibilityScheduleRequestAppScopeClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementRoleEligibilityScheduleRequestAppScopeClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementRoleEligibilityScheduleRequestAppScopeClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementroleeligibilityschedulerequestappscope", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementRoleEligibilityScheduleRequestAppScopeClient: %+v", err)
	}

	return &EntitlementManagementRoleEligibilityScheduleRequestAppScopeClient{
		Client: client,
	}, nil
}

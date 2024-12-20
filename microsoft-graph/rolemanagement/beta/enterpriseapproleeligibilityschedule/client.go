package enterpriseapproleeligibilityschedule

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EnterpriseAppRoleEligibilityScheduleClient struct {
	Client *msgraph.Client
}

func NewEnterpriseAppRoleEligibilityScheduleClientWithBaseURI(sdkApi sdkEnv.Api) (*EnterpriseAppRoleEligibilityScheduleClient, error) {
	client, err := msgraph.NewClient(sdkApi, "enterpriseapproleeligibilityschedule", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EnterpriseAppRoleEligibilityScheduleClient: %+v", err)
	}

	return &EnterpriseAppRoleEligibilityScheduleClient{
		Client: client,
	}, nil
}

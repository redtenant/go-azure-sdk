package privilegedaccessgroupeligibilityscheduleinstance

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegedAccessGroupEligibilityScheduleInstanceClient struct {
	Client *msgraph.Client
}

func NewPrivilegedAccessGroupEligibilityScheduleInstanceClientWithBaseURI(sdkApi sdkEnv.Api) (*PrivilegedAccessGroupEligibilityScheduleInstanceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "privilegedaccessgroupeligibilityscheduleinstance", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PrivilegedAccessGroupEligibilityScheduleInstanceClient: %+v", err)
	}

	return &PrivilegedAccessGroupEligibilityScheduleInstanceClient{
		Client: client,
	}, nil
}

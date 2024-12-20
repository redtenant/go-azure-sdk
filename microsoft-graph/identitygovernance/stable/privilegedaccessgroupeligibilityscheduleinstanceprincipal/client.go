package privilegedaccessgroupeligibilityscheduleinstanceprincipal

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegedAccessGroupEligibilityScheduleInstancePrincipalClient struct {
	Client *msgraph.Client
}

func NewPrivilegedAccessGroupEligibilityScheduleInstancePrincipalClientWithBaseURI(sdkApi sdkEnv.Api) (*PrivilegedAccessGroupEligibilityScheduleInstancePrincipalClient, error) {
	client, err := msgraph.NewClient(sdkApi, "privilegedaccessgroupeligibilityscheduleinstanceprincipal", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PrivilegedAccessGroupEligibilityScheduleInstancePrincipalClient: %+v", err)
	}

	return &PrivilegedAccessGroupEligibilityScheduleInstancePrincipalClient{
		Client: client,
	}, nil
}

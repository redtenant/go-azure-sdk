package privilegedaccessgroupassignmentscheduleinstanceactivatedusing

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegedAccessGroupAssignmentScheduleInstanceActivatedUsingClient struct {
	Client *msgraph.Client
}

func NewPrivilegedAccessGroupAssignmentScheduleInstanceActivatedUsingClientWithBaseURI(sdkApi sdkEnv.Api) (*PrivilegedAccessGroupAssignmentScheduleInstanceActivatedUsingClient, error) {
	client, err := msgraph.NewClient(sdkApi, "privilegedaccessgroupassignmentscheduleinstanceactivatedusing", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PrivilegedAccessGroupAssignmentScheduleInstanceActivatedUsingClient: %+v", err)
	}

	return &PrivilegedAccessGroupAssignmentScheduleInstanceActivatedUsingClient{
		Client: client,
	}, nil
}

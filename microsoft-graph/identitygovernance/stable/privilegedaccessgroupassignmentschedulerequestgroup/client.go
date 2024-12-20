package privilegedaccessgroupassignmentschedulerequestgroup

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegedAccessGroupAssignmentScheduleRequestGroupClient struct {
	Client *msgraph.Client
}

func NewPrivilegedAccessGroupAssignmentScheduleRequestGroupClientWithBaseURI(sdkApi sdkEnv.Api) (*PrivilegedAccessGroupAssignmentScheduleRequestGroupClient, error) {
	client, err := msgraph.NewClient(sdkApi, "privilegedaccessgroupassignmentschedulerequestgroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PrivilegedAccessGroupAssignmentScheduleRequestGroupClient: %+v", err)
	}

	return &PrivilegedAccessGroupAssignmentScheduleRequestGroupClient{
		Client: client,
	}, nil
}

package enterpriseapproleassignmentschedulerequest

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EnterpriseAppRoleAssignmentScheduleRequestClient struct {
	Client *msgraph.Client
}

func NewEnterpriseAppRoleAssignmentScheduleRequestClientWithBaseURI(sdkApi sdkEnv.Api) (*EnterpriseAppRoleAssignmentScheduleRequestClient, error) {
	client, err := msgraph.NewClient(sdkApi, "enterpriseapproleassignmentschedulerequest", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EnterpriseAppRoleAssignmentScheduleRequestClient: %+v", err)
	}

	return &EnterpriseAppRoleAssignmentScheduleRequestClient{
		Client: client,
	}, nil
}

package enterpriseapproleassignmentschedulerequestroledefinition

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EnterpriseAppRoleAssignmentScheduleRequestRoleDefinitionClient struct {
	Client *msgraph.Client
}

func NewEnterpriseAppRoleAssignmentScheduleRequestRoleDefinitionClientWithBaseURI(sdkApi sdkEnv.Api) (*EnterpriseAppRoleAssignmentScheduleRequestRoleDefinitionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "enterpriseapproleassignmentschedulerequestroledefinition", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EnterpriseAppRoleAssignmentScheduleRequestRoleDefinitionClient: %+v", err)
	}

	return &EnterpriseAppRoleAssignmentScheduleRequestRoleDefinitionClient{
		Client: client,
	}, nil
}

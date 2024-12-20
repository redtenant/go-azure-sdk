package enterpriseapproleassignmentscheduleroledefinition

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EnterpriseAppRoleAssignmentScheduleRoleDefinitionClient struct {
	Client *msgraph.Client
}

func NewEnterpriseAppRoleAssignmentScheduleRoleDefinitionClientWithBaseURI(sdkApi sdkEnv.Api) (*EnterpriseAppRoleAssignmentScheduleRoleDefinitionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "enterpriseapproleassignmentscheduleroledefinition", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EnterpriseAppRoleAssignmentScheduleRoleDefinitionClient: %+v", err)
	}

	return &EnterpriseAppRoleAssignmentScheduleRoleDefinitionClient{
		Client: client,
	}, nil
}

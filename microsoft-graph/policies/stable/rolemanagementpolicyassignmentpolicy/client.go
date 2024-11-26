package rolemanagementpolicyassignmentpolicy

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleManagementPolicyAssignmentPolicyClient struct {
	Client *msgraph.Client
}

func NewRoleManagementPolicyAssignmentPolicyClientWithBaseURI(sdkApi sdkEnv.Api) (*RoleManagementPolicyAssignmentPolicyClient, error) {
	client, err := msgraph.NewClient(sdkApi, "rolemanagementpolicyassignmentpolicy", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleManagementPolicyAssignmentPolicyClient: %+v", err)
	}

	return &RoleManagementPolicyAssignmentPolicyClient{
		Client: client,
	}, nil
}

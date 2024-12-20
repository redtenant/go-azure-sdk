package entitlementmanagementroleassignmentapprovalstep

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementRoleAssignmentApprovalStepClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementRoleAssignmentApprovalStepClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementRoleAssignmentApprovalStepClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementroleassignmentapprovalstep", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementRoleAssignmentApprovalStepClient: %+v", err)
	}

	return &EntitlementManagementRoleAssignmentApprovalStepClient{
		Client: client,
	}, nil
}

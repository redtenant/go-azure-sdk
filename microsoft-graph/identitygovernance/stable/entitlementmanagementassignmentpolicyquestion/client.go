package entitlementmanagementassignmentpolicyquestion

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementAssignmentPolicyQuestionClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementAssignmentPolicyQuestionClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementAssignmentPolicyQuestionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementassignmentpolicyquestion", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementAssignmentPolicyQuestionClient: %+v", err)
	}

	return &EntitlementManagementAssignmentPolicyQuestionClient{
		Client: client,
	}, nil
}

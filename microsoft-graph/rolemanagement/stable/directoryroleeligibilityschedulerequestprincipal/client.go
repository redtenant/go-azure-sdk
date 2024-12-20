package directoryroleeligibilityschedulerequestprincipal

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DirectoryRoleEligibilityScheduleRequestPrincipalClient struct {
	Client *msgraph.Client
}

func NewDirectoryRoleEligibilityScheduleRequestPrincipalClientWithBaseURI(sdkApi sdkEnv.Api) (*DirectoryRoleEligibilityScheduleRequestPrincipalClient, error) {
	client, err := msgraph.NewClient(sdkApi, "directoryroleeligibilityschedulerequestprincipal", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DirectoryRoleEligibilityScheduleRequestPrincipalClient: %+v", err)
	}

	return &DirectoryRoleEligibilityScheduleRequestPrincipalClient{
		Client: client,
	}, nil
}

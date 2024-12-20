package roleassignmentschedules

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleAssignmentSchedulesClient struct {
	Client *resourcemanager.Client
}

func NewRoleAssignmentSchedulesClientWithBaseURI(sdkApi sdkEnv.Api) (*RoleAssignmentSchedulesClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "roleassignmentschedules", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleAssignmentSchedulesClient: %+v", err)
	}

	return &RoleAssignmentSchedulesClient{
		Client: client,
	}, nil
}

package devicemanagementroleassignmentdirectoryscope

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementRoleAssignmentDirectoryScopeClient struct {
	Client *msgraph.Client
}

func NewDeviceManagementRoleAssignmentDirectoryScopeClientWithBaseURI(sdkApi sdkEnv.Api) (*DeviceManagementRoleAssignmentDirectoryScopeClient, error) {
	client, err := msgraph.NewClient(sdkApi, "devicemanagementroleassignmentdirectoryscope", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DeviceManagementRoleAssignmentDirectoryScopeClient: %+v", err)
	}

	return &DeviceManagementRoleAssignmentDirectoryScopeClient{
		Client: client,
	}, nil
}

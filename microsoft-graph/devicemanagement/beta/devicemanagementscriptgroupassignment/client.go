package devicemanagementscriptgroupassignment

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementScriptGroupAssignmentClient struct {
	Client *msgraph.Client
}

func NewDeviceManagementScriptGroupAssignmentClientWithBaseURI(sdkApi sdkEnv.Api) (*DeviceManagementScriptGroupAssignmentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "devicemanagementscriptgroupassignment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DeviceManagementScriptGroupAssignmentClient: %+v", err)
	}

	return &DeviceManagementScriptGroupAssignmentClient{
		Client: client,
	}, nil
}

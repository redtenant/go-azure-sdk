package windowsautopilotdeploymentprofileassigneddevice

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsAutopilotDeploymentProfileAssignedDeviceClient struct {
	Client *msgraph.Client
}

func NewWindowsAutopilotDeploymentProfileAssignedDeviceClientWithBaseURI(sdkApi sdkEnv.Api) (*WindowsAutopilotDeploymentProfileAssignedDeviceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "windowsautopilotdeploymentprofileassigneddevice", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating WindowsAutopilotDeploymentProfileAssignedDeviceClient: %+v", err)
	}

	return &WindowsAutopilotDeploymentProfileAssignedDeviceClient{
		Client: client,
	}, nil
}

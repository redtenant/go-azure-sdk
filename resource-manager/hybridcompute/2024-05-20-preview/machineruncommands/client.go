package machineruncommands

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MachineRunCommandsClient struct {
	Client *resourcemanager.Client
}

func NewMachineRunCommandsClientWithBaseURI(sdkApi sdkEnv.Api) (*MachineRunCommandsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "machineruncommands", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MachineRunCommandsClient: %+v", err)
	}

	return &MachineRunCommandsClient{
		Client: client,
	}, nil
}

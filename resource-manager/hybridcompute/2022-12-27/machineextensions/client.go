package machineextensions

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MachineExtensionsClient struct {
	Client *resourcemanager.Client
}

func NewMachineExtensionsClientWithBaseURI(sdkApi sdkEnv.Api) (*MachineExtensionsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "machineextensions", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MachineExtensionsClient: %+v", err)
	}

	return &MachineExtensionsClient{
		Client: client,
	}, nil
}

package managednetwork

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedNetworkClient struct {
	Client *resourcemanager.Client
}

func NewManagedNetworkClientWithBaseURI(sdkApi sdkEnv.Api) (*ManagedNetworkClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "managednetwork", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ManagedNetworkClient: %+v", err)
	}

	return &ManagedNetworkClient{
		Client: client,
	}, nil
}

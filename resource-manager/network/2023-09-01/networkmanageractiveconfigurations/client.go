package networkmanageractiveconfigurations

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkManagerActiveConfigurationsClient struct {
	Client *resourcemanager.Client
}

func NewNetworkManagerActiveConfigurationsClientWithBaseURI(sdkApi sdkEnv.Api) (*NetworkManagerActiveConfigurationsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "networkmanageractiveconfigurations", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating NetworkManagerActiveConfigurationsClient: %+v", err)
	}

	return &NetworkManagerActiveConfigurationsClient{
		Client: client,
	}, nil
}

package virtualnetworkgateways

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualNetworkGatewaysClient struct {
	Client *resourcemanager.Client
}

func NewVirtualNetworkGatewaysClientWithBaseURI(sdkApi sdkEnv.Api) (*VirtualNetworkGatewaysClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "virtualnetworkgateways", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating VirtualNetworkGatewaysClient: %+v", err)
	}

	return &VirtualNetworkGatewaysClient{
		Client: client,
	}, nil
}

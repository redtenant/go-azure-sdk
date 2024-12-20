package virtualnetworkgatewayconnections

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualNetworkGatewayConnectionsClient struct {
	Client *resourcemanager.Client
}

func NewVirtualNetworkGatewayConnectionsClientWithBaseURI(sdkApi sdkEnv.Api) (*VirtualNetworkGatewayConnectionsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "virtualnetworkgatewayconnections", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating VirtualNetworkGatewayConnectionsClient: %+v", err)
	}

	return &VirtualNetworkGatewayConnectionsClient{
		Client: client,
	}, nil
}

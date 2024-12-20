package v2016_06_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/resource-manager/web/2016-06-01/connectiongateways"
	"github.com/redtenant/go-azure-sdk/resource-manager/web/2016-06-01/connections"
	"github.com/redtenant/go-azure-sdk/resource-manager/web/2016-06-01/customapis"
	"github.com/redtenant/go-azure-sdk/resource-manager/web/2016-06-01/managedapis"
	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

type Client struct {
	ConnectionGateways *connectiongateways.ConnectionGatewaysClient
	Connections        *connections.ConnectionsClient
	CustomAPIs         *customapis.CustomAPIsClient
	ManagedAPIs        *managedapis.ManagedAPIsClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	connectionGatewaysClient, err := connectiongateways.NewConnectionGatewaysClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ConnectionGateways client: %+v", err)
	}
	configureFunc(connectionGatewaysClient.Client)

	connectionsClient, err := connections.NewConnectionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Connections client: %+v", err)
	}
	configureFunc(connectionsClient.Client)

	customAPIsClient, err := customapis.NewCustomAPIsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CustomAPIs client: %+v", err)
	}
	configureFunc(customAPIsClient.Client)

	managedAPIsClient, err := managedapis.NewManagedAPIsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ManagedAPIs client: %+v", err)
	}
	configureFunc(managedAPIsClient.Client)

	return &Client{
		ConnectionGateways: connectionGatewaysClient,
		Connections:        connectionsClient,
		CustomAPIs:         customAPIsClient,
		ManagedAPIs:        managedAPIsClient,
	}, nil
}

package v2024_05_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/resource-manager/databricks/2024-05-01/accessconnector"
	"github.com/redtenant/go-azure-sdk/resource-manager/databricks/2024-05-01/delete"
	"github.com/redtenant/go-azure-sdk/resource-manager/databricks/2024-05-01/outboundnetworkdependenciesendpoints"
	"github.com/redtenant/go-azure-sdk/resource-manager/databricks/2024-05-01/privateendpointconnections"
	"github.com/redtenant/go-azure-sdk/resource-manager/databricks/2024-05-01/privatelinkresources"
	"github.com/redtenant/go-azure-sdk/resource-manager/databricks/2024-05-01/put"
	"github.com/redtenant/go-azure-sdk/resource-manager/databricks/2024-05-01/vnetpeering"
	"github.com/redtenant/go-azure-sdk/resource-manager/databricks/2024-05-01/workspaces"
	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

type Client struct {
	AccessConnector                      *accessconnector.AccessConnectorClient
	DELETE                               *delete.DELETEClient
	OutboundNetworkDependenciesEndpoints *outboundnetworkdependenciesendpoints.OutboundNetworkDependenciesEndpointsClient
	PUT                                  *put.PUTClient
	PrivateEndpointConnections           *privateendpointconnections.PrivateEndpointConnectionsClient
	PrivateLinkResources                 *privatelinkresources.PrivateLinkResourcesClient
	VNetPeering                          *vnetpeering.VNetPeeringClient
	Workspaces                           *workspaces.WorkspacesClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	accessConnectorClient, err := accessconnector.NewAccessConnectorClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AccessConnector client: %+v", err)
	}
	configureFunc(accessConnectorClient.Client)

	dELETEClient, err := delete.NewDELETEClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DELETE client: %+v", err)
	}
	configureFunc(dELETEClient.Client)

	outboundNetworkDependenciesEndpointsClient, err := outboundnetworkdependenciesendpoints.NewOutboundNetworkDependenciesEndpointsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building OutboundNetworkDependenciesEndpoints client: %+v", err)
	}
	configureFunc(outboundNetworkDependenciesEndpointsClient.Client)

	pUTClient, err := put.NewPUTClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PUT client: %+v", err)
	}
	configureFunc(pUTClient.Client)

	privateEndpointConnectionsClient, err := privateendpointconnections.NewPrivateEndpointConnectionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PrivateEndpointConnections client: %+v", err)
	}
	configureFunc(privateEndpointConnectionsClient.Client)

	privateLinkResourcesClient, err := privatelinkresources.NewPrivateLinkResourcesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PrivateLinkResources client: %+v", err)
	}
	configureFunc(privateLinkResourcesClient.Client)

	vNetPeeringClient, err := vnetpeering.NewVNetPeeringClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VNetPeering client: %+v", err)
	}
	configureFunc(vNetPeeringClient.Client)

	workspacesClient, err := workspaces.NewWorkspacesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Workspaces client: %+v", err)
	}
	configureFunc(workspacesClient.Client)

	return &Client{
		AccessConnector:                      accessConnectorClient,
		DELETE:                               dELETEClient,
		OutboundNetworkDependenciesEndpoints: outboundNetworkDependenciesEndpointsClient,
		PUT:                                  pUTClient,
		PrivateEndpointConnections:           privateEndpointConnectionsClient,
		PrivateLinkResources:                 privateLinkResourcesClient,
		VNetPeering:                          vNetPeeringClient,
		Workspaces:                           workspacesClient,
	}, nil
}

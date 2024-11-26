package v2021_07_01_preview

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/resource-manager/insights/2021-07-01-preview/privateendpointconnections"
	"github.com/redtenant/go-azure-sdk/resource-manager/insights/2021-07-01-preview/privatelinkresources"
	"github.com/redtenant/go-azure-sdk/resource-manager/insights/2021-07-01-preview/privatelinkscopedresources"
	"github.com/redtenant/go-azure-sdk/resource-manager/insights/2021-07-01-preview/privatelinkscopesapis"
	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

type Client struct {
	PrivateEndpointConnections *privateendpointconnections.PrivateEndpointConnectionsClient
	PrivateLinkResources       *privatelinkresources.PrivateLinkResourcesClient
	PrivateLinkScopedResources *privatelinkscopedresources.PrivateLinkScopedResourcesClient
	PrivateLinkScopesAPIs      *privatelinkscopesapis.PrivateLinkScopesAPIsClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
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

	privateLinkScopedResourcesClient, err := privatelinkscopedresources.NewPrivateLinkScopedResourcesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PrivateLinkScopedResources client: %+v", err)
	}
	configureFunc(privateLinkScopedResourcesClient.Client)

	privateLinkScopesAPIsClient, err := privatelinkscopesapis.NewPrivateLinkScopesAPIsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PrivateLinkScopesAPIs client: %+v", err)
	}
	configureFunc(privateLinkScopesAPIsClient.Client)

	return &Client{
		PrivateEndpointConnections: privateEndpointConnectionsClient,
		PrivateLinkResources:       privateLinkResourcesClient,
		PrivateLinkScopedResources: privateLinkScopedResourcesClient,
		PrivateLinkScopesAPIs:      privateLinkScopesAPIsClient,
	}, nil
}

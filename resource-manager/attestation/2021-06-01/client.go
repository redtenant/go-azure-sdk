package v2021_06_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/resource-manager/attestation/2021-06-01/attestationproviders"
	"github.com/redtenant/go-azure-sdk/resource-manager/attestation/2021-06-01/privateendpointconnections"
	"github.com/redtenant/go-azure-sdk/resource-manager/attestation/2021-06-01/privatelinkresources"
	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

type Client struct {
	AttestationProviders       *attestationproviders.AttestationProvidersClient
	PrivateEndpointConnections *privateendpointconnections.PrivateEndpointConnectionsClient
	PrivateLinkResources       *privatelinkresources.PrivateLinkResourcesClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	attestationProvidersClient, err := attestationproviders.NewAttestationProvidersClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AttestationProviders client: %+v", err)
	}
	configureFunc(attestationProvidersClient.Client)

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

	return &Client{
		AttestationProviders:       attestationProvidersClient,
		PrivateEndpointConnections: privateEndpointConnectionsClient,
		PrivateLinkResources:       privateLinkResourcesClient,
	}, nil
}

package v2021_07_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/resource-manager/purview/2021-07-01/account"
	"github.com/redtenant/go-azure-sdk/resource-manager/purview/2021-07-01/defaultaccount"
	"github.com/redtenant/go-azure-sdk/resource-manager/purview/2021-07-01/privateendpointconnection"
	"github.com/redtenant/go-azure-sdk/resource-manager/purview/2021-07-01/privatelinkresource"
	"github.com/redtenant/go-azure-sdk/resource-manager/purview/2021-07-01/provider"
	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

type Client struct {
	Account                   *account.AccountClient
	DefaultAccount            *defaultaccount.DefaultAccountClient
	PrivateEndpointConnection *privateendpointconnection.PrivateEndpointConnectionClient
	PrivateLinkResource       *privatelinkresource.PrivateLinkResourceClient
	Provider                  *provider.ProviderClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	accountClient, err := account.NewAccountClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Account client: %+v", err)
	}
	configureFunc(accountClient.Client)

	defaultAccountClient, err := defaultaccount.NewDefaultAccountClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DefaultAccount client: %+v", err)
	}
	configureFunc(defaultAccountClient.Client)

	privateEndpointConnectionClient, err := privateendpointconnection.NewPrivateEndpointConnectionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PrivateEndpointConnection client: %+v", err)
	}
	configureFunc(privateEndpointConnectionClient.Client)

	privateLinkResourceClient, err := privatelinkresource.NewPrivateLinkResourceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PrivateLinkResource client: %+v", err)
	}
	configureFunc(privateLinkResourceClient.Client)

	providerClient, err := provider.NewProviderClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Provider client: %+v", err)
	}
	configureFunc(providerClient.Client)

	return &Client{
		Account:                   accountClient,
		DefaultAccount:            defaultAccountClient,
		PrivateEndpointConnection: privateEndpointConnectionClient,
		PrivateLinkResource:       privateLinkResourceClient,
		Provider:                  providerClient,
	}, nil
}

package v2021_01_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/resource-manager/mixedreality/2021-01-01/key"
	"github.com/redtenant/go-azure-sdk/resource-manager/mixedreality/2021-01-01/proxy"
	"github.com/redtenant/go-azure-sdk/resource-manager/mixedreality/2021-01-01/resource"
	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

type Client struct {
	Key      *key.KeyClient
	Proxy    *proxy.ProxyClient
	Resource *resource.ResourceClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	keyClient, err := key.NewKeyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Key client: %+v", err)
	}
	configureFunc(keyClient.Client)

	proxyClient, err := proxy.NewProxyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Proxy client: %+v", err)
	}
	configureFunc(proxyClient.Client)

	resourceClient, err := resource.NewResourceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Resource client: %+v", err)
	}
	configureFunc(resourceClient.Client)

	return &Client{
		Key:      keyClient,
		Proxy:    proxyClient,
		Resource: resourceClient,
	}, nil
}

package virtualnetworkrules

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualNetworkRulesClient struct {
	Client *resourcemanager.Client
}

func NewVirtualNetworkRulesClientWithBaseURI(sdkApi sdkEnv.Api) (*VirtualNetworkRulesClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "virtualnetworkrules", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating VirtualNetworkRulesClient: %+v", err)
	}

	return &VirtualNetworkRulesClient{
		Client: client,
	}, nil
}

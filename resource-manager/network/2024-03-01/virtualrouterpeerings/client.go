package virtualrouterpeerings

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualRouterPeeringsClient struct {
	Client *resourcemanager.Client
}

func NewVirtualRouterPeeringsClientWithBaseURI(sdkApi sdkEnv.Api) (*VirtualRouterPeeringsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "virtualrouterpeerings", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating VirtualRouterPeeringsClient: %+v", err)
	}

	return &VirtualRouterPeeringsClient{
		Client: client,
	}, nil
}

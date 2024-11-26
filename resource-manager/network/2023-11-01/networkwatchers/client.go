package networkwatchers

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkWatchersClient struct {
	Client *resourcemanager.Client
}

func NewNetworkWatchersClientWithBaseURI(sdkApi sdkEnv.Api) (*NetworkWatchersClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "networkwatchers", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating NetworkWatchersClient: %+v", err)
	}

	return &NetworkWatchersClient{
		Client: client,
	}, nil
}

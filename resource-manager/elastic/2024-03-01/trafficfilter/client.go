package trafficfilter

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TrafficFilterClient struct {
	Client *resourcemanager.Client
}

func NewTrafficFilterClientWithBaseURI(sdkApi sdkEnv.Api) (*TrafficFilterClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "trafficfilter", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TrafficFilterClient: %+v", err)
	}

	return &TrafficFilterClient{
		Client: client,
	}, nil
}

package vnetpeering

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VNetPeeringClient struct {
	Client *resourcemanager.Client
}

func NewVNetPeeringClientWithBaseURI(sdkApi sdkEnv.Api) (*VNetPeeringClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "vnetpeering", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating VNetPeeringClient: %+v", err)
	}

	return &VNetPeeringClient{
		Client: client,
	}, nil
}

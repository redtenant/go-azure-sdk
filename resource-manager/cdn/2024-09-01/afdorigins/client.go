package afdorigins

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AFDOriginsClient struct {
	Client *resourcemanager.Client
}

func NewAFDOriginsClientWithBaseURI(sdkApi sdkEnv.Api) (*AFDOriginsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "afdorigins", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AFDOriginsClient: %+v", err)
	}

	return &AFDOriginsClient{
		Client: client,
	}, nil
}

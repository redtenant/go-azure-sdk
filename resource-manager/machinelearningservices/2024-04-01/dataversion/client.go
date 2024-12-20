package dataversion

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataVersionClient struct {
	Client *resourcemanager.Client
}

func NewDataVersionClientWithBaseURI(sdkApi sdkEnv.Api) (*DataVersionClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "dataversion", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DataVersionClient: %+v", err)
	}

	return &DataVersionClient{
		Client: client,
	}, nil
}

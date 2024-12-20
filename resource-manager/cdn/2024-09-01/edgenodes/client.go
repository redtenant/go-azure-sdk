package edgenodes

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EdgenodesClient struct {
	Client *resourcemanager.Client
}

func NewEdgenodesClientWithBaseURI(sdkApi sdkEnv.Api) (*EdgenodesClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "edgenodes", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EdgenodesClient: %+v", err)
	}

	return &EdgenodesClient{
		Client: client,
	}, nil
}

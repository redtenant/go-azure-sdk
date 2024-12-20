package graphapicompute

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GraphAPIComputeClient struct {
	Client *resourcemanager.Client
}

func NewGraphAPIComputeClientWithBaseURI(sdkApi sdkEnv.Api) (*GraphAPIComputeClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "graphapicompute", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GraphAPIComputeClient: %+v", err)
	}

	return &GraphAPIComputeClient{
		Client: client,
	}, nil
}

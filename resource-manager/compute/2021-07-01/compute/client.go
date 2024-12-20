package compute

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ComputeClient struct {
	Client *resourcemanager.Client
}

func NewComputeClientWithBaseURI(sdkApi sdkEnv.Api) (*ComputeClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "compute", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ComputeClient: %+v", err)
	}

	return &ComputeClient{
		Client: client,
	}, nil
}

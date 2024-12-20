package hybridrunbookworker

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HybridRunbookWorkerClient struct {
	Client *resourcemanager.Client
}

func NewHybridRunbookWorkerClientWithBaseURI(sdkApi sdkEnv.Api) (*HybridRunbookWorkerClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "hybridrunbookworker", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating HybridRunbookWorkerClient: %+v", err)
	}

	return &HybridRunbookWorkerClient{
		Client: client,
	}, nil
}

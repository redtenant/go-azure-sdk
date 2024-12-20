package globalschedules

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GlobalSchedulesClient struct {
	Client *resourcemanager.Client
}

func NewGlobalSchedulesClientWithBaseURI(sdkApi sdkEnv.Api) (*GlobalSchedulesClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "globalschedules", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GlobalSchedulesClient: %+v", err)
	}

	return &GlobalSchedulesClient{
		Client: client,
	}, nil
}

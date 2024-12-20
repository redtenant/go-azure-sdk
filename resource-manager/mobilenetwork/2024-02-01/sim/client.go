package sim

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SIMClient struct {
	Client *resourcemanager.Client
}

func NewSIMClientWithBaseURI(sdkApi sdkEnv.Api) (*SIMClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "sim", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SIMClient: %+v", err)
	}

	return &SIMClient{
		Client: client,
	}, nil
}

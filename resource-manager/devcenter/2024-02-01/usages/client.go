package usages

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UsagesClient struct {
	Client *resourcemanager.Client
}

func NewUsagesClientWithBaseURI(sdkApi sdkEnv.Api) (*UsagesClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "usages", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UsagesClient: %+v", err)
	}

	return &UsagesClient{
		Client: client,
	}, nil
}

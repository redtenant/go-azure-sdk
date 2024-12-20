package ipallocations

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IPAllocationsClient struct {
	Client *resourcemanager.Client
}

func NewIPAllocationsClientWithBaseURI(sdkApi sdkEnv.Api) (*IPAllocationsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "ipallocations", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating IPAllocationsClient: %+v", err)
	}

	return &IPAllocationsClient{
		Client: client,
	}, nil
}

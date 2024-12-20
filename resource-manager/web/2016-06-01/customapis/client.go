package customapis

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CustomAPIsClient struct {
	Client *resourcemanager.Client
}

func NewCustomAPIsClientWithBaseURI(sdkApi sdkEnv.Api) (*CustomAPIsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "customapis", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CustomAPIsClient: %+v", err)
	}

	return &CustomAPIsClient{
		Client: client,
	}, nil
}

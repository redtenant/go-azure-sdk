package connection

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConnectionClient struct {
	Client *resourcemanager.Client
}

func NewConnectionClientWithBaseURI(sdkApi sdkEnv.Api) (*ConnectionClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "connection", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ConnectionClient: %+v", err)
	}

	return &ConnectionClient{
		Client: client,
	}, nil
}

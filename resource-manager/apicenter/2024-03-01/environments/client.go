package environments

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EnvironmentsClient struct {
	Client *resourcemanager.Client
}

func NewEnvironmentsClientWithBaseURI(sdkApi sdkEnv.Api) (*EnvironmentsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "environments", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EnvironmentsClient: %+v", err)
	}

	return &EnvironmentsClient{
		Client: client,
	}, nil
}

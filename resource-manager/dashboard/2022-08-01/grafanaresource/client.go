package grafanaresource

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GrafanaResourceClient struct {
	Client *resourcemanager.Client
}

func NewGrafanaResourceClientWithBaseURI(sdkApi sdkEnv.Api) (*GrafanaResourceClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "grafanaresource", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GrafanaResourceClient: %+v", err)
	}

	return &GrafanaResourceClient{
		Client: client,
	}, nil
}

package monitoredresources

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MonitoredResourcesClient struct {
	Client *resourcemanager.Client
}

func NewMonitoredResourcesClientWithBaseURI(sdkApi sdkEnv.Api) (*MonitoredResourcesClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "monitoredresources", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MonitoredResourcesClient: %+v", err)
	}

	return &MonitoredResourcesClient{
		Client: client,
	}, nil
}

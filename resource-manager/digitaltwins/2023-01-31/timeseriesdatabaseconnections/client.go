package timeseriesdatabaseconnections

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TimeSeriesDatabaseConnectionsClient struct {
	Client *resourcemanager.Client
}

func NewTimeSeriesDatabaseConnectionsClientWithBaseURI(sdkApi sdkEnv.Api) (*TimeSeriesDatabaseConnectionsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "timeseriesdatabaseconnections", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TimeSeriesDatabaseConnectionsClient: %+v", err)
	}

	return &TimeSeriesDatabaseConnectionsClient{
		Client: client,
	}, nil
}

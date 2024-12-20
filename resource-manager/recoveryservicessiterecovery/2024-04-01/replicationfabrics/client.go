package replicationfabrics

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReplicationFabricsClient struct {
	Client *resourcemanager.Client
}

func NewReplicationFabricsClientWithBaseURI(sdkApi sdkEnv.Api) (*ReplicationFabricsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "replicationfabrics", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ReplicationFabricsClient: %+v", err)
	}

	return &ReplicationFabricsClient{
		Client: client,
	}, nil
}

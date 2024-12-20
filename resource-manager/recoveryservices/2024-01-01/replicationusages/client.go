package replicationusages

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReplicationUsagesClient struct {
	Client *resourcemanager.Client
}

func NewReplicationUsagesClientWithBaseURI(sdkApi sdkEnv.Api) (*ReplicationUsagesClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "replicationusages", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ReplicationUsagesClient: %+v", err)
	}

	return &ReplicationUsagesClient{
		Client: client,
	}, nil
}

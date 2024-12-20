package replicationprotectioncontainers

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReplicationProtectionContainersClient struct {
	Client *resourcemanager.Client
}

func NewReplicationProtectionContainersClientWithBaseURI(sdkApi sdkEnv.Api) (*ReplicationProtectionContainersClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "replicationprotectioncontainers", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ReplicationProtectionContainersClient: %+v", err)
	}

	return &ReplicationProtectionContainersClient{
		Client: client,
	}, nil
}

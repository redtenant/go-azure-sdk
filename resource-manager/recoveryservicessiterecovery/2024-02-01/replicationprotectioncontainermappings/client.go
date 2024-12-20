package replicationprotectioncontainermappings

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReplicationProtectionContainerMappingsClient struct {
	Client *resourcemanager.Client
}

func NewReplicationProtectionContainerMappingsClientWithBaseURI(sdkApi sdkEnv.Api) (*ReplicationProtectionContainerMappingsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "replicationprotectioncontainermappings", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ReplicationProtectionContainerMappingsClient: %+v", err)
	}

	return &ReplicationProtectionContainerMappingsClient{
		Client: client,
	}, nil
}

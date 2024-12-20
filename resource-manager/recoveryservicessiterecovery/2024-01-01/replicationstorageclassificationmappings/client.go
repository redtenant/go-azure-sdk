package replicationstorageclassificationmappings

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReplicationStorageClassificationMappingsClient struct {
	Client *resourcemanager.Client
}

func NewReplicationStorageClassificationMappingsClientWithBaseURI(sdkApi sdkEnv.Api) (*ReplicationStorageClassificationMappingsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "replicationstorageclassificationmappings", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ReplicationStorageClassificationMappingsClient: %+v", err)
	}

	return &ReplicationStorageClassificationMappingsClient{
		Client: client,
	}, nil
}

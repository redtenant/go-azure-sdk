package replicationrecoveryplans

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReplicationRecoveryPlansClient struct {
	Client *resourcemanager.Client
}

func NewReplicationRecoveryPlansClientWithBaseURI(sdkApi sdkEnv.Api) (*ReplicationRecoveryPlansClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "replicationrecoveryplans", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ReplicationRecoveryPlansClient: %+v", err)
	}

	return &ReplicationRecoveryPlansClient{
		Client: client,
	}, nil
}

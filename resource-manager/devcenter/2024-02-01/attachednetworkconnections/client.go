package attachednetworkconnections

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AttachedNetworkConnectionsClient struct {
	Client *resourcemanager.Client
}

func NewAttachedNetworkConnectionsClientWithBaseURI(sdkApi sdkEnv.Api) (*AttachedNetworkConnectionsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "attachednetworkconnections", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AttachedNetworkConnectionsClient: %+v", err)
	}

	return &AttachedNetworkConnectionsClient{
		Client: client,
	}, nil
}

package restorepoints

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RestorePointsClient struct {
	Client *resourcemanager.Client
}

func NewRestorePointsClientWithBaseURI(sdkApi sdkEnv.Api) (*RestorePointsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "restorepoints", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RestorePointsClient: %+v", err)
	}

	return &RestorePointsClient{
		Client: client,
	}, nil
}

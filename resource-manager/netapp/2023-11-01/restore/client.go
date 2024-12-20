package restore

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RestoreClient struct {
	Client *resourcemanager.Client
}

func NewRestoreClientWithBaseURI(sdkApi sdkEnv.Api) (*RestoreClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "restore", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RestoreClient: %+v", err)
	}

	return &RestoreClient{
		Client: client,
	}, nil
}

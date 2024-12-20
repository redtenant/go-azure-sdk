package updateruns

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateRunsClient struct {
	Client *resourcemanager.Client
}

func NewUpdateRunsClientWithBaseURI(sdkApi sdkEnv.Api) (*UpdateRunsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "updateruns", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UpdateRunsClient: %+v", err)
	}

	return &UpdateRunsClient{
		Client: client,
	}, nil
}

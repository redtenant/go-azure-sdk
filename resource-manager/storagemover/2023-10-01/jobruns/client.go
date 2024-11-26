package jobruns

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JobRunsClient struct {
	Client *resourcemanager.Client
}

func NewJobRunsClientWithBaseURI(sdkApi sdkEnv.Api) (*JobRunsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "jobruns", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating JobRunsClient: %+v", err)
	}

	return &JobRunsClient{
		Client: client,
	}, nil
}

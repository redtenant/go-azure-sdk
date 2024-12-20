package pipelineruns

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PipelineRunsClient struct {
	Client *resourcemanager.Client
}

func NewPipelineRunsClientWithBaseURI(sdkApi sdkEnv.Api) (*PipelineRunsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "pipelineruns", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PipelineRunsClient: %+v", err)
	}

	return &PipelineRunsClient{
		Client: client,
	}, nil
}

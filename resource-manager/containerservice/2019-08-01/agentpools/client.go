package agentpools

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AgentPoolsClient struct {
	Client *resourcemanager.Client
}

func NewAgentPoolsClientWithBaseURI(sdkApi sdkEnv.Api) (*AgentPoolsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "agentpools", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AgentPoolsClient: %+v", err)
	}

	return &AgentPoolsClient{
		Client: client,
	}, nil
}

package workflowresource

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkflowResourceClient struct {
	Client *resourcemanager.Client
}

func NewWorkflowResourceClientWithBaseURI(sdkApi sdkEnv.Api) (*WorkflowResourceClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "workflowresource", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating WorkflowResourceClient: %+v", err)
	}

	return &WorkflowResourceClient{
		Client: client,
	}, nil
}

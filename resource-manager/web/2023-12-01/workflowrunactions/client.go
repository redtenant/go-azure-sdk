package workflowrunactions

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkflowRunActionsClient struct {
	Client *resourcemanager.Client
}

func NewWorkflowRunActionsClientWithBaseURI(sdkApi sdkEnv.Api) (*WorkflowRunActionsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "workflowrunactions", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating WorkflowRunActionsClient: %+v", err)
	}

	return &WorkflowRunActionsClient{
		Client: client,
	}, nil
}

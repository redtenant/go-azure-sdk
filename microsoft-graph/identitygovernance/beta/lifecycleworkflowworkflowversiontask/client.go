package lifecycleworkflowworkflowversiontask

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LifecycleWorkflowWorkflowVersionTaskClient struct {
	Client *msgraph.Client
}

func NewLifecycleWorkflowWorkflowVersionTaskClientWithBaseURI(sdkApi sdkEnv.Api) (*LifecycleWorkflowWorkflowVersionTaskClient, error) {
	client, err := msgraph.NewClient(sdkApi, "lifecycleworkflowworkflowversiontask", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LifecycleWorkflowWorkflowVersionTaskClient: %+v", err)
	}

	return &LifecycleWorkflowWorkflowVersionTaskClient{
		Client: client,
	}, nil
}

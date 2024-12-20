package lifecycleworkflowworkflowrunuserprocessingresulttaskprocessingresulttask

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LifecycleWorkflowWorkflowRunUserProcessingResultTaskProcessingResultTaskClient struct {
	Client *msgraph.Client
}

func NewLifecycleWorkflowWorkflowRunUserProcessingResultTaskProcessingResultTaskClientWithBaseURI(sdkApi sdkEnv.Api) (*LifecycleWorkflowWorkflowRunUserProcessingResultTaskProcessingResultTaskClient, error) {
	client, err := msgraph.NewClient(sdkApi, "lifecycleworkflowworkflowrunuserprocessingresulttaskprocessingresulttask", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LifecycleWorkflowWorkflowRunUserProcessingResultTaskProcessingResultTaskClient: %+v", err)
	}

	return &LifecycleWorkflowWorkflowRunUserProcessingResultTaskProcessingResultTaskClient{
		Client: client,
	}, nil
}

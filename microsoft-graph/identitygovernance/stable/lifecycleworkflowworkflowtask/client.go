package lifecycleworkflowworkflowtask

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LifecycleWorkflowWorkflowTaskClient struct {
	Client *msgraph.Client
}

func NewLifecycleWorkflowWorkflowTaskClientWithBaseURI(sdkApi sdkEnv.Api) (*LifecycleWorkflowWorkflowTaskClient, error) {
	client, err := msgraph.NewClient(sdkApi, "lifecycleworkflowworkflowtask", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LifecycleWorkflowWorkflowTaskClient: %+v", err)
	}

	return &LifecycleWorkflowWorkflowTaskClient{
		Client: client,
	}, nil
}

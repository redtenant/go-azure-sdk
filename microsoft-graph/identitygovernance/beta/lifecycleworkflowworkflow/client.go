package lifecycleworkflowworkflow

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LifecycleWorkflowWorkflowClient struct {
	Client *msgraph.Client
}

func NewLifecycleWorkflowWorkflowClientWithBaseURI(sdkApi sdkEnv.Api) (*LifecycleWorkflowWorkflowClient, error) {
	client, err := msgraph.NewClient(sdkApi, "lifecycleworkflowworkflow", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LifecycleWorkflowWorkflowClient: %+v", err)
	}

	return &LifecycleWorkflowWorkflowClient{
		Client: client,
	}, nil
}

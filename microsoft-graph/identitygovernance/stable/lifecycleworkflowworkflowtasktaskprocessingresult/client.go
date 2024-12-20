package lifecycleworkflowworkflowtasktaskprocessingresult

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LifecycleWorkflowWorkflowTaskTaskProcessingResultClient struct {
	Client *msgraph.Client
}

func NewLifecycleWorkflowWorkflowTaskTaskProcessingResultClientWithBaseURI(sdkApi sdkEnv.Api) (*LifecycleWorkflowWorkflowTaskTaskProcessingResultClient, error) {
	client, err := msgraph.NewClient(sdkApi, "lifecycleworkflowworkflowtasktaskprocessingresult", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LifecycleWorkflowWorkflowTaskTaskProcessingResultClient: %+v", err)
	}

	return &LifecycleWorkflowWorkflowTaskTaskProcessingResultClient{
		Client: client,
	}, nil
}

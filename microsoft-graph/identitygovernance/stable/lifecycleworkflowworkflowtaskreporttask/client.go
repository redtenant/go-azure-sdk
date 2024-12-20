package lifecycleworkflowworkflowtaskreporttask

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LifecycleWorkflowWorkflowTaskReportTaskClient struct {
	Client *msgraph.Client
}

func NewLifecycleWorkflowWorkflowTaskReportTaskClientWithBaseURI(sdkApi sdkEnv.Api) (*LifecycleWorkflowWorkflowTaskReportTaskClient, error) {
	client, err := msgraph.NewClient(sdkApi, "lifecycleworkflowworkflowtaskreporttask", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LifecycleWorkflowWorkflowTaskReportTaskClient: %+v", err)
	}

	return &LifecycleWorkflowWorkflowTaskReportTaskClient{
		Client: client,
	}, nil
}

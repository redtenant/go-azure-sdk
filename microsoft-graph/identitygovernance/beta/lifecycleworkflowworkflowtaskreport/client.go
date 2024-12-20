package lifecycleworkflowworkflowtaskreport

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LifecycleWorkflowWorkflowTaskReportClient struct {
	Client *msgraph.Client
}

func NewLifecycleWorkflowWorkflowTaskReportClientWithBaseURI(sdkApi sdkEnv.Api) (*LifecycleWorkflowWorkflowTaskReportClient, error) {
	client, err := msgraph.NewClient(sdkApi, "lifecycleworkflowworkflowtaskreport", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LifecycleWorkflowWorkflowTaskReportClient: %+v", err)
	}

	return &LifecycleWorkflowWorkflowTaskReportClient{
		Client: client,
	}, nil
}

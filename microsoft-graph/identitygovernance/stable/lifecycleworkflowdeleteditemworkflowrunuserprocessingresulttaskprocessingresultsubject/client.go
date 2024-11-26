package lifecycleworkflowdeleteditemworkflowrunuserprocessingresulttaskprocessingresultsubject

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultTaskProcessingResultSubjectClient struct {
	Client *msgraph.Client
}

func NewLifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultTaskProcessingResultSubjectClientWithBaseURI(sdkApi sdkEnv.Api) (*LifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultTaskProcessingResultSubjectClient, error) {
	client, err := msgraph.NewClient(sdkApi, "lifecycleworkflowdeleteditemworkflowrunuserprocessingresulttaskprocessingresultsubject", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultTaskProcessingResultSubjectClient: %+v", err)
	}

	return &LifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultTaskProcessingResultSubjectClient{
		Client: client,
	}, nil
}

package lifecycleworkflowworkflowrunuserprocessingresulttaskprocessingresultsubjectserviceprovisioningerror

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LifecycleWorkflowWorkflowRunUserProcessingResultTaskProcessingResultSubjectServiceProvisioningErrorClient struct {
	Client *msgraph.Client
}

func NewLifecycleWorkflowWorkflowRunUserProcessingResultTaskProcessingResultSubjectServiceProvisioningErrorClientWithBaseURI(sdkApi sdkEnv.Api) (*LifecycleWorkflowWorkflowRunUserProcessingResultTaskProcessingResultSubjectServiceProvisioningErrorClient, error) {
	client, err := msgraph.NewClient(sdkApi, "lifecycleworkflowworkflowrunuserprocessingresulttaskprocessingresultsubjectserviceprovisioningerror", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LifecycleWorkflowWorkflowRunUserProcessingResultTaskProcessingResultSubjectServiceProvisioningErrorClient: %+v", err)
	}

	return &LifecycleWorkflowWorkflowRunUserProcessingResultTaskProcessingResultSubjectServiceProvisioningErrorClient{
		Client: client,
	}, nil
}

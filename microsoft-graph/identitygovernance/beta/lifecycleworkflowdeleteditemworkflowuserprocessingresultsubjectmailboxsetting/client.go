package lifecycleworkflowdeleteditemworkflowuserprocessingresultsubjectmailboxsetting

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LifecycleWorkflowDeletedItemWorkflowUserProcessingResultSubjectMailboxSettingClient struct {
	Client *msgraph.Client
}

func NewLifecycleWorkflowDeletedItemWorkflowUserProcessingResultSubjectMailboxSettingClientWithBaseURI(sdkApi sdkEnv.Api) (*LifecycleWorkflowDeletedItemWorkflowUserProcessingResultSubjectMailboxSettingClient, error) {
	client, err := msgraph.NewClient(sdkApi, "lifecycleworkflowdeleteditemworkflowuserprocessingresultsubjectmailboxsetting", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LifecycleWorkflowDeletedItemWorkflowUserProcessingResultSubjectMailboxSettingClient: %+v", err)
	}

	return &LifecycleWorkflowDeletedItemWorkflowUserProcessingResultSubjectMailboxSettingClient{
		Client: client,
	}, nil
}

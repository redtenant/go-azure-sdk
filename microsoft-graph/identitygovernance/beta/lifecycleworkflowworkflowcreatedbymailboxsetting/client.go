package lifecycleworkflowworkflowcreatedbymailboxsetting

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LifecycleWorkflowWorkflowCreatedByMailboxSettingClient struct {
	Client *msgraph.Client
}

func NewLifecycleWorkflowWorkflowCreatedByMailboxSettingClientWithBaseURI(sdkApi sdkEnv.Api) (*LifecycleWorkflowWorkflowCreatedByMailboxSettingClient, error) {
	client, err := msgraph.NewClient(sdkApi, "lifecycleworkflowworkflowcreatedbymailboxsetting", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LifecycleWorkflowWorkflowCreatedByMailboxSettingClient: %+v", err)
	}

	return &LifecycleWorkflowWorkflowCreatedByMailboxSettingClient{
		Client: client,
	}, nil
}

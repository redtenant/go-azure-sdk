package lifecycleworkflowdeleteditemworkflowversionlastmodifiedbymailboxsetting

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LifecycleWorkflowDeletedItemWorkflowVersionLastModifiedByMailboxSettingClient struct {
	Client *msgraph.Client
}

func NewLifecycleWorkflowDeletedItemWorkflowVersionLastModifiedByMailboxSettingClientWithBaseURI(sdkApi sdkEnv.Api) (*LifecycleWorkflowDeletedItemWorkflowVersionLastModifiedByMailboxSettingClient, error) {
	client, err := msgraph.NewClient(sdkApi, "lifecycleworkflowdeleteditemworkflowversionlastmodifiedbymailboxsetting", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LifecycleWorkflowDeletedItemWorkflowVersionLastModifiedByMailboxSettingClient: %+v", err)
	}

	return &LifecycleWorkflowDeletedItemWorkflowVersionLastModifiedByMailboxSettingClient{
		Client: client,
	}, nil
}

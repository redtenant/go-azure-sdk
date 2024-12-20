package lifecycleworkflowworkflowversionlastmodifiedby

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LifecycleWorkflowWorkflowVersionLastModifiedByClient struct {
	Client *msgraph.Client
}

func NewLifecycleWorkflowWorkflowVersionLastModifiedByClientWithBaseURI(sdkApi sdkEnv.Api) (*LifecycleWorkflowWorkflowVersionLastModifiedByClient, error) {
	client, err := msgraph.NewClient(sdkApi, "lifecycleworkflowworkflowversionlastmodifiedby", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LifecycleWorkflowWorkflowVersionLastModifiedByClient: %+v", err)
	}

	return &LifecycleWorkflowWorkflowVersionLastModifiedByClient{
		Client: client,
	}, nil
}

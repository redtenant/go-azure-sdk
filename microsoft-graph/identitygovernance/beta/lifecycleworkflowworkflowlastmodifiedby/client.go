package lifecycleworkflowworkflowlastmodifiedby

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LifecycleWorkflowWorkflowLastModifiedByClient struct {
	Client *msgraph.Client
}

func NewLifecycleWorkflowWorkflowLastModifiedByClientWithBaseURI(sdkApi sdkEnv.Api) (*LifecycleWorkflowWorkflowLastModifiedByClient, error) {
	client, err := msgraph.NewClient(sdkApi, "lifecycleworkflowworkflowlastmodifiedby", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LifecycleWorkflowWorkflowLastModifiedByClient: %+v", err)
	}

	return &LifecycleWorkflowWorkflowLastModifiedByClient{
		Client: client,
	}, nil
}

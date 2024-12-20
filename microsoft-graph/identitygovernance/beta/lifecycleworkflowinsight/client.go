package lifecycleworkflowinsight

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LifecycleWorkflowInsightClient struct {
	Client *msgraph.Client
}

func NewLifecycleWorkflowInsightClientWithBaseURI(sdkApi sdkEnv.Api) (*LifecycleWorkflowInsightClient, error) {
	client, err := msgraph.NewClient(sdkApi, "lifecycleworkflowinsight", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LifecycleWorkflowInsightClient: %+v", err)
	}

	return &LifecycleWorkflowInsightClient{
		Client: client,
	}, nil
}

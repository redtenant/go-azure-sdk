package operationprogress

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OperationProgressClient struct {
	Client *resourcemanager.Client
}

func NewOperationProgressClientWithBaseURI(sdkApi sdkEnv.Api) (*OperationProgressClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "operationprogress", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating OperationProgressClient: %+v", err)
	}

	return &OperationProgressClient{
		Client: client,
	}, nil
}

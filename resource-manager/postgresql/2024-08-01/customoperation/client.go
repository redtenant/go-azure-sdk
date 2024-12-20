package customoperation

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CustomOperationClient struct {
	Client *resourcemanager.Client
}

func NewCustomOperationClientWithBaseURI(sdkApi sdkEnv.Api) (*CustomOperationClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "customoperation", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CustomOperationClient: %+v", err)
	}

	return &CustomOperationClient{
		Client: client,
	}, nil
}

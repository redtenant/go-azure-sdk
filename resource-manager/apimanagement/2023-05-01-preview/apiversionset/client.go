package apiversionset

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApiVersionSetClient struct {
	Client *resourcemanager.Client
}

func NewApiVersionSetClientWithBaseURI(sdkApi sdkEnv.Api) (*ApiVersionSetClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "apiversionset", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ApiVersionSetClient: %+v", err)
	}

	return &ApiVersionSetClient{
		Client: client,
	}, nil
}

package codecontainer

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CodeContainerClient struct {
	Client *resourcemanager.Client
}

func NewCodeContainerClientWithBaseURI(sdkApi sdkEnv.Api) (*CodeContainerClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "codecontainer", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CodeContainerClient: %+v", err)
	}

	return &CodeContainerClient{
		Client: client,
	}, nil
}

package createresource

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateResourceClient struct {
	Client *resourcemanager.Client
}

func NewCreateResourceClientWithBaseURI(sdkApi sdkEnv.Api) (*CreateResourceClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "createresource", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CreateResourceClient: %+v", err)
	}

	return &CreateResourceClient{
		Client: client,
	}, nil
}

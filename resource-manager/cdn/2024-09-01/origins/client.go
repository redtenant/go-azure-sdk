package origins

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OriginsClient struct {
	Client *resourcemanager.Client
}

func NewOriginsClientWithBaseURI(sdkApi sdkEnv.Api) (*OriginsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "origins", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating OriginsClient: %+v", err)
	}

	return &OriginsClient{
		Client: client,
	}, nil
}

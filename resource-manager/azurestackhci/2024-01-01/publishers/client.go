package publishers

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PublishersClient struct {
	Client *resourcemanager.Client
}

func NewPublishersClientWithBaseURI(sdkApi sdkEnv.Api) (*PublishersClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "publishers", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PublishersClient: %+v", err)
	}

	return &PublishersClient{
		Client: client,
	}, nil
}

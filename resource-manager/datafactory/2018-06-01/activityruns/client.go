package activityruns

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ActivityrunsClient struct {
	Client *resourcemanager.Client
}

func NewActivityrunsClientWithBaseURI(sdkApi sdkEnv.Api) (*ActivityrunsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "activityruns", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ActivityrunsClient: %+v", err)
	}

	return &ActivityrunsClient{
		Client: client,
	}, nil
}

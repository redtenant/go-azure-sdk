package ipgroups

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IPGroupsClient struct {
	Client *resourcemanager.Client
}

func NewIPGroupsClientWithBaseURI(sdkApi sdkEnv.Api) (*IPGroupsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "ipgroups", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating IPGroupsClient: %+v", err)
	}

	return &IPGroupsClient{
		Client: client,
	}, nil
}

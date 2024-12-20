package staticmembers

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StaticMembersClient struct {
	Client *resourcemanager.Client
}

func NewStaticMembersClientWithBaseURI(sdkApi sdkEnv.Api) (*StaticMembersClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "staticmembers", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating StaticMembersClient: %+v", err)
	}

	return &StaticMembersClient{
		Client: client,
	}, nil
}

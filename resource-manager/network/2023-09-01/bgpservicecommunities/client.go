package bgpservicecommunities

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BgpServiceCommunitiesClient struct {
	Client *resourcemanager.Client
}

func NewBgpServiceCommunitiesClientWithBaseURI(sdkApi sdkEnv.Api) (*BgpServiceCommunitiesClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "bgpservicecommunities", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating BgpServiceCommunitiesClient: %+v", err)
	}

	return &BgpServiceCommunitiesClient{
		Client: client,
	}, nil
}

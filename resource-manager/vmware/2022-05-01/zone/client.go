package zone

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ZoneClient struct {
	Client *resourcemanager.Client
}

func NewZoneClientWithBaseURI(sdkApi sdkEnv.Api) (*ZoneClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "zone", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ZoneClient: %+v", err)
	}

	return &ZoneClient{
		Client: client,
	}, nil
}

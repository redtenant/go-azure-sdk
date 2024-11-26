package componentproactivedetectionapis

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ComponentProactiveDetectionAPIsClient struct {
	Client *resourcemanager.Client
}

func NewComponentProactiveDetectionAPIsClientWithBaseURI(sdkApi sdkEnv.Api) (*ComponentProactiveDetectionAPIsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "componentproactivedetectionapis", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ComponentProactiveDetectionAPIsClient: %+v", err)
	}

	return &ComponentProactiveDetectionAPIsClient{
		Client: client,
	}, nil
}

package applicationgatewaywafdynamicmanifests

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApplicationGatewayWafDynamicManifestsClient struct {
	Client *resourcemanager.Client
}

func NewApplicationGatewayWafDynamicManifestsClientWithBaseURI(sdkApi sdkEnv.Api) (*ApplicationGatewayWafDynamicManifestsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "applicationgatewaywafdynamicmanifests", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ApplicationGatewayWafDynamicManifestsClient: %+v", err)
	}

	return &ApplicationGatewayWafDynamicManifestsClient{
		Client: client,
	}, nil
}

package expressroutecircuitauthorizations

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExpressRouteCircuitAuthorizationsClient struct {
	Client *resourcemanager.Client
}

func NewExpressRouteCircuitAuthorizationsClientWithBaseURI(sdkApi sdkEnv.Api) (*ExpressRouteCircuitAuthorizationsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "expressroutecircuitauthorizations", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ExpressRouteCircuitAuthorizationsClient: %+v", err)
	}

	return &ExpressRouteCircuitAuthorizationsClient{
		Client: client,
	}, nil
}

package expressroutecircuitroutestablesummary

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExpressRouteCircuitRoutesTableSummaryClient struct {
	Client *resourcemanager.Client
}

func NewExpressRouteCircuitRoutesTableSummaryClientWithBaseURI(sdkApi sdkEnv.Api) (*ExpressRouteCircuitRoutesTableSummaryClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "expressroutecircuitroutestablesummary", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ExpressRouteCircuitRoutesTableSummaryClient: %+v", err)
	}

	return &ExpressRouteCircuitRoutesTableSummaryClient{
		Client: client,
	}, nil
}

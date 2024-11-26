package oraclesubscriptions

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OracleSubscriptionsClient struct {
	Client *resourcemanager.Client
}

func NewOracleSubscriptionsClientWithBaseURI(sdkApi sdkEnv.Api) (*OracleSubscriptionsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "oraclesubscriptions", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating OracleSubscriptionsClient: %+v", err)
	}

	return &OracleSubscriptionsClient{
		Client: client,
	}, nil
}

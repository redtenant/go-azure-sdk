package v2023_06_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/resource-manager/maps/2023-06-01/accounts"
	"github.com/redtenant/go-azure-sdk/resource-manager/maps/2023-06-01/creators"
	"github.com/redtenant/go-azure-sdk/resource-manager/maps/2023-06-01/operations"
	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

type Client struct {
	Accounts   *accounts.AccountsClient
	Creators   *creators.CreatorsClient
	Operations *operations.OperationsClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	accountsClient, err := accounts.NewAccountsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Accounts client: %+v", err)
	}
	configureFunc(accountsClient.Client)

	creatorsClient, err := creators.NewCreatorsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Creators client: %+v", err)
	}
	configureFunc(creatorsClient.Client)

	operationsClient, err := operations.NewOperationsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Operations client: %+v", err)
	}
	configureFunc(operationsClient.Client)

	return &Client{
		Accounts:   accountsClient,
		Creators:   creatorsClient,
		Operations: operationsClient,
	}, nil
}

package accounts

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccountsClient struct {
	Client *resourcemanager.Client
}

func NewAccountsClientWithBaseURI(sdkApi sdkEnv.Api) (*AccountsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "accounts", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AccountsClient: %+v", err)
	}

	return &AccountsClient{
		Client: client,
	}, nil
}

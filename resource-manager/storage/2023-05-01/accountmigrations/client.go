package accountmigrations

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccountMigrationsClient struct {
	Client *resourcemanager.Client
}

func NewAccountMigrationsClientWithBaseURI(sdkApi sdkEnv.Api) (*AccountMigrationsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "accountmigrations", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AccountMigrationsClient: %+v", err)
	}

	return &AccountMigrationsClient{
		Client: client,
	}, nil
}

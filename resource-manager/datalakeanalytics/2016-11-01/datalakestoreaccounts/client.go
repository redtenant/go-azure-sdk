package datalakestoreaccounts

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataLakeStoreAccountsClient struct {
	Client *resourcemanager.Client
}

func NewDataLakeStoreAccountsClientWithBaseURI(sdkApi sdkEnv.Api) (*DataLakeStoreAccountsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "datalakestoreaccounts", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DataLakeStoreAccountsClient: %+v", err)
	}

	return &DataLakeStoreAccountsClient{
		Client: client,
	}, nil
}

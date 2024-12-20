package recoverabledatabases

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecoverableDatabasesClient struct {
	Client *resourcemanager.Client
}

func NewRecoverableDatabasesClientWithBaseURI(sdkApi sdkEnv.Api) (*RecoverableDatabasesClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "recoverabledatabases", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RecoverableDatabasesClient: %+v", err)
	}

	return &RecoverableDatabasesClient{
		Client: client,
	}, nil
}

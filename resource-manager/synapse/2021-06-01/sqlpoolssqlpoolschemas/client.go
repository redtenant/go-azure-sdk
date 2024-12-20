package sqlpoolssqlpoolschemas

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SqlPoolsSqlPoolSchemasClient struct {
	Client *resourcemanager.Client
}

func NewSqlPoolsSqlPoolSchemasClientWithBaseURI(sdkApi sdkEnv.Api) (*SqlPoolsSqlPoolSchemasClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "sqlpoolssqlpoolschemas", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SqlPoolsSqlPoolSchemasClient: %+v", err)
	}

	return &SqlPoolsSqlPoolSchemasClient{
		Client: client,
	}, nil
}

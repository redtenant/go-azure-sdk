package sqlpoolsblobauditing

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SqlPoolsBlobAuditingClient struct {
	Client *resourcemanager.Client
}

func NewSqlPoolsBlobAuditingClientWithBaseURI(sdkApi sdkEnv.Api) (*SqlPoolsBlobAuditingClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "sqlpoolsblobauditing", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SqlPoolsBlobAuditingClient: %+v", err)
	}

	return &SqlPoolsBlobAuditingClient{
		Client: client,
	}, nil
}

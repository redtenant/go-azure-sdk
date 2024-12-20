package tableserviceproperties

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TableServicePropertiesClient struct {
	Client *resourcemanager.Client
}

func NewTableServicePropertiesClientWithBaseURI(sdkApi sdkEnv.Api) (*TableServicePropertiesClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "tableserviceproperties", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TableServicePropertiesClient: %+v", err)
	}

	return &TableServicePropertiesClient{
		Client: client,
	}, nil
}

package webapps

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WebAppsClient struct {
	Client *resourcemanager.Client
}

func NewWebAppsClientWithBaseURI(sdkApi sdkEnv.Api) (*WebAppsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "webapps", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating WebAppsClient: %+v", err)
	}

	return &WebAppsClient{
		Client: client,
	}, nil
}

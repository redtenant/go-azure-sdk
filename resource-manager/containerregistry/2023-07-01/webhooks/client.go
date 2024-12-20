package webhooks

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WebHooksClient struct {
	Client *resourcemanager.Client
}

func NewWebHooksClientWithBaseURI(sdkApi sdkEnv.Api) (*WebHooksClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "webhooks", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating WebHooksClient: %+v", err)
	}

	return &WebHooksClient{
		Client: client,
	}, nil
}

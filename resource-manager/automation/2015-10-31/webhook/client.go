package webhook

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WebhookClient struct {
	Client *resourcemanager.Client
}

func NewWebhookClientWithBaseURI(sdkApi sdkEnv.Api) (*WebhookClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "webhook", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating WebhookClient: %+v", err)
	}

	return &WebhookClient{
		Client: client,
	}, nil
}

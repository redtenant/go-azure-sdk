package privatelinks

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivateLinksClient struct {
	Client *resourcemanager.Client
}

func NewPrivateLinksClientWithBaseURI(sdkApi sdkEnv.Api) (*PrivateLinksClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "privatelinks", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PrivateLinksClient: %+v", err)
	}

	return &PrivateLinksClient{
		Client: client,
	}, nil
}

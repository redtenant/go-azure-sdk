package authorizationloginlinks

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthorizationLoginLinksClient struct {
	Client *resourcemanager.Client
}

func NewAuthorizationLoginLinksClientWithBaseURI(sdkApi sdkEnv.Api) (*AuthorizationLoginLinksClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "authorizationloginlinks", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AuthorizationLoginLinksClient: %+v", err)
	}

	return &AuthorizationLoginLinksClient{
		Client: client,
	}, nil
}

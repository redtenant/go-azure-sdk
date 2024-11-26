package authenticationplatformcredentialmethod

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationPlatformCredentialMethodClient struct {
	Client *msgraph.Client
}

func NewAuthenticationPlatformCredentialMethodClientWithBaseURI(sdkApi sdkEnv.Api) (*AuthenticationPlatformCredentialMethodClient, error) {
	client, err := msgraph.NewClient(sdkApi, "authenticationplatformcredentialmethod", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AuthenticationPlatformCredentialMethodClient: %+v", err)
	}

	return &AuthenticationPlatformCredentialMethodClient{
		Client: client,
	}, nil
}

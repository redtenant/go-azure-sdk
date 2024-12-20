package authenticationphonemethod

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationPhoneMethodClient struct {
	Client *msgraph.Client
}

func NewAuthenticationPhoneMethodClientWithBaseURI(sdkApi sdkEnv.Api) (*AuthenticationPhoneMethodClient, error) {
	client, err := msgraph.NewClient(sdkApi, "authenticationphonemethod", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AuthenticationPhoneMethodClient: %+v", err)
	}

	return &AuthenticationPhoneMethodClient{
		Client: client,
	}, nil
}

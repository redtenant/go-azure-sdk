package signin

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SignInClient struct {
	Client *msgraph.Client
}

func NewSignInClientWithBaseURI(sdkApi sdkEnv.Api) (*SignInClient, error) {
	client, err := msgraph.NewClient(sdkApi, "signin", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SignInClient: %+v", err)
	}

	return &SignInClient{
		Client: client,
	}, nil
}

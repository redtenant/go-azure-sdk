package serviceprincipalsigninactivity

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServicePrincipalSignInActivityClient struct {
	Client *msgraph.Client
}

func NewServicePrincipalSignInActivityClientWithBaseURI(sdkApi sdkEnv.Api) (*ServicePrincipalSignInActivityClient, error) {
	client, err := msgraph.NewClient(sdkApi, "serviceprincipalsigninactivity", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ServicePrincipalSignInActivityClient: %+v", err)
	}

	return &ServicePrincipalSignInActivityClient{
		Client: client,
	}, nil
}

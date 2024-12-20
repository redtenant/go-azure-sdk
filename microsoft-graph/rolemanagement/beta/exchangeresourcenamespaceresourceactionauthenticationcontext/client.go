package exchangeresourcenamespaceresourceactionauthenticationcontext

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExchangeResourceNamespaceResourceActionAuthenticationContextClient struct {
	Client *msgraph.Client
}

func NewExchangeResourceNamespaceResourceActionAuthenticationContextClientWithBaseURI(sdkApi sdkEnv.Api) (*ExchangeResourceNamespaceResourceActionAuthenticationContextClient, error) {
	client, err := msgraph.NewClient(sdkApi, "exchangeresourcenamespaceresourceactionauthenticationcontext", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ExchangeResourceNamespaceResourceActionAuthenticationContextClient: %+v", err)
	}

	return &ExchangeResourceNamespaceResourceActionAuthenticationContextClient{
		Client: client,
	}, nil
}

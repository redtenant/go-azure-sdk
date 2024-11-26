package exchangeresourcenamespaceresourceactionresourcescope

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExchangeResourceNamespaceResourceActionResourceScopeClient struct {
	Client *msgraph.Client
}

func NewExchangeResourceNamespaceResourceActionResourceScopeClientWithBaseURI(sdkApi sdkEnv.Api) (*ExchangeResourceNamespaceResourceActionResourceScopeClient, error) {
	client, err := msgraph.NewClient(sdkApi, "exchangeresourcenamespaceresourceactionresourcescope", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ExchangeResourceNamespaceResourceActionResourceScopeClient: %+v", err)
	}

	return &ExchangeResourceNamespaceResourceActionResourceScopeClient{
		Client: client,
	}, nil
}

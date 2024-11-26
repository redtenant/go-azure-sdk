package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/microsoft-graph/identity/beta/conditionalaccessauthenticationcontextclassreference"
	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

type Client struct {
	ConditionalAccessAuthenticationContextClassReference *conditionalaccessauthenticationcontextclassreference.ConditionalAccessAuthenticationContextClassReferenceClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *msgraph.Client)) (*Client, error) {
	conditionalAccessAuthenticationContextClassReferenceClient, err := conditionalaccessauthenticationcontextclassreference.NewConditionalAccessAuthenticationContextClassReferenceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ConditionalAccessAuthenticationContextClassReference client: %+v", err)
	}
	configureFunc(conditionalAccessAuthenticationContextClassReferenceClient.Client)

	return &Client{
		ConditionalAccessAuthenticationContextClassReference: conditionalAccessAuthenticationContextClassReferenceClient,
	}, nil
}

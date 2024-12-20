package exchangetransitiveroleassignment

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExchangeTransitiveRoleAssignmentClient struct {
	Client *msgraph.Client
}

func NewExchangeTransitiveRoleAssignmentClientWithBaseURI(sdkApi sdkEnv.Api) (*ExchangeTransitiveRoleAssignmentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "exchangetransitiveroleassignment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ExchangeTransitiveRoleAssignmentClient: %+v", err)
	}

	return &ExchangeTransitiveRoleAssignmentClient{
		Client: client,
	}, nil
}

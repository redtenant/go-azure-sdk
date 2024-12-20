package sla

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SlaClient struct {
	Client *msgraph.Client
}

func NewSlaClientWithBaseURI(sdkApi sdkEnv.Api) (*SlaClient, error) {
	client, err := msgraph.NewClient(sdkApi, "sla", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SlaClient: %+v", err)
	}

	return &SlaClient{
		Client: client,
	}, nil
}

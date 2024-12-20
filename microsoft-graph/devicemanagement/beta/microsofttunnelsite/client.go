package microsofttunnelsite

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MicrosoftTunnelSiteClient struct {
	Client *msgraph.Client
}

func NewMicrosoftTunnelSiteClientWithBaseURI(sdkApi sdkEnv.Api) (*MicrosoftTunnelSiteClient, error) {
	client, err := msgraph.NewClient(sdkApi, "microsofttunnelsite", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MicrosoftTunnelSiteClient: %+v", err)
	}

	return &MicrosoftTunnelSiteClient{
		Client: client,
	}, nil
}

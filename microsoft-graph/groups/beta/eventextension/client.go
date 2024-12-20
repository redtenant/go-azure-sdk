package eventextension

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EventExtensionClient struct {
	Client *msgraph.Client
}

func NewEventExtensionClientWithBaseURI(sdkApi sdkEnv.Api) (*EventExtensionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "eventextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EventExtensionClient: %+v", err)
	}

	return &EventExtensionClient{
		Client: client,
	}, nil
}

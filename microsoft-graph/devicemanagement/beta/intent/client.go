package intent

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IntentClient struct {
	Client *msgraph.Client
}

func NewIntentClientWithBaseURI(sdkApi sdkEnv.Api) (*IntentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "intent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating IntentClient: %+v", err)
	}

	return &IntentClient{
		Client: client,
	}, nil
}

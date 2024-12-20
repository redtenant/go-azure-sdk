package querytexts

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type QueryTextsClient struct {
	Client *resourcemanager.Client
}

func NewQueryTextsClientWithBaseURI(sdkApi sdkEnv.Api) (*QueryTextsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "querytexts", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating QueryTextsClient: %+v", err)
	}

	return &QueryTextsClient{
		Client: client,
	}, nil
}

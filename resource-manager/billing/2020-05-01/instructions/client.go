package instructions

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InstructionsClient struct {
	Client *resourcemanager.Client
}

func NewInstructionsClientWithBaseURI(sdkApi sdkEnv.Api) (*InstructionsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "instructions", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating InstructionsClient: %+v", err)
	}

	return &InstructionsClient{
		Client: client,
	}, nil
}

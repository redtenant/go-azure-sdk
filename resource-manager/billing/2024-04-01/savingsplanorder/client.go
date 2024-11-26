package savingsplanorder

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SavingsPlanOrderClient struct {
	Client *resourcemanager.Client
}

func NewSavingsPlanOrderClientWithBaseURI(sdkApi sdkEnv.Api) (*SavingsPlanOrderClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "savingsplanorder", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SavingsPlanOrderClient: %+v", err)
	}

	return &SavingsPlanOrderClient{
		Client: client,
	}, nil
}

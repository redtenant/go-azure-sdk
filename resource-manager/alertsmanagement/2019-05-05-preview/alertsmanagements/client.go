package alertsmanagements

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AlertsManagementsClient struct {
	Client *resourcemanager.Client
}

func NewAlertsManagementsClientWithBaseURI(sdkApi sdkEnv.Api) (*AlertsManagementsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "alertsmanagements", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AlertsManagementsClient: %+v", err)
	}

	return &AlertsManagementsClient{
		Client: client,
	}, nil
}

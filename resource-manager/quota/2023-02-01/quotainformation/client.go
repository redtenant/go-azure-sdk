package quotainformation

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type QuotaInformationClient struct {
	Client *resourcemanager.Client
}

func NewQuotaInformationClientWithBaseURI(sdkApi sdkEnv.Api) (*QuotaInformationClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "quotainformation", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating QuotaInformationClient: %+v", err)
	}

	return &QuotaInformationClient{
		Client: client,
	}, nil
}

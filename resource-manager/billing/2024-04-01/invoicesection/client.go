package invoicesection

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InvoiceSectionClient struct {
	Client *resourcemanager.Client
}

func NewInvoiceSectionClientWithBaseURI(sdkApi sdkEnv.Api) (*InvoiceSectionClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "invoicesection", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating InvoiceSectionClient: %+v", err)
	}

	return &InvoiceSectionClient{
		Client: client,
	}, nil
}

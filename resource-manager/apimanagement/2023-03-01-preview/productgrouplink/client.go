package productgrouplink

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProductGroupLinkClient struct {
	Client *resourcemanager.Client
}

func NewProductGroupLinkClientWithBaseURI(sdkApi sdkEnv.Api) (*ProductGroupLinkClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "productgrouplink", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ProductGroupLinkClient: %+v", err)
	}

	return &ProductGroupLinkClient{
		Client: client,
	}, nil
}

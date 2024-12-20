package contentproductpackages

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ContentProductPackagesClient struct {
	Client *resourcemanager.Client
}

func NewContentProductPackagesClientWithBaseURI(sdkApi sdkEnv.Api) (*ContentProductPackagesClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "contentproductpackages", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ContentProductPackagesClient: %+v", err)
	}

	return &ContentProductPackagesClient{
		Client: client,
	}, nil
}

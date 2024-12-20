package images

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ImagesClient struct {
	Client *resourcemanager.Client
}

func NewImagesClientWithBaseURI(sdkApi sdkEnv.Api) (*ImagesClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "images", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ImagesClient: %+v", err)
	}

	return &ImagesClient{
		Client: client,
	}, nil
}

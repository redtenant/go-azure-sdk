package sharedgalleryimages

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SharedGalleryImagesClient struct {
	Client *resourcemanager.Client
}

func NewSharedGalleryImagesClientWithBaseURI(sdkApi sdkEnv.Api) (*SharedGalleryImagesClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "sharedgalleryimages", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SharedGalleryImagesClient: %+v", err)
	}

	return &SharedGalleryImagesClient{
		Client: client,
	}, nil
}

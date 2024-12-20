package componentannotationsapis

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ComponentAnnotationsAPIsClient struct {
	Client *resourcemanager.Client
}

func NewComponentAnnotationsAPIsClientWithBaseURI(sdkApi sdkEnv.Api) (*ComponentAnnotationsAPIsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "componentannotationsapis", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ComponentAnnotationsAPIsClient: %+v", err)
	}

	return &ComponentAnnotationsAPIsClient{
		Client: client,
	}, nil
}

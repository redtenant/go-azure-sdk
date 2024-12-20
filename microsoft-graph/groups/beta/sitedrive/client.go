package sitedrive

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteDriveClient struct {
	Client *msgraph.Client
}

func NewSiteDriveClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteDriveClient, error) {
	client, err := msgraph.NewClient(sdkApi, "sitedrive", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteDriveClient: %+v", err)
	}

	return &SiteDriveClient{
		Client: client,
	}, nil
}

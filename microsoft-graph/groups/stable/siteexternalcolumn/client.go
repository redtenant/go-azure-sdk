package siteexternalcolumn

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteExternalColumnClient struct {
	Client *msgraph.Client
}

func NewSiteExternalColumnClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteExternalColumnClient, error) {
	client, err := msgraph.NewClient(sdkApi, "siteexternalcolumn", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteExternalColumnClient: %+v", err)
	}

	return &SiteExternalColumnClient{
		Client: client,
	}, nil
}

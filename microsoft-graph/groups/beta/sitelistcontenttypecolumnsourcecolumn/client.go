package sitelistcontenttypecolumnsourcecolumn

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteListContentTypeColumnSourceColumnClient struct {
	Client *msgraph.Client
}

func NewSiteListContentTypeColumnSourceColumnClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteListContentTypeColumnSourceColumnClient, error) {
	client, err := msgraph.NewClient(sdkApi, "sitelistcontenttypecolumnsourcecolumn", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteListContentTypeColumnSourceColumnClient: %+v", err)
	}

	return &SiteListContentTypeColumnSourceColumnClient{
		Client: client,
	}, nil
}

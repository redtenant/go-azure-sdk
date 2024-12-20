package sitelistlastmodifiedbyuser

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteListLastModifiedByUserClient struct {
	Client *msgraph.Client
}

func NewSiteListLastModifiedByUserClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteListLastModifiedByUserClient, error) {
	client, err := msgraph.NewClient(sdkApi, "sitelistlastmodifiedbyuser", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteListLastModifiedByUserClient: %+v", err)
	}

	return &SiteListLastModifiedByUserClient{
		Client: client,
	}, nil
}

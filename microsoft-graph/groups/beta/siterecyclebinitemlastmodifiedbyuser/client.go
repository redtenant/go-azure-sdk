package siterecyclebinitemlastmodifiedbyuser

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteRecycleBinItemLastModifiedByUserClient struct {
	Client *msgraph.Client
}

func NewSiteRecycleBinItemLastModifiedByUserClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteRecycleBinItemLastModifiedByUserClient, error) {
	client, err := msgraph.NewClient(sdkApi, "siterecyclebinitemlastmodifiedbyuser", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteRecycleBinItemLastModifiedByUserClient: %+v", err)
	}

	return &SiteRecycleBinItemLastModifiedByUserClient{
		Client: client,
	}, nil
}

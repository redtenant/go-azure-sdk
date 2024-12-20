package siterecyclebinitemlastmodifiedbyusermailboxsetting

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteRecycleBinItemLastModifiedByUserMailboxSettingClient struct {
	Client *msgraph.Client
}

func NewSiteRecycleBinItemLastModifiedByUserMailboxSettingClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteRecycleBinItemLastModifiedByUserMailboxSettingClient, error) {
	client, err := msgraph.NewClient(sdkApi, "siterecyclebinitemlastmodifiedbyusermailboxsetting", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteRecycleBinItemLastModifiedByUserMailboxSettingClient: %+v", err)
	}

	return &SiteRecycleBinItemLastModifiedByUserMailboxSettingClient{
		Client: client,
	}, nil
}

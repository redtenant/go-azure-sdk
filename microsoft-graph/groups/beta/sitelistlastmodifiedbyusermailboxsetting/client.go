package sitelistlastmodifiedbyusermailboxsetting

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteListLastModifiedByUserMailboxSettingClient struct {
	Client *msgraph.Client
}

func NewSiteListLastModifiedByUserMailboxSettingClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteListLastModifiedByUserMailboxSettingClient, error) {
	client, err := msgraph.NewClient(sdkApi, "sitelistlastmodifiedbyusermailboxsetting", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteListLastModifiedByUserMailboxSettingClient: %+v", err)
	}

	return &SiteListLastModifiedByUserMailboxSettingClient{
		Client: client,
	}, nil
}

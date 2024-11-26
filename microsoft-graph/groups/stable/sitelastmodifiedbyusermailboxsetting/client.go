package sitelastmodifiedbyusermailboxsetting

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteLastModifiedByUserMailboxSettingClient struct {
	Client *msgraph.Client
}

func NewSiteLastModifiedByUserMailboxSettingClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteLastModifiedByUserMailboxSettingClient, error) {
	client, err := msgraph.NewClient(sdkApi, "sitelastmodifiedbyusermailboxsetting", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteLastModifiedByUserMailboxSettingClient: %+v", err)
	}

	return &SiteLastModifiedByUserMailboxSettingClient{
		Client: client,
	}, nil
}

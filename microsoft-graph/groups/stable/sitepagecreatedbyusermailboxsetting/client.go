package sitepagecreatedbyusermailboxsetting

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SitePageCreatedByUserMailboxSettingClient struct {
	Client *msgraph.Client
}

func NewSitePageCreatedByUserMailboxSettingClientWithBaseURI(sdkApi sdkEnv.Api) (*SitePageCreatedByUserMailboxSettingClient, error) {
	client, err := msgraph.NewClient(sdkApi, "sitepagecreatedbyusermailboxsetting", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SitePageCreatedByUserMailboxSettingClient: %+v", err)
	}

	return &SitePageCreatedByUserMailboxSettingClient{
		Client: client,
	}, nil
}

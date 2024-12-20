package siteonenotesectiongroupsectiongroup

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteOnenoteSectionGroupSectionGroupClient struct {
	Client *msgraph.Client
}

func NewSiteOnenoteSectionGroupSectionGroupClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteOnenoteSectionGroupSectionGroupClient, error) {
	client, err := msgraph.NewClient(sdkApi, "siteonenotesectiongroupsectiongroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteOnenoteSectionGroupSectionGroupClient: %+v", err)
	}

	return &SiteOnenoteSectionGroupSectionGroupClient{
		Client: client,
	}, nil
}

package siteanalyticsitemactivitystatactivitydriveitemcontentstream

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteAnalyticsItemActivityStatActivityDriveItemContentStreamClient struct {
	Client *msgraph.Client
}

func NewSiteAnalyticsItemActivityStatActivityDriveItemContentStreamClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteAnalyticsItemActivityStatActivityDriveItemContentStreamClient, error) {
	client, err := msgraph.NewClient(sdkApi, "siteanalyticsitemactivitystatactivitydriveitemcontentstream", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteAnalyticsItemActivityStatActivityDriveItemContentStreamClient: %+v", err)
	}

	return &SiteAnalyticsItemActivityStatActivityDriveItemContentStreamClient{
		Client: client,
	}, nil
}

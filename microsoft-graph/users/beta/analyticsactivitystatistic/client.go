package analyticsactivitystatistic

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AnalyticsActivityStatisticClient struct {
	Client *msgraph.Client
}

func NewAnalyticsActivityStatisticClientWithBaseURI(sdkApi sdkEnv.Api) (*AnalyticsActivityStatisticClient, error) {
	client, err := msgraph.NewClient(sdkApi, "analyticsactivitystatistic", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AnalyticsActivityStatisticClient: %+v", err)
	}

	return &AnalyticsActivityStatisticClient{
		Client: client,
	}, nil
}

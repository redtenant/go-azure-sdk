package analytics

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AnalyticsClient struct {
	Client *msgraph.Client
}

func NewAnalyticsClientWithBaseURI(sdkApi sdkEnv.Api) (*AnalyticsClient, error) {
	client, err := msgraph.NewClient(sdkApi, "analytics", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AnalyticsClient: %+v", err)
	}

	return &AnalyticsClient{
		Client: client,
	}, nil
}

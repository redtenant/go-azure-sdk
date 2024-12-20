package userexperienceanalyticsanomalycorrelationgroupoverview

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsAnomalyCorrelationGroupOverviewClient struct {
	Client *msgraph.Client
}

func NewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewClientWithBaseURI(sdkApi sdkEnv.Api) (*UserExperienceAnalyticsAnomalyCorrelationGroupOverviewClient, error) {
	client, err := msgraph.NewClient(sdkApi, "userexperienceanalyticsanomalycorrelationgroupoverview", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserExperienceAnalyticsAnomalyCorrelationGroupOverviewClient: %+v", err)
	}

	return &UserExperienceAnalyticsAnomalyCorrelationGroupOverviewClient{
		Client: client,
	}, nil
}

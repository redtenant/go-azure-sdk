package userexperienceanalyticsworkfromanywheremetric

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsWorkFromAnywhereMetricClient struct {
	Client *msgraph.Client
}

func NewUserExperienceAnalyticsWorkFromAnywhereMetricClientWithBaseURI(sdkApi sdkEnv.Api) (*UserExperienceAnalyticsWorkFromAnywhereMetricClient, error) {
	client, err := msgraph.NewClient(sdkApi, "userexperienceanalyticsworkfromanywheremetric", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserExperienceAnalyticsWorkFromAnywhereMetricClient: %+v", err)
	}

	return &UserExperienceAnalyticsWorkFromAnywhereMetricClient{
		Client: client,
	}, nil
}

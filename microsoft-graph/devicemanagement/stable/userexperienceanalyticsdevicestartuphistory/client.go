package userexperienceanalyticsdevicestartuphistory

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsDeviceStartupHistoryClient struct {
	Client *msgraph.Client
}

func NewUserExperienceAnalyticsDeviceStartupHistoryClientWithBaseURI(sdkApi sdkEnv.Api) (*UserExperienceAnalyticsDeviceStartupHistoryClient, error) {
	client, err := msgraph.NewClient(sdkApi, "userexperienceanalyticsdevicestartuphistory", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserExperienceAnalyticsDeviceStartupHistoryClient: %+v", err)
	}

	return &UserExperienceAnalyticsDeviceStartupHistoryClient{
		Client: client,
	}, nil
}

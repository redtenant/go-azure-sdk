package userexperienceanalyticsdevicestartupprocess

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsDeviceStartupProcessClient struct {
	Client *msgraph.Client
}

func NewUserExperienceAnalyticsDeviceStartupProcessClientWithBaseURI(sdkApi sdkEnv.Api) (*UserExperienceAnalyticsDeviceStartupProcessClient, error) {
	client, err := msgraph.NewClient(sdkApi, "userexperienceanalyticsdevicestartupprocess", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserExperienceAnalyticsDeviceStartupProcessClient: %+v", err)
	}

	return &UserExperienceAnalyticsDeviceStartupProcessClient{
		Client: client,
	}, nil
}

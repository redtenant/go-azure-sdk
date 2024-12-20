package userexperienceanalyticsbatteryhealthcapacitydetail

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsBatteryHealthCapacityDetailClient struct {
	Client *msgraph.Client
}

func NewUserExperienceAnalyticsBatteryHealthCapacityDetailClientWithBaseURI(sdkApi sdkEnv.Api) (*UserExperienceAnalyticsBatteryHealthCapacityDetailClient, error) {
	client, err := msgraph.NewClient(sdkApi, "userexperienceanalyticsbatteryhealthcapacitydetail", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserExperienceAnalyticsBatteryHealthCapacityDetailClient: %+v", err)
	}

	return &UserExperienceAnalyticsBatteryHealthCapacityDetailClient{
		Client: client,
	}, nil
}

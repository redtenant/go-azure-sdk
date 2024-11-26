package userinsightdaily

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserInsightDailyClient struct {
	Client *msgraph.Client
}

func NewUserInsightDailyClientWithBaseURI(sdkApi sdkEnv.Api) (*UserInsightDailyClient, error) {
	client, err := msgraph.NewClient(sdkApi, "userinsightdaily", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserInsightDailyClient: %+v", err)
	}

	return &UserInsightDailyClient{
		Client: client,
	}, nil
}

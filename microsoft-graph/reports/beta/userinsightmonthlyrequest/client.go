package userinsightmonthlyrequest

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserInsightMonthlyRequestClient struct {
	Client *msgraph.Client
}

func NewUserInsightMonthlyRequestClientWithBaseURI(sdkApi sdkEnv.Api) (*UserInsightMonthlyRequestClient, error) {
	client, err := msgraph.NewClient(sdkApi, "userinsightmonthlyrequest", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserInsightMonthlyRequestClient: %+v", err)
	}

	return &UserInsightMonthlyRequestClient{
		Client: client,
	}, nil
}

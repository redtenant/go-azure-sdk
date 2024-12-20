package usersubscription

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserSubscriptionClient struct {
	Client *resourcemanager.Client
}

func NewUserSubscriptionClientWithBaseURI(sdkApi sdkEnv.Api) (*UserSubscriptionClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "usersubscription", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserSubscriptionClient: %+v", err)
	}

	return &UserSubscriptionClient{
		Client: client,
	}, nil
}

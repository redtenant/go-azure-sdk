package v2018_04_16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/resource-manager/insights/2018-04-16/scheduledqueryrules"
	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

type Client struct {
	ScheduledQueryRules *scheduledqueryrules.ScheduledQueryRulesClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	scheduledQueryRulesClient, err := scheduledqueryrules.NewScheduledQueryRulesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ScheduledQueryRules client: %+v", err)
	}
	configureFunc(scheduledQueryRulesClient.Client)

	return &Client{
		ScheduledQueryRules: scheduledQueryRulesClient,
	}, nil
}

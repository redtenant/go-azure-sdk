package webapplicationfirewallpolicies

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WebApplicationFirewallPoliciesClient struct {
	Client *resourcemanager.Client
}

func NewWebApplicationFirewallPoliciesClientWithBaseURI(sdkApi sdkEnv.Api) (*WebApplicationFirewallPoliciesClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "webapplicationfirewallpolicies", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating WebApplicationFirewallPoliciesClient: %+v", err)
	}

	return &WebApplicationFirewallPoliciesClient{
		Client: client,
	}, nil
}

package listtenantconfigurationviolationsoperations

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListTenantConfigurationViolationsOperationsClient struct {
	Client *resourcemanager.Client
}

func NewListTenantConfigurationViolationsOperationsClientWithBaseURI(sdkApi sdkEnv.Api) (*ListTenantConfigurationViolationsOperationsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "listtenantconfigurationviolationsoperations", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ListTenantConfigurationViolationsOperationsClient: %+v", err)
	}

	return &ListTenantConfigurationViolationsOperationsClient{
		Client: client,
	}, nil
}

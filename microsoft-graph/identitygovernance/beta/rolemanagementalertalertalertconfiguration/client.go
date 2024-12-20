package rolemanagementalertalertalertconfiguration

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleManagementAlertAlertAlertConfigurationClient struct {
	Client *msgraph.Client
}

func NewRoleManagementAlertAlertAlertConfigurationClientWithBaseURI(sdkApi sdkEnv.Api) (*RoleManagementAlertAlertAlertConfigurationClient, error) {
	client, err := msgraph.NewClient(sdkApi, "rolemanagementalertalertalertconfiguration", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleManagementAlertAlertAlertConfigurationClient: %+v", err)
	}

	return &RoleManagementAlertAlertAlertConfigurationClient{
		Client: client,
	}, nil
}

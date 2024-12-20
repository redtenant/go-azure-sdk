package cloudpcroledefinitioninheritspermissionsfrom

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCRoleDefinitionInheritsPermissionsFromClient struct {
	Client *msgraph.Client
}

func NewCloudPCRoleDefinitionInheritsPermissionsFromClientWithBaseURI(sdkApi sdkEnv.Api) (*CloudPCRoleDefinitionInheritsPermissionsFromClient, error) {
	client, err := msgraph.NewClient(sdkApi, "cloudpcroledefinitioninheritspermissionsfrom", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CloudPCRoleDefinitionInheritsPermissionsFromClient: %+v", err)
	}

	return &CloudPCRoleDefinitionInheritsPermissionsFromClient{
		Client: client,
	}, nil
}

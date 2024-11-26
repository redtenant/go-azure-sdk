package directoryroleassignmentscheduleinstanceroledefinition

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DirectoryRoleAssignmentScheduleInstanceRoleDefinitionClient struct {
	Client *msgraph.Client
}

func NewDirectoryRoleAssignmentScheduleInstanceRoleDefinitionClientWithBaseURI(sdkApi sdkEnv.Api) (*DirectoryRoleAssignmentScheduleInstanceRoleDefinitionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "directoryroleassignmentscheduleinstanceroledefinition", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DirectoryRoleAssignmentScheduleInstanceRoleDefinitionClient: %+v", err)
	}

	return &DirectoryRoleAssignmentScheduleInstanceRoleDefinitionClient{
		Client: client,
	}, nil
}

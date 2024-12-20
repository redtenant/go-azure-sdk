package directoryroleassignmentscheduleinstancedirectoryscope

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DirectoryRoleAssignmentScheduleInstanceDirectoryScopeClient struct {
	Client *msgraph.Client
}

func NewDirectoryRoleAssignmentScheduleInstanceDirectoryScopeClientWithBaseURI(sdkApi sdkEnv.Api) (*DirectoryRoleAssignmentScheduleInstanceDirectoryScopeClient, error) {
	client, err := msgraph.NewClient(sdkApi, "directoryroleassignmentscheduleinstancedirectoryscope", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DirectoryRoleAssignmentScheduleInstanceDirectoryScopeClient: %+v", err)
	}

	return &DirectoryRoleAssignmentScheduleInstanceDirectoryScopeClient{
		Client: client,
	}, nil
}

package directoryroleassignmentschedulerequest

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DirectoryRoleAssignmentScheduleRequestClient struct {
	Client *msgraph.Client
}

func NewDirectoryRoleAssignmentScheduleRequestClientWithBaseURI(sdkApi sdkEnv.Api) (*DirectoryRoleAssignmentScheduleRequestClient, error) {
	client, err := msgraph.NewClient(sdkApi, "directoryroleassignmentschedulerequest", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DirectoryRoleAssignmentScheduleRequestClient: %+v", err)
	}

	return &DirectoryRoleAssignmentScheduleRequestClient{
		Client: client,
	}, nil
}

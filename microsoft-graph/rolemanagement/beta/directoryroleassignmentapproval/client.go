package directoryroleassignmentapproval

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DirectoryRoleAssignmentApprovalClient struct {
	Client *msgraph.Client
}

func NewDirectoryRoleAssignmentApprovalClientWithBaseURI(sdkApi sdkEnv.Api) (*DirectoryRoleAssignmentApprovalClient, error) {
	client, err := msgraph.NewClient(sdkApi, "directoryroleassignmentapproval", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DirectoryRoleAssignmentApprovalClient: %+v", err)
	}

	return &DirectoryRoleAssignmentApprovalClient{
		Client: client,
	}, nil
}

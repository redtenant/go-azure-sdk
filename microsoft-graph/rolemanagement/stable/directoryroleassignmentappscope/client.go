package directoryroleassignmentappscope

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DirectoryRoleAssignmentAppScopeClient struct {
	Client *msgraph.Client
}

func NewDirectoryRoleAssignmentAppScopeClientWithBaseURI(sdkApi sdkEnv.Api) (*DirectoryRoleAssignmentAppScopeClient, error) {
	client, err := msgraph.NewClient(sdkApi, "directoryroleassignmentappscope", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DirectoryRoleAssignmentAppScopeClient: %+v", err)
	}

	return &DirectoryRoleAssignmentAppScopeClient{
		Client: client,
	}, nil
}

package rolescopetagassignment

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleScopeTagAssignmentClient struct {
	Client *msgraph.Client
}

func NewRoleScopeTagAssignmentClientWithBaseURI(sdkApi sdkEnv.Api) (*RoleScopeTagAssignmentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "rolescopetagassignment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleScopeTagAssignmentClient: %+v", err)
	}

	return &RoleScopeTagAssignmentClient{
		Client: client,
	}, nil
}

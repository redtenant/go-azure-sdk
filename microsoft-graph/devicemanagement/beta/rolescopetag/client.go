package rolescopetag

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleScopeTagClient struct {
	Client *msgraph.Client
}

func NewRoleScopeTagClientWithBaseURI(sdkApi sdkEnv.Api) (*RoleScopeTagClient, error) {
	client, err := msgraph.NewClient(sdkApi, "rolescopetag", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleScopeTagClient: %+v", err)
	}

	return &RoleScopeTagClient{
		Client: client,
	}, nil
}

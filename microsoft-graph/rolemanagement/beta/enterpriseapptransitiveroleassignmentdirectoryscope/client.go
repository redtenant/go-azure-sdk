package enterpriseapptransitiveroleassignmentdirectoryscope

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EnterpriseAppTransitiveRoleAssignmentDirectoryScopeClient struct {
	Client *msgraph.Client
}

func NewEnterpriseAppTransitiveRoleAssignmentDirectoryScopeClientWithBaseURI(sdkApi sdkEnv.Api) (*EnterpriseAppTransitiveRoleAssignmentDirectoryScopeClient, error) {
	client, err := msgraph.NewClient(sdkApi, "enterpriseapptransitiveroleassignmentdirectoryscope", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EnterpriseAppTransitiveRoleAssignmentDirectoryScopeClient: %+v", err)
	}

	return &EnterpriseAppTransitiveRoleAssignmentDirectoryScopeClient{
		Client: client,
	}, nil
}

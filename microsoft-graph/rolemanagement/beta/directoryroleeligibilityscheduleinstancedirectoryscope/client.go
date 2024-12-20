package directoryroleeligibilityscheduleinstancedirectoryscope

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DirectoryRoleEligibilityScheduleInstanceDirectoryScopeClient struct {
	Client *msgraph.Client
}

func NewDirectoryRoleEligibilityScheduleInstanceDirectoryScopeClientWithBaseURI(sdkApi sdkEnv.Api) (*DirectoryRoleEligibilityScheduleInstanceDirectoryScopeClient, error) {
	client, err := msgraph.NewClient(sdkApi, "directoryroleeligibilityscheduleinstancedirectoryscope", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DirectoryRoleEligibilityScheduleInstanceDirectoryScopeClient: %+v", err)
	}

	return &DirectoryRoleEligibilityScheduleInstanceDirectoryScopeClient{
		Client: client,
	}, nil
}

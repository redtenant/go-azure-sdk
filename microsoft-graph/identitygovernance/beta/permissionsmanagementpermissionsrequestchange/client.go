package permissionsmanagementpermissionsrequestchange

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PermissionsManagementPermissionsRequestChangeClient struct {
	Client *msgraph.Client
}

func NewPermissionsManagementPermissionsRequestChangeClientWithBaseURI(sdkApi sdkEnv.Api) (*PermissionsManagementPermissionsRequestChangeClient, error) {
	client, err := msgraph.NewClient(sdkApi, "permissionsmanagementpermissionsrequestchange", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PermissionsManagementPermissionsRequestChangeClient: %+v", err)
	}

	return &PermissionsManagementPermissionsRequestChangeClient{
		Client: client,
	}, nil
}

package devicemanagementresourcenamespaceresourceactionresourcescope

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementResourceNamespaceResourceActionResourceScopeClient struct {
	Client *msgraph.Client
}

func NewDeviceManagementResourceNamespaceResourceActionResourceScopeClientWithBaseURI(sdkApi sdkEnv.Api) (*DeviceManagementResourceNamespaceResourceActionResourceScopeClient, error) {
	client, err := msgraph.NewClient(sdkApi, "devicemanagementresourcenamespaceresourceactionresourcescope", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DeviceManagementResourceNamespaceResourceActionResourceScopeClient: %+v", err)
	}

	return &DeviceManagementResourceNamespaceResourceActionResourceScopeClient{
		Client: client,
	}, nil
}

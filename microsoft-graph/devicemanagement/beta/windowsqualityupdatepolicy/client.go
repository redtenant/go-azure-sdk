package windowsqualityupdatepolicy

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsQualityUpdatePolicyClient struct {
	Client *msgraph.Client
}

func NewWindowsQualityUpdatePolicyClientWithBaseURI(sdkApi sdkEnv.Api) (*WindowsQualityUpdatePolicyClient, error) {
	client, err := msgraph.NewClient(sdkApi, "windowsqualityupdatepolicy", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating WindowsQualityUpdatePolicyClient: %+v", err)
	}

	return &WindowsQualityUpdatePolicyClient{
		Client: client,
	}, nil
}

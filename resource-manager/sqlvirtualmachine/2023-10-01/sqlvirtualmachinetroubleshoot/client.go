package sqlvirtualmachinetroubleshoot

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SqlVirtualMachineTroubleshootClient struct {
	Client *resourcemanager.Client
}

func NewSqlVirtualMachineTroubleshootClientWithBaseURI(sdkApi sdkEnv.Api) (*SqlVirtualMachineTroubleshootClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "sqlvirtualmachinetroubleshoot", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SqlVirtualMachineTroubleshootClient: %+v", err)
	}

	return &SqlVirtualMachineTroubleshootClient{
		Client: client,
	}, nil
}

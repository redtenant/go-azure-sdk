package windowsdriverupdateprofile

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsDriverUpdateProfileClient struct {
	Client *msgraph.Client
}

func NewWindowsDriverUpdateProfileClientWithBaseURI(sdkApi sdkEnv.Api) (*WindowsDriverUpdateProfileClient, error) {
	client, err := msgraph.NewClient(sdkApi, "windowsdriverupdateprofile", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating WindowsDriverUpdateProfileClient: %+v", err)
	}

	return &WindowsDriverUpdateProfileClient{
		Client: client,
	}, nil
}

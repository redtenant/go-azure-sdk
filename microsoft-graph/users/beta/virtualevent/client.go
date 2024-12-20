package virtualevent

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualEventClient struct {
	Client *msgraph.Client
}

func NewVirtualEventClientWithBaseURI(sdkApi sdkEnv.Api) (*VirtualEventClient, error) {
	client, err := msgraph.NewClient(sdkApi, "virtualevent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating VirtualEventClient: %+v", err)
	}

	return &VirtualEventClient{
		Client: client,
	}, nil
}

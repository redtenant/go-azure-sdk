package provisioning

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProvisioningClient struct {
	Client *msgraph.Client
}

func NewProvisioningClientWithBaseURI(sdkApi sdkEnv.Api) (*ProvisioningClient, error) {
	client, err := msgraph.NewClient(sdkApi, "provisioning", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ProvisioningClient: %+v", err)
	}

	return &ProvisioningClient{
		Client: client,
	}, nil
}

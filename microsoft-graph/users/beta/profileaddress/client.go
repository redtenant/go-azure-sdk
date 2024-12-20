package profileaddress

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProfileAddressClient struct {
	Client *msgraph.Client
}

func NewProfileAddressClientWithBaseURI(sdkApi sdkEnv.Api) (*ProfileAddressClient, error) {
	client, err := msgraph.NewClient(sdkApi, "profileaddress", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ProfileAddressClient: %+v", err)
	}

	return &ProfileAddressClient{
		Client: client,
	}, nil
}

package profilename

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProfileNameClient struct {
	Client *msgraph.Client
}

func NewProfileNameClientWithBaseURI(sdkApi sdkEnv.Api) (*ProfileNameClient, error) {
	client, err := msgraph.NewClient(sdkApi, "profilename", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ProfileNameClient: %+v", err)
	}

	return &ProfileNameClient{
		Client: client,
	}, nil
}

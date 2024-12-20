package profileaward

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProfileAwardClient struct {
	Client *msgraph.Client
}

func NewProfileAwardClientWithBaseURI(sdkApi sdkEnv.Api) (*ProfileAwardClient, error) {
	client, err := msgraph.NewClient(sdkApi, "profileaward", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ProfileAwardClient: %+v", err)
	}

	return &ProfileAwardClient{
		Client: client,
	}, nil
}

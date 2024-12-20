package activity

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ActivityClient struct {
	Client *msgraph.Client
}

func NewActivityClientWithBaseURI(sdkApi sdkEnv.Api) (*ActivityClient, error) {
	client, err := msgraph.NewClient(sdkApi, "activity", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ActivityClient: %+v", err)
	}

	return &ActivityClient{
		Client: client,
	}, nil
}

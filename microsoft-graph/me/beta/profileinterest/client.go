package profileinterest

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProfileInterestClient struct {
	Client *msgraph.Client
}

func NewProfileInterestClientWithBaseURI(sdkApi sdkEnv.Api) (*ProfileInterestClient, error) {
	client, err := msgraph.NewClient(sdkApi, "profileinterest", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ProfileInterestClient: %+v", err)
	}

	return &ProfileInterestClient{
		Client: client,
	}, nil
}

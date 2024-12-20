package profileanniversary

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProfileAnniversaryClient struct {
	Client *msgraph.Client
}

func NewProfileAnniversaryClientWithBaseURI(sdkApi sdkEnv.Api) (*ProfileAnniversaryClient, error) {
	client, err := msgraph.NewClient(sdkApi, "profileanniversary", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ProfileAnniversaryClient: %+v", err)
	}

	return &ProfileAnniversaryClient{
		Client: client,
	}, nil
}

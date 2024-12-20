package teamowner

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamOwnerClient struct {
	Client *msgraph.Client
}

func NewTeamOwnerClientWithBaseURI(sdkApi sdkEnv.Api) (*TeamOwnerClient, error) {
	client, err := msgraph.NewClient(sdkApi, "teamowner", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TeamOwnerClient: %+v", err)
	}

	return &TeamOwnerClient{
		Client: client,
	}, nil
}

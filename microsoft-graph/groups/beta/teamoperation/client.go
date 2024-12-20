package teamoperation

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamOperationClient struct {
	Client *msgraph.Client
}

func NewTeamOperationClientWithBaseURI(sdkApi sdkEnv.Api) (*TeamOperationClient, error) {
	client, err := msgraph.NewClient(sdkApi, "teamoperation", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TeamOperationClient: %+v", err)
	}

	return &TeamOperationClient{
		Client: client,
	}, nil
}

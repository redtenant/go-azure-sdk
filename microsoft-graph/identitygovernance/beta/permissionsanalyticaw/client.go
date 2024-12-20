package permissionsanalyticaw

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PermissionsAnalyticAwClient struct {
	Client *msgraph.Client
}

func NewPermissionsAnalyticAwClientWithBaseURI(sdkApi sdkEnv.Api) (*PermissionsAnalyticAwClient, error) {
	client, err := msgraph.NewClient(sdkApi, "permissionsanalyticaw", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PermissionsAnalyticAwClient: %+v", err)
	}

	return &PermissionsAnalyticAwClient{
		Client: client,
	}, nil
}

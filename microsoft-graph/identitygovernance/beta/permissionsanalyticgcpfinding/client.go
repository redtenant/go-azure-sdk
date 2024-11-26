package permissionsanalyticgcpfinding

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PermissionsAnalyticGcpFindingClient struct {
	Client *msgraph.Client
}

func NewPermissionsAnalyticGcpFindingClientWithBaseURI(sdkApi sdkEnv.Api) (*PermissionsAnalyticGcpFindingClient, error) {
	client, err := msgraph.NewClient(sdkApi, "permissionsanalyticgcpfinding", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PermissionsAnalyticGcpFindingClient: %+v", err)
	}

	return &PermissionsAnalyticGcpFindingClient{
		Client: client,
	}, nil
}

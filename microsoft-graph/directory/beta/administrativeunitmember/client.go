package administrativeunitmember

import (
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/redtenant/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AdministrativeUnitMemberClient struct {
	Client *msgraph.Client
}

func NewAdministrativeUnitMemberClientWithBaseURI(sdkApi sdkEnv.Api) (*AdministrativeUnitMemberClient, error) {
	client, err := msgraph.NewClient(sdkApi, "administrativeunitmember", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AdministrativeUnitMemberClient: %+v", err)
	}

	return &AdministrativeUnitMemberClient{
		Client: client,
	}, nil
}
